package models

import "gorm.io/gorm"

type ResultSql struct {
	gorm.Model
	Result
}

// result
type Result struct {
	CPUTime  string `json:"cpu_time"`
	Memory   string `json:"memory"`
	RealTime string `json:"real_time"`
	CodeId   string `json:"code_id"`
	Status   string `json:"status"`
	SubmitID string `json:"submit_id"`
	Time     string `json:"time"`
	UserID   string `json:"user_id"`
}
