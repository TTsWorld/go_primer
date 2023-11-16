package mtime

import (
	"time"
)

// seconds
const (
	MinuteSeconds    = 60
	HourSeconds      = MinuteSeconds * 60
	DaySeconds       = HourSeconds * 24
	MaxUnixTimeStamp = 2147483647
)

// milliSeconds
const (
	MinuteMilliSeconds = MinuteSeconds * 1000
	HourMillSeconds    = MinuteMilliSeconds * 60
	DayMillSeconds     = HourMillSeconds * 24
)

func DateTime64StringFromTimestampMilli(now int64, tz_seconds int) string {
	localtime := now + int64(tz_seconds)*1000
	tm := time.Unix(0, localtime*int64(time.Millisecond))
	return tm.In(time.UTC).Format("2006-01-02 15:04:05.000")
}

func DateStringFromTimestampMilli(now int64, tz_seconds int) string {
	localtime := now/1000 + int64(tz_seconds)
	return time.Unix(localtime, 0).In(time.UTC).Format("2006-01-02")
}

func YesterdayStringFromTimestampMilli(now int64, tz_seconds int) string {
	localtime := now/1000 + int64(tz_seconds)
	return time.Unix(localtime, 0).In(time.UTC).AddDate(0, 0, -1).Format("2006-01-02")
}

func DateStringYMDHMSFromTimestampMilli(now int64, tz_seconds int) string {
	localtime := now/1000 + int64(tz_seconds)
	return time.Unix(localtime, 0).In(time.UTC).Format("2006-01-02 15:04:05")
}

func CurrentHourStringFromTimestampMilli(now int64, tz_seconds int) string {
	localtime := now/1000 + int64(tz_seconds)
	tm := time.Unix(localtime, 0).In(time.UTC)

	hourStart := time.Date(tm.Year(), tm.Month(), tm.Day(), tm.Hour(), 0, 0, 0, time.Local)
	dt := hourStart.Format("2006-01-02 15:04:05")
	return dt
}

func LastMondayStringFromTimestampMilli(now int64, tz_seconds int) string {
	localtime := now/1000 + int64(tz_seconds)
	tm := time.Unix(localtime, 0).In(time.UTC)

	offset := (tm.Weekday() + 6) % 7

	l1time := tm.AddDate(0, 0, -int(offset)-7)

	dt := l1time.Format("2006-01-02")
	return dt
}

func MondayStringFromTimestampMilli(now int64, tz_seconds int) string {
	localtime := now/1000 + int64(tz_seconds)
	tm := time.Unix(localtime, 0).In(time.UTC)

	offset := int(time.Monday - tm.Weekday())
	if offset > 0 {
		offset = -6
	}

	l1time := tm.AddDate(0, 0, offset)

	dt := l1time.Format("2006-01-02")
	return dt
}

func DateTimeStringToTimestampMilli(dt string) int64 {
	tm, err := time.ParseInLocation("2006-01-02 15:04:05", dt, time.UTC)
	if err != nil {
		return 0
	}

	return tm.Unix() * 1000
}

func DateStringToTimestampMilli1(dt string) int64 {
	tm, err := time.ParseInLocation("2006-01-02", dt, time.UTC)
	if err != nil {
		return 0
	}

	return tm.Unix() * 1000
}

// 流水小时统计使用【存储对应小时的毫秒值】
func DateInt64FromTimestamp(now int64) int64 {
	in := time.Unix(now/1000, 0).In(time.UTC)
	return time.Date(in.Year(), in.Month(), in.Day(), in.Hour(), 0, 0, 0, in.Location()).Unix() * 1000
}

func DateInt64FromTimestampMilli(now int64, tz_seconds int) int64 {
	localtime := now + int64(tz_seconds)*1000
	in := time.Unix(0, localtime*int64(time.Millisecond)).In(time.UTC)
	return time.Date(in.Year(), in.Month(), in.Day(), in.Hour(), 0, 0, 0, in.Location()).Unix() * 1000
}

// 获取当前的时间 - Unix时间戳
func GetCurrentUnix() int64 {
	return time.Now().Unix()
}

// 获取当前的小时的时间戳
func GetCurrentHourMilliUnix() int64 {
	tm := time.Unix(GetCurrentUnix(), 0).In(time.UTC)

	hourStart := time.Date(tm.Year(), tm.Month(), tm.Day(), tm.Hour(), 0, 0, 0, time.UTC)
	return hourStart.Unix()

}

// 获取当前的天的时间戳
func GetCurrentDayMilliUnix(reduceDay int) int64 {
	tm := time.Unix(GetCurrentUnix(), 0).In(time.UTC)

	hourStart := time.Date(tm.Year(), tm.Month(), tm.Day()-reduceDay, 0, 0, 0, 0, time.UTC)
	return hourStart.Unix()

}

// 获取当前的时间 - 毫秒级时间戳
func GetCurrentMilliUnix() int64 {
	return time.Now().UnixNano() / 1000000
}

// 获取当前的时间 - 纳秒级时间戳
func GetCurrentNanoUnix() int64 {
	return time.Now().UnixNano()
}

// 获取传入的时间所在月份的第一天，即某月第一天的0点。如传入time.Now(), 返回当前月份的第一天0点时间。
func GetFirstDateOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)
	return GetZeroTime(d)
}

// 获取传入的时间所在月份的最后一天，即某月最后一天的0点。如传入time.Now(), 返回当前月份的最后一天0点时间。
func GetLastDateOfMonth(d time.Time) time.Time {
	return GetFirstDateOfMonth(d).AddDate(0, 1, -1)
}

// 获取某一天的0点时间
func GetZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

// 获取某一天的1点时间
func GetOneTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 1, 0, 0, 0, d.Location())
}

// 获取当前时间的小时
func GetCurrentTimeHour(now int64, tzSeconds int) int {
	localtime := now/1000 + int64(tzSeconds)
	return time.Unix(localtime, 0).In(time.UTC).Hour()
}

// 获取两个日期之间相差的天数
func GetTwoDateDays(startTime, endTime string, loc *time.Location) int64 {
	startUnix, _ := time.ParseInLocation("2006-01-02", startTime, loc)
	endUnix, _ := time.ParseInLocation("2006-01-02", endTime, loc)
	// 求相差天数
	date := (endUnix.Unix() - startUnix.Unix()) / 86400
	return date
}

func GetTodayTimeLeft() int64 {
	timeTemplate := "2006-01-02 15:04:05"
	timeStr := time.Now().Format("2006-01-02")
	now_time_str := time.Now().Format(timeTemplate)
	today_end_time_str := timeStr + " 23:59:59"

	formatTime1, _ := time.Parse(timeTemplate, now_time_str)
	formatTime2, _ := time.Parse(timeTemplate, today_end_time_str)

	t1 := formatTime1.Unix()
	t2 := formatTime2.Unix()

	return t2 - t1
}

var HourMinuteSecondsLayout = "15:04:000"

func GetWeekEndTime(now int64, tz_seconds int) int64 {
	localtime := now/1000 + int64(tz_seconds)
	tm := time.Unix(localtime, 0).In(time.UTC)

	offset := int(time.Monday - tm.Weekday())
	if offset > 0 {
		offset = -6
	}
	//sunday
	in := tm.AddDate(0, 0, offset+6)
	out := time.Date(in.Year(), in.Month(), in.Day(), 23, 59, 59, 999, time.UTC)

	return out.UnixMilli()
}
