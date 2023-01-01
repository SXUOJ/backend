package models

import "gorm.io/gorm"

type UserSql struct {
	*gorm.Model
	User
}

type User struct {
	UserId    string `json:"user_id" gorm:"user_id"`
	Username  string `json:"username" gorm:"username"`
	Password  string `json:"password" gorm:"password"`
	Usergroup string `json:"usergroup"gorm:"usergroup"`
	Truename  string `json:"truename" gorm:"truename"`
	Email     string `json:"email,omitempty" gorm:"email"`
	School    string `json:"school,omitempty" gorm:"school"`
	Score     string `json:"score,omitempty" gorm:"score"`
}

type UserSignUp struct {
	UserId   string `json:"user_id" gorm:"user_id"`
	Username string `json:"username" gorm:"username"`
	Password string `json:"password" gorm:"password"`
}

type UserInMysql struct {
	UserId   string `json:"user_id" gorm:"user_id"`
	Username string `json:"username" gorm:"username"`
	Password string `json:"password" gorm:"password"`
}
