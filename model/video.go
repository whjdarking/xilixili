package model

import (
	"github.com/jinzhu/gorm"
	"strconv"
	"xilixili/cache"
)

type Video struct{
	gorm.Model
	Title string
	Info string `sql:"type:text"`
	URL string
	Avatar string
}
func (video *Video) View() uint64 {
	countStr, _ := cache.RedisClient.Get(cache.VideoViewKey(video.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

func (video *Video) AddView() {
	//增加视频点击数
	cache.RedisClient.Incr(cache.VideoViewKey(video.ID))

	cache.RedisClient.ZIncrBy(cache.WeeklyRankKey, 1, strconv.Itoa(int(video.ID)))
}

