package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
)

func GetStatusList(qid string, uid string, amount string, page string) ([]*models.Result, error) {
	re, err := mysql.GetStatusListByQid(qid, uid, amount, page)
	return re, err
}

func GetStatusDetail(sid string) (models.Result, error) {
	re, err := mysql.GetStatusByAdmitId(sid)
	return re, err
}
