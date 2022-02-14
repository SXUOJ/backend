package redis

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

// 声明一个全局的rdb变量
var rdb *redis.Client

// 初始化连接
func Init() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", viper.GetString("redis.host"), viper.GetInt("radis. port")),
		Password: viper.GetString("radis.password"),   // no password set
		DB:       viper.GetInt("radis.max_pool_size"), // use default DB
	})

	_, err = rdb.Ping().Result()
	return err
}

func Close() {
	rdb.Close()
}
