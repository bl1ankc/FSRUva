package utils

import "time"

func Now() int64 {
	return time.Now().UnixNano() / 1e6
}

func NowString() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
