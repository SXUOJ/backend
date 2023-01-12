package logic

import (
	"github.com/SXUOJ/backend/dao"
	"github.com/SXUOJ/backend/models"
)

func CreateJudger(addr models.JudgerAddr) error {
	return dao.InsertJudger(addr)
}

func GetJudgerList() ([]*models.JudgerAddr, error) {
	return dao.GetJudger()
}
