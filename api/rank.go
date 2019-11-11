package api

import (
	"github.com/gin-gonic/gin"
	"xilixili/service"
)

func WeeklyRank(c *gin.Context) {
	service := service.WeeklyRankService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Get()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}