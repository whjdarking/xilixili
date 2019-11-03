package serializer

import "xilixili/model"

type Video struct {
	ID        uint   `json:"id"`
	Title  string `json:"title"`
	Info   string `json:"info"`
	URL   string `json:"url"`
	CreatedAt int64  `json:"created_at"`
}

// BuildVideo 序列化单个视频
func BuildVideo(item model.Video) Video {
	return Video{
		ID:        item.ID,
		Title:  item.Title,
		Info:  item.Info,
		URL:  item.URL,
		CreatedAt: item.CreatedAt.Unix(),
	}
}

// BuildVideos 序列化视频list
func BuildVideos(items []model.Video) (videos []Video) {
	for _,item := range items{
		video := BuildVideo(item)
		videos = append(videos, video)
	}
	return videos
}