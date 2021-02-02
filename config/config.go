package config

import (
	"Refactor_xilixili/cache"
	"Refactor_xilixili/model"
	"github.com/joho/godotenv"
	"os"
)

func Init() {
	//加载环境变量
	godotenv.Load()
	//连接mysql
	model.Database(os.Getenv("MYSQL_DSN"))
	//连接redis
	cache.Redis()
	//开始计时任务
	CronJob()
}
