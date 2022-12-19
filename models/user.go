package models

type User struct {
	UserId    string `json:"user_id" db:"user_id"`
	Username  string `json:"username" db:"username"`
	Password  string `json:"password" db:"password"`
	Usergroup string `json:"usergroup"db:"usergroup"`
	Truename  string `json:"truename" db:"truename"`
	Email     string `json:"email" db:"email"`
	School    string `json:"school" db:"school"`
	Score     string `json:"score" db:"score"`
}

type UserSignUp struct {
	UserId   string `json:"user_id" db:"user_id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

type UserInMysql struct {
	UserId   string `json:"user_id" db:"user_id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}
