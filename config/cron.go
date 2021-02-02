package config

import (
	"Refactor_xilixili/cache"
	"context"
	"fmt"
	"github.com/robfig/cron/v3"
)

var ctx = context.TODO()
var c *cron.Cron

func CronJob() {

	if c == nil {
		c = cron.New()
	}

	c.AddFunc("@weekly", func() { cache.RedisClient.Del(ctx, "rank:weekly") }) //invoked
	c.Start()
	fmt.Println("cron start......")
}
