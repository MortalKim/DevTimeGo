package redis

import (
	"WakaTImeGo/config"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"time"
)

var log = logrus.New()

var (
	MyRedis *redis.Client
	ctx     = context.Background()
)

func Setup() {
	redisHost, redisPasswd, redisPort, redisDb := config.GetRedisConfig()

	MyRedis = redis.NewClient(&redis.Options{
		Addr:     redisHost + ":" + fmt.Sprint(redisPort),
		Password: redisPasswd, // no password set
		DB:       redisDb,     // use default DB
	})
	_, err := MyRedis.Ping(context.Background()).Result()
	if err != nil {
		log.Error("Redis connect ping failed, err:", err)
		return
	}
	log.Info("Redis connect succeeded")
	return
}

func SetRedis(key string, value string, t int64) bool {
	expire := time.Duration(t) * time.Second
	if err := MyRedis.Set(ctx, key, value, expire).Err(); err != nil {
		return false
	}
	return true
}

func GetRedis(key string) string {
	result, err := MyRedis.Get(ctx, key).Result()
	if err != nil {
		return ""
	}
	return result
}

func DelRedis(key string) bool {
	_, err := MyRedis.Del(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func ExpireRedis(key string, t int64) bool {
	expire := time.Duration(t) * time.Second
	if err := MyRedis.Expire(ctx, key, expire).Err(); err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
