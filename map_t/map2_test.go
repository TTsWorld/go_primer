package map_t

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/spf13/cast"
)

var dateFormt = "%s|%s|%s|%s"

func TestMap2(t *testing.T) {
	dateMap := map[string]int{} // key = 2023-10-23|sugo_event_kafka|event_id|vc|

	for i := 0; i < 1000; i++ {
		j := rand.Intn(3)
		topic := "topic" + cast.ToString(j)
		date := "date" + cast.ToString(j)
		eventId := "date" + cast.ToString(j)
		vc := "date" + cast.ToString(j)

		key := fmt.Sprintf(dateFormt, date, topic, eventId, vc)
		dateMap[key]++
		fmt.Printf("ConsumeClaim||dateMap:%+v \n", dateMap)
	}
}

// 测试 mapvar 声明候是否需要 make 申请内存
func Test0202(t *testing.T) {
	var m map[int]int
	m = make(map[int]int)
	m[1] = 0
	m[2] = 0
	fmt.Printf("%+v", m)
}

type Data struct {
	a int
}

// 测试 mapvar 声明候是否需要 make 申请内存
func Test0203(t *testing.T) {

	m := make(map[int]*Data)

	if m[1] != nil {
		fmt.Printf("%+v", *m[1])
		return
	}
	fmt.Printf("%+v", m[1])
}
