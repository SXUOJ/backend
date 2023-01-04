package logic

import (
	"strconv"

	"github.com/SXUOJ/backend/dao"
	"github.com/SXUOJ/backend/models"
)

func GetStatusList(qid string, uid string, amount string, page string) ([]*models.Result, error) {
	pageInt, err := strconv.Atoi(page)
	amountInt, err := strconv.Atoi(amount)
	re, err := dao.GetStatusListByQid(qid, uid, amountInt, pageInt)
	return re, err
}

func GetStatusDetail(sid string) (models.Result, error) {
	re, err := dao.GetStatusByAdmitId(sid)
	return re, err
}
