package logic

import (
	"github.com/SXUOJ/backend/dao"
	"github.com/SXUOJ/backend/models"
	"github.com/SXUOJ/backend/pb"
	"github.com/SXUOJ/backend/pkg/grpc"
	"github.com/SXUOJ/backend/pkg/uuid"
	"go.uber.org/zap"
	"strconv"
	"sync"
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
	addrs := []string{}
	var addr string
	for _, a := range addrs {
		_, err := grpc.Ping(a)
		if err == nil {
			addr = a
			break
		}
	}
	code.CodeID, _ = uuid.Getuuid()
	Quest := pb.JudgeRequest{
		SubmitId:    code.CodeID,
		Type:        code.CodeType,
		Source:      code.Source,
		TimeLimit:   code.TimeLimit,
		MemoryLimit: code.MemoryLimit,
		Samples:     nil,
	}
	re, err := grpc.Judge(addr, &Quest)

	if err != nil {
		return nil, err
	}

	Subid, _ := uuid.Getuuid()
	Result := models.Result{
		SubmitID: Subid,
		CodeId:   code.CodeID,
		UserID:   code.UserID,
		Time:     time.Now().String(),
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

	var wg sync.WaitGroup

	wg.Add(2)
	go InsertReGoroutine(Result, code, &err, &wg)

	var errNew error
	go InsertSoGoroutine(code, &errNew, &wg)

	wg.Wait()
	if err != nil {
		return nil, err
	}
	if errNew != nil {
		return nil, err
	}
	return &Result, nil
}

func InsertReGoroutine(Result models.Result, code models.Code, err *error, wg *sync.WaitGroup) {
	if Result.IfAC == true {
		errD := dao.InsertAc(models.AC{
			UserId: code.UserID,
			QueId:  code.QuestionID,
		})
		if errD != nil {
			err = &errD
			wg.Done()
			return
		}
	}
	errD := dao.InsertStatus(Result)
	err = &errD
	wg.Done()
}

func InsertSoGoroutine(code models.Code, err *error, wg *sync.WaitGroup) {
	solution := models.Solution{
		CodeID:     code.CodeID,
		CodeType:   code.CodeType,
		Public:     code.Public,
		QuestionID: code.QuestionID,
		Source:     code.Source,
		Time:       time.Now().String(),
		UserID:     code.UserID,
	}
	*err = dao.InsertSolution(solution)
	wg.Done()
}
