package model

import (
	"Refactor_xilixili/cache"
	"context"
	"gorm.io/gorm"
	"os"
	"strconv"
)

type Video struct {
	gorm.Model
	Title  string
	Info   string `sql:"type:text"`
	URL    string
	Avatar string
	Belong string
}

var ctx = context.TODO()

// AvatarURL 封面地址
func (video *Video) AvatarURL() string {
	//println("3")
	if len(video.Avatar) < 1 {
		return ""
	}
	signedGetURL := "https://xilixili." + os.Getenv("OSS_END_POINT") + "/" + video.Avatar
	//if len(video.URL) < 2 {
	//	return ""
	//}
	//client, _ := oss.New(os.Getenv("OSS_END_POINT"), os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"))
	//bucket, _ := client.Bucket(os.Getenv("OSS_BUCKET"))
	//signedGetURL, _ := bucket.SignURL(video.Avatar, oss.HTTPGet, 600)
	//println("4")
	return signedGetURL
}

// VideoURL 视频地址
func (video *Video) VideoURL() string {
	if len(video.URL) < 1 {
		return ""
	}
	signedGetURL := "https://xilixili." + os.Getenv("OSS_END_POINT") + "/" + video.URL
	//client, _ := oss.New(os.Getenv("OSS_END_POINT"), os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"))
	//bucket, _ := client.Bucket(os.Getenv("OSS_BUCKET"))
	//signedGetURL, _ := bucket.SignURL(video.URL, oss.HTTPGet, 600)
	return signedGetURL
}

func (video *Video) View() uint64 {
	count, _ := cache.RedisClient.Get(ctx, cache.VideoViewKey(video.ID)).Uint64()
	//count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

func (video *Video) AddView() {
	//增加视频点击数
	cache.RedisClient.Incr(ctx, cache.VideoViewKey(video.ID))
	cache.RedisClient.ZIncrBy(ctx, cache.WeeklyRankKey, 1, strconv.Itoa(int(video.ID)))
}
