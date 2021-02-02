package service

import (
	"Refactor_xilixili/model"
	"Refactor_xilixili/serializer"
)

type CreateVideoService struct {
	Title  string `json:"title" binding:"required,min=2,max=50"`
	Info   string `json:"info" binding:"min=0,max=300"`
	URL    string `json:"url"`
	Avatar string `json:"avatar"`
	Belong string `json:"belong"`
}

func (service *CreateVideoService) Create() serializer.Response {
	println(service.Belong)
	video := model.Video{
		Title:  service.Title,
		Info:   service.Info,
		URL:    service.URL,
		Avatar: service.Avatar,
		Belong: service.Belong,
	}
	//全局单例DB
	err := model.DB.Create(&video).Error
	if err != nil {
		return serializer.Response{
			Code:  40001,
			Msg:   "视频创建保存失败",
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
