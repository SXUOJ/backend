package logic

import (
	"github.com/SXUOJ/backend/dao"
	"github.com/SXUOJ/backend/models"
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
