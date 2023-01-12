package logic

import (
	"github.com/SXUOJ/backend/dao"
	"github.com/SXUOJ/backend/models"
	"github.com/SXUOJ/backend/pb"
	"github.com/SXUOJ/backend/pkg/grpc"
	"github.com/SXUOJ/backend/pkg/uuid"
	"go.uber.org/zap"
	"os"
	"path"
	"strconv"
	"time"
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

func GetQuestionList(page int, amount int) (data []*models.Question, err error) {
	//查库
	data, err = dao.GetQuestionList(page, amount)
	if err != nil {
		zap.L().Error("dao.GetQuestionList(page, amount) err ", zap.Error(err))
		return nil, err
	}
	return data, nil
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

func PushJudge(code models.Code) (*models.Result, error) {
	//1.获取判题机地址并选择合适地址（暂时只便利）
	var addr string
	addrs, err := GetJudgerList()
	for _, a := range addrs {
		_, err := grpc.Ping(a.Addr)
		if err == nil {
			addr = a.Addr
			break
		}
	}
	//2.构建grpc判题请求模型
	//2.1 创建代码id
	code.CodeID, _ = uuid.Getuuid()
	Quest := pb.JudgeRequest{
		SubmitId:    code.CodeID,
		Type:        code.CodeType,
		Source:      code.Source,
		TimeLimit:   code.TimeLimit,
		MemoryLimit: code.MemoryLimit,
		Samples:     nil,
	}
	//2.2获取样例
	res, err := os.ReadDir("./file/sample/123/sample")
	samples := []*pb.Sample{}
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(res); i += 2 {
		s := new(pb.Sample)
		for j := i; j < i+2; j++ {
			if path.Ext(res[j].Name()) == ".in" {
				b, err := os.ReadFile("./file/sample/123/sample/" + res[j].Name())
				s.Input = string(b)
				if err != nil {
					return nil, err
				}
			} else if path.Ext(res[j].Name()) == ".out" {
				b, err := os.ReadFile("./file/sample/123/sample/" + res[j].Name())
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

	if err != nil {
		return nil, err
	}

	//3.创建结果结构体
	Subid, _ := uuid.Getuuid()
	Result := models.Result{
		SubmitID:   Subid,
		CodeId:     code.CodeID,
		QuestionID: code.QuestionID,
		UserID:     code.UserID,
		Time:       time.Now().String(),
	}
	Result.IfAC = true
	for i := range re.Results {
		if re.Results[i].Status != 1 {
			Result.IfAC = false
		}
		Result.Results = append(Result.Results, models.ResultOne{
			Status:   re.Results[i].Status,
			Memory:   strconv.FormatFloat(re.Results[i].Memory, 'f', -1, 32),
			RealTime: strconv.FormatFloat(re.Results[i].RealTime, 'f', -1, 32),
			CPUTime:  strconv.FormatFloat(re.Results[i].CpuTime, 'f', -1, 32),
			ErrorMsg: re.Results[i].Error,
		})
	}

	errCh1 := make(chan error)
	errCh2 := make(chan error)

	//4.插入结果goroutine
	go InsertReGoroutine(Result, code, errCh1)

	//5.插入解决代码goroutine
	go InsertSoGoroutine(code, errCh2)
	err = <-errCh1
	errNew := <-errCh2

	if err != nil {
		return nil, err
	}
	if errNew != nil {
		return nil, err
	}
	return &Result, nil
}

func InsertReGoroutine(Result models.Result, code models.Code, err chan error) {
	errD := dao.InsertStatus(Result)
	err <- errD
}

func InsertSoGoroutine(code models.Code, err chan error) {
	solution := models.Solution{
		CodeID:     code.CodeID,
		CodeType:   code.CodeType,
		Public:     code.Public,
		QuestionID: code.QuestionID,
		Source:     code.Source,
		Time:       time.Now().String(),
		UserID:     code.UserID,
	}
	errD := dao.InsertSolution(solution)
	err <- errD
}
