package cache

import (
	"fmt"
	"strconv"
)

const (
	WeeklyRankKey = "rank:weekly"
)

//view: video
//view: article?
func VideoViewKey(id uint) string {
	return fmt.Sprintf("view:video:%s", strconv.Itoa(int(id)))
}
