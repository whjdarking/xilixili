package cache

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"os"
	"strconv"
)

var RedisClient *redis.Client
var ctx = context.TODO()

func Redis() {
	db, _ := strconv.ParseUint(os.Getenv("REDIS_DB"), 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PW"),
		DB:       int(db),
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		panic("连接Redis不成功," + os.Getenv("REDIS_ADDR"))
	}
	RedisClient = client
}

const (
	WeeklyRankKey = "rank:weekly"
)

//通过id生成key
//格式为view:video:id -> 得到value
func VideoViewKey(id uint) string {
	return fmt.Sprintf("view:video:%s", strconv.Itoa(int(id)))
}
