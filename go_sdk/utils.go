package go_sdk

import (
	"fmt"
	"time"
)

func TimeStrToStamp(timeString string) time.Time {
	layout := TimeLaout
	t, err := time.Parse(layout, timeString)
	if err != nil {
		fmt.Println("解析失败:", err)
		return t
	}
	return t
}
