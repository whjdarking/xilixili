package service

import (
	"Refactor_xilixili/model"
	"Refactor_xilixili/serializer"
)

type UpdateVideoService struct {
	Title string `json:"title" binding:"required,min=2,max=20"`
	Info  string `json:"info" binding:"required,min=0,max=200"`
}

func (service *UpdateVideoService) Update(id string) serializer.Response {
	var video model.Video
	//全局单例DB
	err := model.DB.Where("id = ?", id).Find(&video).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "不存在",
			Error: err.Error(), //打出错误信息
		}
	}

	video.Title = service.Title
	video.Info = service.Info
	err = model.DB.Save(&video).Error
	if err != nil {
		return serializer.Response{
			Code:  50002,
			Msg:   "视频保存失败",
			Error: err.Error(), //打出错误信息
		}
	}
	return serializer.Response{
		Code:  0,
		Data:  serializer.BuildVideo(video),
		Msg:   "",
		Error: "",
	}
}
