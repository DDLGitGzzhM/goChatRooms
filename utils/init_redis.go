package utils

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var Red *redis.Client

// InitRedis
// 初始化Redis
func InitRedis() {
	Red = redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.addr"),
		Password:     viper.GetString("redis.password"),
		DB:           viper.GetInt("redis.DB"),
		PoolTimeout:  viper.GetDuration("redis.poolSize"),
		MinIdleConns: viper.GetInt("redis.minIdleConn"),
	})

	ctx := context.Background()
	pong, err := Red.Ping(ctx).Result()
	if err != nil {
		fmt.Println("init  redis....", err)
	} else {
		fmt.Println("Redis inited ... ", pong)
	}
}
