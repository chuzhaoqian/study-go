package timetools

import (
	"time"
)

func Now() time.Time {
	return time.Now()
}

func NowDateTime() string {
	now := Now()
	return now.Format("2006-01-02 15:04:05")
}
