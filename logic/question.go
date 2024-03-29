package logic

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/SXUOJ/backend/dao"
	"github.com/SXUOJ/backend/models"
	"github.com/SXUOJ/backend/pb"
	"github.com/SXUOJ/backend/pkg/grpc"
	"github.com/SXUOJ/backend/pkg/uuid"
	"go.uber.org/zap"
)

func GetQuestionDetail(Qid string) (que *models.Question, err error) {
	//查库
	que, err = dao.GetQuestionDetail(Qid)
	if err != nil {
		zap.L().Error("GetQuestionDetail(Qid string) err...", zap.Error(err))
		return nil, err
	}
	return que, nil
}

func GetQuestionList(page int, amount int, uid string) (data *[]models.QueList, count int64, err error) {
	//查库
	data, count, err = dao.GetQuestionList(page, amount, uid)
	if err != nil {
		zap.L().Error("dao.GetQuestionList(page, amount) err ", zap.Error(err))
		return nil, 0, err
	}
	return data, count, nil
}

func GetSearchList(keyword string, amount int, page int, uid string) (data *[]models.QueList, nums int64, err error) {
	data, nums, err = dao.GetSearchList(keyword, page, amount, uid)
	if err != nil {
		zap.L().Error("dao.GetQuestionList(page, amount) err ", zap.Error(err))
		return nil, 0, err
	}
	return data, nums, nil
}

func CreateQuestion(que models.Question) error {
	err := dao.InsertQuestion(que)
	return err
}

func ChangeQuestion(qid string, que models.Question) error {
	err := dao.UpdateQuestion(qid, que)
	return err
}

func DelQuestion(qid string) error {
	err := dao.DeleteQuestion(qid)
	return err
}

func PushJudge(code models.Submit) (*models.SubmitResult, error) {
	//1.获取判题机地址并选择合适地址（暂时只便利）
	var addr string
	addrs, err := GetJudgerList()
	for _, a := range addrs {
		re, err := grpc.Ping(a.Addr)
		fmt.Println(re, a.Addr)
		if err == nil {
			addr = a.Addr
			break
		}
	}
	defer func() { //recover写在调用方兜底
		if r := recover(); r != nil {
			fmt.Printf("我是PushJudg方法中捕获的panic异常的recover，异常信息为：%s\n", r)
		}
	}()

	//2.构建grpc判题请求模型
	//2.1 创建代码id
	submitId, _ := uuid.Getuuid()
	Quest := pb.JudgeRequest{
		SubmitId: submitId,
		Source:   code.Source,
		Samples:  nil,
	}
	Quest.Type, _ = strconv.ParseUint(code.CodeType, 10, 64)
	Quest.TimeLimit, _ = strconv.ParseUint(code.TimeLimit, 10, 64)
	Quest.MemoryLimit, _ = strconv.ParseUint(code.MemoryLimit, 10, 64)
	Quest.MemoryLimit = Quest.MemoryLimit * 1024 * 1024
	//2.2获取样例
	res, err := os.ReadDir("./file/sample/" + code.QuestionID + "/sample")
	samples := []*pb.Sample{}
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(res); i += 2 {
		s := new(pb.Sample)
		for j := i; j < i+2; j++ {
			if path.Ext(res[j].Name()) == ".in" {
				b, err := os.ReadFile("./file/sample/" + code.QuestionID + "/sample/" + res[j].Name())
				s.Input = string(b)
				if err != nil {
					return nil, err
				}
			} else if path.Ext(res[j].Name()) == ".out" {
				b, err := os.ReadFile("./file/sample/" + code.QuestionID + "/sample/" + res[j].Name())
				if err != nil {
					return nil, err
				}
				s.Output = string(b)
			}
		}
		samples = append(samples, s)
	}
	Quest.Samples = samples

	//2.3传入判题机
	re, err := grpc.Judge(addr, &Quest)
	fmt.Println(re)
	if err != nil {
		return nil, err
	}
	var cstZone = time.FixedZone("CST", 8*3600) //上海
	//3.创建结果结构体
	Subid, _ := uuid.Getuuid()
	Result := models.Result{
		SubmitID:   Subid,
		QuestionID: code.QuestionID,
		UserID:     code.UserID,
		Public:     code.Public,
		Source:     code.Source,
		Time:       time.Now().In(cstZone).Format("2006-01-02 15:04:05"),
	}
	Result.CodeType, _ = strconv.ParseUint(code.CodeType, 10, 64)
	Result.IfAC = 1
	Results := []models.ResultOfOneSample{}
	for i := range re.Results {
		if re.Results[i].Status != 1 {
			Result.IfAC = 0
		}
		Results = append(Results, models.ResultOfOneSample{
			Status:   re.Results[i].Status,
			Memory:   strconv.FormatFloat(re.Results[i].Memory, 'f', -1, 32),
			RealTime: strconv.FormatFloat(re.Results[i].RealTime, 'f', -1, 32),
			CPUTime:  strconv.FormatFloat(re.Results[i].CpuTime, 'f', -1, 32),
			ErrorMsg: re.Results[i].Error,
		})
	}
	Rjson, err := json.Marshal(Results)
	if err != nil {
		return nil, err
	}
	Result.Results = string(Rjson)
	err = dao.InsertStatus(Result)
	if Result.IfAC == 1 {
		dao.InsertAc(models.Ac{
			UserID:       code.UserID,
			AcQuestionID: code.QuestionID,
			IfAC:         1,
		})
	}
	if err != nil {
		return nil, err
	}

	return &models.SubmitResult{
		SubmitID:   submitId,
		UserID:     Result.UserID,
		QuestionID: Result.QuestionID,
		Time:       Result.Time,
		IfAC:       Result.IfAC,
		Results:    Results,
	}, nil
}
