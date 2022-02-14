package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

func Init() (err error) {
	viper.SetConfigFile("config.yaml") // 指定配置文件
	viper.AddConfigPath("./")          // 指定查找配置文件的路径
	err = viper.ReadInConfig()         // 读取配置信息
	if err != nil {
		// 读取配置信息失败
		fmt.Printf("viper.ReadInConfig err: %v", err)
		return err
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// 监控配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了")
	})
	return
	//r := gin.Default()
	//// 访问/version的返回值会随配置文件的变化而变化
	//r.GET("/version", func(c *gin.Context) {
	//	c.String(http.StatusOK, viper.GetString("version"))
	//})
	//
	//if err := r.Run(
	//	fmt.Sprintf(":%d", viper.GetInt("port"))); err != nil {
	//	panic(err)
	//}
}
