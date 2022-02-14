package logic

import (
	"github.com/pkg/errors"
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/jwt"
	"web_app/pkg/uuid"
)

func Register(user *models.UserSignUp) (string, error) {
	//生成userID
	userId, err := uuid.Getuuid()
	newuser := new(models.UserInMysql)
	newuser.UserId = userId
	newuser.Username = user.Username
	newuser.Password = user.Password
	//入库
	ok, err := mysql.Register(newuser)
	if !ok {
		return "nil", err
	}
	//生成token
	token, err := jwt.GenToken(user.Username)
	if err != nil {
		return "nil", err
	}
	return token, err
}

func Login(user *models.UserSignUp) (string, error) {
	//操作数据库校验登陆
	err := mysql.Login(user)
	if err != nil {
		if err.Error() == "密码错误" {
			return "", errors.New("密码错误")
		}
		return "", err
	}
	//生成token
	token, err := jwt.GenToken(user.Username)
	if err != nil {
		return "", err
	}
	return token, err
}

func GetUserInfo(username string) (userinfo *models.User, err error) {
	//传入username进行查库操作
	userinfo, err = mysql.GetUserInfo(username)
	if err != nil {
		return nil, err
	}
	return userinfo, nil
}

func PutUserInfo(user *models.User) error {
	//对新的UserInfo进行入库操作
	err := mysql.PutUserInfo(user)
	if err != nil {
		return err
	}
	return nil
}
