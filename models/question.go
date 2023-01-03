package models

import "gorm.io/gorm"

type QuestionSql struct {
	*gorm.Model
	Question
}

// Question
type Question struct {
	Context     Empty       `json:"context"`
	Information Information `json:"information"`
	Limit       Limit       `json:"limit"`
	Title       string      `json:"title"` // 标题
}

type Empty struct {
	Description  string   `json:"description"` // 描述
	ImgPath      string   `json:"img_path"`
	Input        string   `json:"input"`         // 输入描述
	InputSample  []string `json:"input_sample"`  // 输入样例
	Output       string   `json:"output"`        // 输出描述
	OutputSample []string `json:"output_sample"` // 输出样例
	Source       string   `json:"source"`        // 来源
}

type Information struct {
	Creator    string `json:"creator"`     // 创建者
	Level      string `json:"level"`       // 等级
	QuestionID string `json:"question_id"` // 问题ID
	Tags       string `json:"tags"`        // 标签
}

type Limit struct {
	MemLimit  string `json:"mem_limit"`  // 内存限制
	TimeLimit string `json:"time_limit"` // 时间限制
}
