package models

import (
	"gorm.io/gorm"
)

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
	SubmitID   string `json:"submit_id"`
	UserID     string `json:"user_id"`
	QuestionID string `json:"question_id"`
	Time       string `json:"time"`
	IfAC       int    `json:"if_ac"`
	CodeType   uint64 `json:"code_type"` // 代码类型
	Source     string `json:"source"`    // 源码
	Public     int64  `json:"public"`
	Results    string `json:"results"`
}

type ResultOfOneSample struct {
	Status   uint64 `json:"status"`
	Memory   string `json:"memory"`
	RealTime string `json:"real_time"`
	CPUTime  string `json:"cpu_time"`
	ErrorMsg string `json:"error_msg"`
}

// 非入库数据模型
type Submit struct {
	QuestionID  string `json:"question_id"` // 题目ID
	UserID      string `json:"user_id"`     // 用户ID
	CodeType    string `json:"code_type"`   // 代码类型
	Public      int64  `json:"public"`
	Source      string `json:"source"` // 源码
	Time        string `json:"time"`   // 提交时间
	TimeLimit   string `json:"time_limit"`
	MemoryLimit string `json:"memory_limit"`
}

type SubmitResult struct {
	SubmitID   string              `json:"submit_id"`
	UserID     string              `json:"user_id"`
	QuestionID string              `json:"question_id"`
	Time       string              `json:"time"`
	IfAC       int                 `json:"if_ac"`
	Results    []ResultOfOneSample `json:"results"`
}

// result
type ResultDetail struct {
	SubmitID   string              `json:"submit_id"`
	UserID     string              `json:"user_id"`
	QuestionID string              `json:"question_id"`
	Time       string              `json:"time"`
	IfAC       int                 `json:"if_ac"`
	CodeType   uint64              `json:"code_type"` // 代码类型
	Source     string              `json:"source"`    // 源码
	Public     int64               `json:"public"`
	Results    []ResultOfOneSample `json:"results"`
}

type AcSql struct {
	gorm.Model
	Ac
}

type Ac struct {
	UserID       string `json:"user_id"`
	AcQuestionID string `json:"ac_question_id" gorm:"unique;not null"`
	IfAC         int    `json:"if_ac"`
}
