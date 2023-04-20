package main

import (
	"fmt"
	"github.com/SXUOJ/backend/dao"
	"github.com/SXUOJ/backend/logger"
	"github.com/SXUOJ/backend/routes"
	"github.com/SXUOJ/backend/settings"
	"github.com/spf13/viper"

	"go.uber.org/zap"
)

func main() {
	//1、加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("settings.Init() err:%v", err)
		return
	}
	//2.初始化日志
	if err := logger.Init(); err != nil {
		fmt.Printf("logger.Init() err:%v", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("zap init sucessed...")
	//3.SQL
	if err := dao.DBInit(); err != nil {
		fmt.Printf("dao.Init() err:%v", err)
		return
	}
	//defer mysql.Close()
	//4.注册路由
	r := routes.Setup()
	//5.启动服务
	err := r.Run(fmt.Sprintf(":%d", viper.GetInt64("app.port")))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
}
