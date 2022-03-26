package models

type Question struct {
	Title       string       `json:"title" db:"title"`
	Context     *Context     `json:"context"`
	Information *Information `json:"information"`
	Statistic   *Statistic   `json:"statistic"`
}

type Context struct {
	Description  string `json:"description" db:"description"`
	Input        string `json:"input" db:"input"`
	Output       string `json:"output" db:"output"`
	SampleInput  string `json:"sample_input" db:"sampleInput"`
	SampleOutput string `json:"sample_output" db:"sampleOutput"`
	Source       string `json:"source" db:"source"`
}

type Information struct {
	Id        string `json:"id" db:"id"`
	TimeLimit string `json:"time_limit" db:"timeLimit"`
	MemLimit  string `json:"mem_limit" db:"memLimit"`
	IoMode    string `json:"io_mode" db:"ioMode"`
	CreateBy  string `json:"create_by" db:"createBy"`
	Level     string `json:"level" db:"level"`
	Tags      string `json:"tags" db:"tags"`
}

type Statistic struct {
	Ac string `json:"ac" db:"ac"`
	Wa string `json:"wa" db:"wa"`
}

// MySqlAll 问题类型的mysql模型
type MySqlAll struct {
	Title        string `json:"title" db:"title"`
	Description  string `json:"description" db:"description"`
	Input        string `json:"input" db:"input"`
	Output       string `json:"output" db:"output"`
	SampleOutput string `json:"sample_output" db:"sampleOutput"`
	SampleInput  string `json:"sample_input" db:"sampleInput"`
	Source       string `json:"source" db:"source"`
	Id           string `json:"id" db:"id"`
	TimeLimit    string `json:"time_limit" db:"timeLimit"`
	MemLimit     string `json:"mem_limit" db:"memLimit"`
	IoMode       string `json:"io_mode" db:"ioMode"`
	CreateBy     string `json:"create_by" db:"createBy"`
	Level        string `json:"level" db:"level"`
	Tags         string `json:"tags" db:"tags"`
	Ac           string `json:"ac" db:"ac"`
	Wa           string `json:"wa" db:"wa"`
}

type QueList struct {
	Title string `json:"title" db:"title"`
	Id    string `json:"id" db:"id"`
	QueId string `json:"que_id" db:"que_id"`
	Tags  string `json:"tags" db:"tags"`
}
