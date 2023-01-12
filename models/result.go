package models

import "gorm.io/gorm"

type ResultSql struct {
	gorm.Model
	Result
}

type Status uint64

const (
	Success           Status = 0
	Accepted          Status = 1
	WrongAnswer       Status = 2
	CompileError      Status = 3
	RuntimeError      Status = 4
	TimeLimitExceed   Status = 5
	MemoryLimitExceed Status = 6
	OutputLimitExceed Status = 7
	PresentationError Status = 8
	SystemError       Status = 9
	UnkownError       Status = 10
)

// result
type Result struct {
	SubmitID string `json:"submit_id"`
	CodeId   string `json:"code_id"`
	UserID   string `json:"user_id"`
	Time     string `json:"time"`
	IfAC     bool   `json:"if_ac"`
	Results  []ResultOne
}

type ResultOne struct {
	Status   uint64 `json:"status"`
	Memory   string `json:"memory"`
	RealTime string `json:"real_time"`
	CPUTime  string `json:"cpu_time"`
	ErrorMsg string `json:"error_msg"`
}
