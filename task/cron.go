package task

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"xilixili/cache"
)

var c *cron.Cron
func CronJob() {
	if c == nil {
		c = cron.New()
	}

	c.AddFunc("@weekly", func() {cache.RedisClient.Del("rank:weekly") }) //invoked
	c.Start()
	fmt.Println("cron start......")
}



//func ResetWeeklyRank() error {
//	return cache.RedisClient.Del("rank:weekly").Err()
//}