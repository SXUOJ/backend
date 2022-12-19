package logic

import "web_app/models"

func GetStatusList(qid string, uid string, amount string, page string) ([]*models.Result, error) {
	return nil, nil
}

func GetStatusDetail(sid string) (models.Result, error) {
	return *new(models.Result), nil
}
