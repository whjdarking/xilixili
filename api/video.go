package api

import (
	"github.com/gin-gonic/gin"
	"xilixili/service"
)

//视频投稿
func CreateVideo(c *gin.Context) {
	service := service.CreateVideoService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

//视频详情
func ShowVideo(c *gin.Context) {
	service := service.ShowVideoService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Show(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func ListVideo(c *gin.Context) {
	service := service.ListVideoService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

//更新视频
func UpdateVideo(c *gin.Context) {
	service := service.UpdateVideoService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func DeleteVideo(c *gin.Context) {
	service := service.DeleteVideoService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Delete(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}