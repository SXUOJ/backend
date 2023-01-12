package dao

import (
	"fmt"

	"github.com/SXUOJ/backend/models"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func DBInit() (err error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		viper.GetString("db.host"), viper.GetString("db.port"),
		viper.GetString("db.user"), viper.GetString("db.password"),
		viper.GetString("db.db_name"), viper.GetString("db.ssl_mode"),
		viper.GetString("db.time_zone"),
	)

	// 连接数据库
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("connect postgres failed: %v", zap.Error(err))
		return
	}

	// 绑定模型
	err = db.AutoMigrate(
		&models.UserSql{},
		&models.QuestionSql{},
		&models.JudgerAddrSql{},
		&models.ResultSql{},
	)
	if err != nil {
		fmt.Printf("Binding model failed: %v", zap.Error(err))
		return
	}

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Printf("Set config failed: %v", zap.Error(err))
		return
	}
	sqlDB.SetMaxOpenConns(viper.GetInt("postgres.max_open_conns"))
	sqlDB.SetMaxIdleConns(viper.GetInt("postgres.max_idle_conns"))
	return
}
