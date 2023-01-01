package mysql

import (
	"fmt"
	"web_app/models"

	"go.uber.org/zap"

	"github.com/spf13/viper"

	_ "github.com/go-sql-driver/mysql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Init() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		viper.GetString("mysql.user"), viper.GetString("mysql.password"),
		viper.GetString("mysql.host"), viper.GetInt("mysql.port"), viper.GetString("mysql.dbname"),
	)

	// 连接数据库
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("connect DB failed", zap.Error(err))
		return
	}

	// 绑定模型
	err = db.AutoMigrate(&models.UserSql{}, &models.QuestionSql{})
	if err != nil {
		fmt.Printf("Binding model failed", zap.Error(err))
		return
	}

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Printf("Set config failed", zap.Error(err))
		return
	}
	sqlDB.SetMaxOpenConns(viper.GetInt("mysql.max_open_conns"))
	sqlDB.SetMaxIdleConns(viper.GetInt("mysql.max_idle_conns"))
	return
}
