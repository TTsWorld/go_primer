package time

import (
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
