package mtime

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	now := time.Now()
	nowUnix := now.Unix()
	println(now.String())
	tm := time.Unix(nowUnix, 0)
	println(tm.Add(time.Minute * 10).String())

}

// 获取指定本地时分秒，日期是今日的时间戳
func TestTimeParse(t *testing.T) {
	value := "18:00:000"
	tm, err := time.Parse(HourMinuteSecondsLayout, value)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	localNow := time.Now()
	tt := time.Date(localNow.Year(), localNow.Month(), localNow.Day(), tm.Hour(), tm.Minute(), tm.Second(), 0, time.UTC)
	fmt.Printf("%+v", tt)

}

func Test0202(t *testing.T) {
	tm, _ := time.ParseInLocation(time.DateOnly, "2023-09-12", time.UTC)
	fmt.Println(GetZeroTime(tm))
	fmt.Println(GetZeroTime(time.Now().In(time.UTC)))
	fmt.Println(GetZeroTime(time.Now().In(time.UTC)).Sub(GetZeroTime(tm)).Hours() / 24)

}
