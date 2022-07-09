package timer

import (
	"time"
)

// GetNowTime 获取当前时间
func GetNowTime() time.Time {
	return time.Now()
}

// GetCalTime 返回传入的时间累加上指定的时间结果
func GetCalTime(cur time.Time, s string) (time.Time, error) {
	duration, err := time.ParseDuration(s)
	if err != nil {
		return time.Time{}, err
	}
	return cur.Add(duration), nil
}
