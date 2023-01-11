package models

import "gorm.io/gorm"

type JudgerAddrSql struct {
	gorm.Model
	JudgerAddr
}

type JudgerAddr struct {
	Name string `json:"name" gorm:"name;unique"`
	Addr string `json:"addr"`
}
