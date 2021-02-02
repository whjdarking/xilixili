package service

import (
	"Refactor_xilixili/model"
	"Refactor_xilixili/serializer"
)

type DeleteVideoService struct {
}

func (service *DeleteVideoService) Delete(id string) serializer.Response {
	var video model.Video
	//全局单例DB
	err := model.DB.Where("id = ?", id).Find(&video).Error
	if err != nil {
		return serializer.Response{
			//如果要处理code需要到api那里给c
			Code:  404,
			Msg:   "不存在",
			Error: err.Error(), //打出错误信息
		}
	}
	err = model.DB.Delete(&video).Error
	if err != nil {
		return serializer.Response{
			//如果要处理code需要到api那里给c
			Code:  50000,
			Msg:   "视频删除失败",
			Error: err.Error(), //打出错误信息
		}
	}
	return serializer.Response{}
}
