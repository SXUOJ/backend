package models

type Question struct {
	title   string   `json:"title" db:"title"`
	context *Context `json:"context"`
}

type Context struct {
	description string   `json:"description"`
	input       string   `json:"input"`
	output      string   `json:"output"`
	sampleInput []string `json:"sampleInput"`
	source      string   `json:"source"`
}

type Information struct {
	id        string `json:"id"`
	timeLimit string `json:"timeLimit"`
	memLimit  string `json:"memLimit"`
	ioMode    string `json:"ioMode"`
	createBy  string `json:"createBy"`
	level     string `json:"level"`
	tags      string `json:"tags"`
}
