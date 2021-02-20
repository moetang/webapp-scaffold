package utils

import "time"

func UnixTime(t time.Time) int64 {
	return t.Unix()*1000 + t.UnixNano()/1000/1000%1000
}
