package service

import (
	"xilixili/model"
	"xilixili/serializer"
)

type ListVideoService struct {
}
func (service *ListVideoService) List() serializer.Response{
	var videos []model.Video
	//全局单例DB
	err := model.DB.Find(&videos).Error
	if err != nil {
		return serializer.Response{
			//如果要处理需要到api那里给c
			Code: 404,
			Msg:  "不存在",
			Error: err.Error(), //打出错误信息
		}
	}
	return serializer.Response{
		Code:  0,
		Data:  serializer.BuildVideos(videos),
		Msg:   "",
		Error: "",
	}
}