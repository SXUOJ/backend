package logic

import (
	"go.uber.org/zap"
	"web_app/dao/mysql"
	"web_app/models"
)

func GetQuestionDetail(Qid string) (que *models.Question, err error) {
	//查库
	que, err = mysql.GetQuestionDetail(Qid)
	if err != nil {
		zap.L().Error("GetQuestionDetail(Qid string) err...", zap.Error(err))
		return nil, err
	}
	return que, nil
}

func GetQuestionList(page int, amount int) (data []*models.Question, err error) {
	//查库
	data, err = mysql.GetQuestionList(page, amount)
	if err != nil {
		zap.L().Error("mysql.GetQuestionList(page, amount) err ", zap.Error(err))
		return nil, err
	}
	return data, nil
}

func CreateQuestion(que models.Question) error {
	return nil
}

func ChangeQuestion(qid string) error {
	return nil
}

func DelQuestion(qid string) error {
	return nil
}
