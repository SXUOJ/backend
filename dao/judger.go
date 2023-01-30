package dao

import "github.com/SXUOJ/backend/models"

func InsertJudger(addr models.JudgerAddr) error {
	return db.Create(&models.JudgerAddrSql{JudgerAddr: addr}).Error
}

func GetJudger() ([]*models.JudgerAddr, error) {
	var (
		addrSqls []*models.JudgerAddrSql
		addrs    []*models.JudgerAddr
	)

	if err := db.Find(&addrSqls).Error; err != nil {
		return nil, err
	}

	for i := range addrSqls {
		addrs = append(addrs, &addrSqls[i].JudgerAddr)
	}

	return addrs, nil
}
