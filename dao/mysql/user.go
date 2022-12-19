package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"web_app/models"
)

const secret = "jiaomaster"

func CheckUserExist(UserName string) (bool, error) {
	//1.根据用户名与库中用户名匹配
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err := db.Get(&count, sqlStr, UserName); err != nil {
		return false, err
	}
	return count > 0, nil
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
	sqlStr := `insert into user (user_id,username,password) values(?,?,?)`
	_, err = db.Exec(sqlStr, user.UserId, user.Username, user.Password)
	if err != nil {
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
	sqlStr := "select user_id,password from user where username=?"
	var Rpassword string
	err = db.Get(&Rpassword, sqlStr, user.Username)
	if err != nil {
		zap.L().Error("db.Get(password) err...", zap.Error(err))
		return "", err
	}
	//解密当前密码
	Npassword := encryptPassword(user.Password)
	if err != nil {
		zap.L().Error("uuid.ParseUuid err...", zap.Error(err))
		return "", err
	}
	if Rpassword != Npassword {
		zap.L().Debug("Login failed... password err...")
		return "", errors.New("密码错误")
	}
	//校验成功
	return user.UserId, nil
}

// 获取用户信息
func GetUserInfo(username string) (userinfo *models.User, err error) {
	userinfo = new(models.User)
	sqlStr := `select user_id, username, truename, email, school, score from user where username = ?`
	err = db.Get(userinfo, sqlStr, username)
	if err != nil {
		zap.L().Error("db.Get(userinfo,sqlStr,username) err", zap.Error(err))
		return nil, err
	}
	return userinfo, nil
}

// 修改用户信息
func PutUserInfo(user *models.User) error {
	sqlStr := "update user set truename=?,email=?,school=? where username = ?"
	ret, err := db.Exec(sqlStr, user.Truename, user.Email, user.School, user.Username)
	if err != nil {
		zap.L().Error("update failed, err:", zap.Error(err))
		return err
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		zap.L().Error("get RowsAffected failed, err:", zap.Error(err))
		return err
	}
	zap.L().Debug("update success...", zap.Int64(" affected rows:", n))
	return nil
}
