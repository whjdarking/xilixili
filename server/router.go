package server

import (
	"github.com/gin-gonic/gin"
	"os"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	gin.SetMode(os.Getenv("GIN_MODE"))

	// 启用cookie-based session
	r.Use(Session(os.Getenv("SESSION_SECRET")))
	//提取当前用户(若是能提取到，说明第一步的启用session已经成功了，也说明曾经是被下发过session cookie的)
	r.Use(CurrentUser())
	//配置跨域
	r.Use(Cors())

	// 路由
	v1 := r.Group("/api/v1")
	{
		//测试连通接口
		v1.POST("ping", Ping)

		// 用户注册
		v1.POST("user/register", UserRegister)
		// 用户登录
		v1.POST("user/login", UserLogin)

		//创建视频
		v1.POST("videos", CreateVideo)
		//通过id获取视频
		v1.GET("video/:id", ShowVideo)
		//获取视频列表
		v1.GET("videos", ListVideo)

		//周排行
		v1.GET("rank/weekly", WeeklyRank)

		// 需要登录保护的
		auth := v1.Group("")
		auth.Use(AuthRequired()) //确认当前有没有提取的用户
		{
			// User Routing
			auth.GET("user/me", UserMe)
			auth.DELETE("user/logout", UserLogout)
		}

		//前端未实现
		//更新视频
		v1.PUT("video/:id", UpdateVideo)
		//删除视频
		v1.DELETE("video/:id", DeleteVideo)

		//阿里云OSS相关权限
		v1.POST("upload/token", UploadToken)
	}
	return r
}
