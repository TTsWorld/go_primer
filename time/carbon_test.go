package mtime

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang-module/carbon"
)

func TestCarbon(t *testing.T) {
	day := 3
	for i := day; i > 0; i-- {
		println(i)
		fmt.Println(carbon.Now().AddDays(-i).StartOfDay().ToDateTimeString())
		fmt.Println(carbon.Now().AddDays(-i).StartOfDay().TimestampMilli())
		fmt.Println(carbon.Now().AddDays(-i + 1).StartOfDay().ToDateTimeString())
		fmt.Println(carbon.Now().AddDays(-i + 1).StartOfDay().TimestampMilli())
		if i == 1 {
			println("xxxxxxxxxxxxx")

		}
	}
}

func TestCarbon2(t *testing.T) {
	d1 := carbon.Now().AddHours(-13)
	d2 := carbon.Now()
	fmt.Println(d1)
	fmt.Println(d2)
	fmt.Println(d1.DiffAbsInDays(d2))
	fmt.Println(d1.StartOfDay())
	fmt.Println(d2.StartOfDay())
	fmt.Println(d1.StartOfDay().DiffAbsInDays(d2.StartOfDay()))
}

func TestCarbonWeek(t *testing.T) {
	d1 := carbon.Now()
	d1.SetWeekStartsAt("Monday")
	fmt.Println(d1)
	fmt.Println(d1.Week())
	fmt.Println(d1.WeekOfMonth())
	fmt.Println(d1.WeekOfYear())
	fmt.Println(d1.DayOfWeek())
	fmt.Println(d1.EndOfWeek())
	fmt.Println(d1.StartOfWeek())
	fmt.Println(d1.IsWeekday())
	fmt.Println(d1.ToShortWeekString())
	fmt.Println(d1.ToWeekString())
}

func TestCarbon3(t *testing.T) {
	fmt.Println(carbon.Now().ToDateTimeString())
	fmt.Println(carbon.Now().AddHours(2).ToDateTimeString())
	fmt.Println(carbon.Now(carbon.UTC).ToDateTimeString())
	fmt.Println(carbon.Now(carbon.UTC).AddHours(2).ToDateTimeString())
	fmt.Println(carbon.Now(carbon.UTC).AddHours(25).StartOfDay())
}

func TestCarbon4(t *testing.T) {
	fmt.Println(carbon.Now().ToDateTimeString())
	fmt.Println(carbon.Now().AddHours(20).ToDateTimeString())
	fmt.Println(carbon.Now().DiffAbsInDays(carbon.Now().AddHours(20)))
}

func TestCarbon5(t *testing.T) {
	fmt.Println(carbon.Now().StartOfWeek())
	fmt.Println(carbon.Now().EndOfWeek())
	fmt.Println(carbon.Now().StartOfWeek().Second())
	fmt.Println(carbon.Now().EndOfWeek().Year())
	fmt.Println(carbon.Now().EndOfWeek().Month())
	fmt.Println(carbon.Now().EndOfWeek().Day())
	fmt.Println(time.Now().Weekday())
}
func TestCarbon6(t *testing.T) {
	c := carbon.CreateFromTimestampMilli(1700031300000)
	//fmt.Println(c.ToDateTimeString(carbon.UTC))

	c = carbon.Parse("2023-11-03 02:40:05", carbon.UTC)
	fmt.Println(c.ToDateTimeString(carbon.UTC))
	fmt.Println(c.ToDateTimeString(carbon.Shanghai))
	fmt.Println(c.TimestampMilli())
}
