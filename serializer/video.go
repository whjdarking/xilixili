package serializer

import "Refactor_xilixili/model"

type Video struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Info      string `json:"info"`
	URL       string `json:"url"`
	Avatar    string `json:"avatar"`
	Belong    string `json:"belong"`
	View      uint64 `json:"view"`
	CreatedAt int64  `json:"created_at"`
}

// BuildVideo 序列化单个视频
func BuildVideo(item model.Video) Video {
	return Video{
		ID:        item.ID,
		Title:     item.Title,
		Info:      item.Info,
		URL:       item.URL,
		Avatar:    item.AvatarURL(),
		Belong:    item.Belong,
		View:      item.View(),
		CreatedAt: item.CreatedAt.Unix(),
	}
}

// BuildVideos 序列化视频list
func BuildVideos(items []model.Video) (videos []Video) {

	for _, item := range items {
		//println("building...")
		video := BuildVideo(item)
		videos = append(videos, video)
	}
	return videos
}
