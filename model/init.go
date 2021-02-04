package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB

func Database(connString string) {

	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{})
	if err != nil {
		panic("数据库连接不成功," + connString)
	}
	//设置连接池
	sqlDB, err := db.DB()
	if err != nil {
		panic("设置通用数据库对象失败")
	}
	//空闲
	sqlDB.SetMaxIdleConns(30)
	//打开
	sqlDB.SetMaxOpenConns(50)
	//超时
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	DB = db

	DB.AutoMigrate(&User{}, &Video{})
}
