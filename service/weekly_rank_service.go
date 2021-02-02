package service

import (
	"Refactor_xilixili/cache"
	"Refactor_xilixili/model"
	"Refactor_xilixili/serializer"
	"context"
	"fmt"
	"strings"
)

type WeeklyRankService struct {
}

var ctx = context.TODO()

func (service *WeeklyRankService) Get() serializer.Response {
	var videos []model.Video

	vids, _ := cache.RedisClient.ZRevRange(ctx, cache.WeeklyRankKey, 0, 4).Result() //取出前5

	if len(vids) > 1 {
		order := fmt.Sprintf("FIELD(id, %s)", strings.Join(vids, ","))
		err := model.DB.Where("id in (?)", vids).Order(order).Find(&videos).Error

		if err != nil {
			return serializer.Response{
				Code:  50000,
				Msg:   "redis连接错误",
				Error: err.Error(),
			}
		}
	}

	return serializer.Response{
		Data: serializer.BuildVideos(videos),
	}
}
