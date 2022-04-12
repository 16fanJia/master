package dbops

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var RDB *redis.Client

// InitRedis 初始化redis
func InitRedis() {
	ctx := context.Background()

	RDB = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Password: viper.GetString("redis.password"),
		DB:       0,
	})

	//ping redis 是否联通
	_, err := RDB.Ping(ctx).Result()
	if err != nil {
		panic("redis connection failed err:" + err.Error())
		return
	}

}
