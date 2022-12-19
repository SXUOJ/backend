package models

// solution
type Solution struct {
	CodeID     string `json:"code_id"`   // 本次提交代码ID
	CodeType   string `json:"code_type"` // 代码类型
	Public     int64  `json:"public"`
	QuestionID string `json:"question_id"` // 题目ID
	Source     string `json:"source"`      // 源码
	Time       string `json:"time"`        // 提交时间
	UserID     string `json:"user_id"`     // 用户ID
}
