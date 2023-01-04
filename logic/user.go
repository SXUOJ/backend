package logic

import (
	"github.com/SXUOJ/backend/dao"
	"github.com/SXUOJ/backend/models"
	"github.com/SXUOJ/backend/pkg/jwt"
	"github.com/SXUOJ/backend/pkg/uuid"
	"github.com/pkg/errors"
)

func Register(user *models.UserSignUp) (string, error) {
	//生成userID
	userId, err := uuid.Getuuid()
	newuser := new(models.UserInMysql)
	newuser.UserId = userId
	newuser.Username = user.Username
	newuser.Password = user.Password
	//入库
	ok, err := dao.Register(newuser)
	if !ok {
		return "nil", err
	}
	//生成token
	token, err := jwt.GenToken(userId, user.Username)
	if err != nil {
		return "nil", err
	}
	return token, nil
}

func Login(user *models.UserSignUp) (string, error) {
	//操作数据库校验登陆
	userid, err := dao.Login(user)
	if err != nil {
		if err.Error() == "密码错误" {
			return "", errors.New("密码错误")
		}
		return "", err
	}
	//生成token
	token, err := jwt.GenToken(userid, user.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func GetUserInfo(username string) (userinfo *models.User, err error) {
	//传入username进行查库操作
	userinfo, err = dao.GetUserInfo(username)
	if err != nil {
		return nil, err
	}
	return userinfo, nil
}

func PutUserInfo(user *models.User) error {
	//对新的UserInfo进行入库操作
	err := dao.PutUserInfo(user)
	if err != nil {
		return err
	}
	return nil
}
