package mtime

import (
	"fmt"
	"github.com/spf13/cast"
	"testing"
	"time"
)

func getPublishTime(publishTime int64) string {
	if time.Now().Unix()-publishTime > 86400 { //大于一天
		days := time.Since(time.Unix(publishTime, 0)).Hours() / 24
		return fmt.Sprintf("%d天前更新", cast.ToInt32(days))
	} else if time.Now().Unix()-publishTime < 3600 {
		return "1小时前更新"
	} else {
		hours := time.Since(time.Unix(publishTime, 0)).Hours()
		return fmt.Sprintf("%d小时前更新", cast.ToInt32(hours))
	}
}
func TestC(t *testing.T) {
	fmt.Println(getPublishTime(1740196569))

}
