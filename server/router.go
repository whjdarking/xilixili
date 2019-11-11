package server

import (
	"os"
	"xilixili/api"
	"xilixili/middleware"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser()) //提取当前用户

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("ping", api.Ping)

		// 用户登录
		v1.POST("user/register", api.UserRegister)

		// 用户登录
		v1.POST("user/login", api.UserLogin)

		//创建视频
        v1.POST("videos", api.CreateVideo)
		//通过id获取视频
		v1.GET("video/:id", api.ShowVideo)
		//获取视频列表
		v1.GET("videos", api.ListVideo)
		//更新视频
		v1.PUT("video/:id", api.UpdateVideo)
		//删除视频
		v1.DELETE("video/:id", api.DeleteVideo)

		//周排行
		v1.GET("rank/weekly", api.WeeklyRank)

		// 需要登录保护的
		auth := v1.Group("")
		auth.Use(middleware.AuthRequired()) //确认刚才提取的用户是存在的
		{
			// User Routing
			auth.GET("user/me", api.UserMe)
			auth.DELETE("user/logout", api.UserLogout)
		}
	}
	return r
}
