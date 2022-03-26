package models

type Code struct {
	CodeId     string `json:"submit_id,omitempty" db:"codeId"`
	CodeQId    string `json:"problem_id,omitempty" db:"codeQId"`
	CodeUserId string `json:"code_userid,omitempty" db:"codeUserId"`
	CodeQType  string `json:"problem_type,omitempty" db:"codeQType"`
	CodeType   string `json:"code_type,omitempty" db:"codeType"`
	CodeTime   string `json:"time_limit,omitempty" db:"codeTime"`
	CodeMemory string `json:"memory_limit" db:"codeMemory"`
	CodeSource string `json:"code_source_path,omitempty" db:"codeSource"`
	CodeState  string `json:"code_state,omitempty" db:"codeState"`
}
