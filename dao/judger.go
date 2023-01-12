package dao

import "github.com/SXUOJ/backend/models"

func InsertJudger(addr models.JudgerAddr) error {
	return db.Create(&models.JudgerAddrSql{JudgerAddr: addr}).Error
}

func GetJudger() (addrs []*models.JudgerAddr, err error) {
	return nil, nil
}
