package models

// result
type Result struct {
	CPUTime  string `json:"cpu_time"`
	Memory   string `json:"memory"`
	RealTime string `json:"real_time"`
	Source   string `json:"source"`
	Status   string `json:"status"`
	SubmitID string `json:"submit_id"`
	Time     string `json:"time"`
	UserID   string `json:"user_id"`
}
