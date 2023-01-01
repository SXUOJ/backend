package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"web_app/models"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

const secret = "jiaomaster"

func CheckUserExist(UserName string) (bool, error) {
	// 根据用户名与库中用户名匹配
	result := db.Where("usernmae = ?", UserName).Find(models.User{})
	if result.RowsAffected > 0 || result.Error != nil {
		return true, result.Error
	}
	return false, nil
}

// 对密码加密
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

func Register(user *models.UserInMysql) (ok bool, err error) {
	//1.检查账号是否重复
	username := user.Username
	ok, err = CheckUserExist(username)
	if ok {
		zap.L().Debug("CheckUserExist(userId) fail...", zap.String("DeBug", "账号存在"))
		return false, errors.New("账号存在")
	}
	if !ok {
		zap.L().Debug("CheckUserExist(userId) !ok...", zap.Error(err))
	}
	//2.密码加密
	userPassword := encryptPassword(user.Password)
	user.Password = userPassword
	//3.数据入库
	if err := db.Create(&models.UserSql{User: models.User{UserId: user.UserId, Username: user.Username, Password: user.Password}}).Error; err != nil {
		return false, err
	}
	return ok, err
}

func Login(user *models.UserSignUp) (string, error) {
	//校验账号
	username := user.Username
	ok, err := CheckUserExist(username)
	if !ok {
		zap.L().Debug("CheckUserExist(userId) !ok...", zap.Error(err))
		return "", errors.New("账号不存在")
	}
	//校验密码
	var userSql models.UserSql
	if err := db.Model(&models.UserSql{}).Where("username = ?", user.Username).Find(&userSql).Error; err != nil {
		zap.L().Error("db.Get(password) err...", zap.Error(err))
		return "", err
	}
	//加密提交的密码
	Npassword := encryptPassword(user.Password)
	if err != nil {
		zap.L().Error("uuid.ParseUuid err...", zap.Error(err))
		return "", err
	}
	if userSql.Password != Npassword {
		zap.L().Debug("Login failed... password err...")
		return "", errors.New("密码错误")
	}
	//校验成功
	return user.UserId, nil
}

// 获取用户信息
func GetUserInfo(username string) (*models.User, error) {
	var userSql models.UserSql
	if err := db.Model(&models.UserSql{}).Where("username = ?", username).Find(&userSql).Error; err != nil {
		zap.L().Error("db.Get(userinfo,sqlStr,username) err", zap.Error(err))
		return nil, err
	}
	return &userSql.User, nil
}

// 修改用户信息
func PutUserInfo(user *models.User) error {
	if err := db.Model(&models.UserSql{}).Where("username = ?", user.Username).Updates(models.UserSql{
		User: *user,
	}).Error; err != nil {
		zap.L().Error("update failed, err:", zap.Error(err))
		return err
	}

	zap.L().Debug("update success...")
	return nil
}
