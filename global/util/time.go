package util

import (
	"fmt"
	"time"
)

var IntervalLabels = []string{"秒", "分钟", "小时", "天", "月", "年"}

func ElapsedTimeString(timestamp int64) string {
	elapsed := time.Now().Unix() - timestamp
	if elapsed < 0 {
		elapsed = -elapsed
	}

	intervals := []int64{
		elapsed,
		elapsed / 60,
		elapsed / 3600,
		elapsed / 86400,    //60 * 60 * 24
		elapsed / 2592000,  //60 * 60 * 24 * 30
		elapsed / 31104000, //60 * 60 * 24 * 30 * 12
	}

	for i := 5; i >= 0; i-- {
		interval := intervals[i]
		if interval > 0 {
			return fmt.Sprintf("%d%s", interval, IntervalLabels[i])
		}
	}

	return ""
}
