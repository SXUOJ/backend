package models

type Code struct {
	CodeId     string `json:"code_id,omitempty"`
	CodeQId    string `json:"code_question_id,omitempty"`
	CodeUserId string `json:"code_userid,omitempty"`
	CodeType   string `json:"code_type,omitempty"`
	CodeTime   string `json:"code_time,omitempty"`
	CodeSource string `json:"code_source,omitempty"`
}
