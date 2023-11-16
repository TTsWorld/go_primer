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

func Test0301(t *testing.T) {
	GetWeekEndTime(time.Now().UnixMilli(), 8*3600)
}

func TestUnix(t *testing.T) {
	tm := time.Now()
	utm := tm.UTC()
	//时间戳格式化为本地时间，需要加上本地时间和 utc0的时区差，或将时区设置为本地
	fmt.Println(tm)                                                //本地时间
	fmt.Println(tm.UTC())                                          //UTC时间
	fmt.Println(tm.Unix())                                         //本地时间时间戳
	fmt.Println(tm.UTC().Unix())                                   //UTC时间时间戳
	fmt.Println(tm.UTC().Add(time.Hour * 8).Format(time.DateTime)) //utc时间 -> 格式化为本地时间
	fmt.Println(tm.UTC().In(tm.Location()).Format(time.DateTime))  //utc时间 -> 格式化为本地时间

	fmt.Println("===============")
	tstr := time.Now().Format(time.DateTime)
	fmt.Println(tstr)                            //时间串
	fmt.Println(time.Parse(time.DateTime, tstr)) //默认解析为 utc
	fmt.Println(time.ParseInLocation(time.DateTime, tstr, tm.Location()))
	fmt.Println(time.ParseInLocation(time.DateTime, tstr, utm.Location()))

	//本地时间串获取 unix 时间戳, parse完后是对应时区的时间串对应的时间
	t1, _ := time.Parse(time.DateTime, tstr)
	t2, _ := time.ParseInLocation(time.DateTime, tstr, tm.Location())
	t3, _ := time.ParseInLocation(time.DateTime, tstr, utm.Location())
	fmt.Println(t1.Add(-time.Hour * 8).Unix()) //时间串
	fmt.Println(t2.Unix())                     //时间串
	fmt.Println(t3.Unix())                     //时间串

}
