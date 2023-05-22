package go_redis

import (
	"fmt"
	"github.com/spf13/cast"
	"testing"
)

func TestZsetExists(t *testing.T) {
	val, err := RedisDB.ZScore(Gctx, "a", "b").Result()
	fmt.Printf("%+v\n", val)
	fmt.Printf("%+v\n", err)
}

func TestZRevRank(t *testing.T) {
	val, err := RedisDB.ZRevRank(Gctx, "a", "a").Result()
	fmt.Printf("%+v\n", val)
	fmt.Printf("%+v\n", err)
}

func TestSomething(t *testing.T) {
	arr := []int{10, 9, 8, 7, 6, 5, 3, 3, 2, 1}
	v := 4
	lenArr := len(arr)
	for i := lenArr - 1; i >= 0; i-- {
		tmp := arr[i]
		if tmp > v {
			fmt.Printf("idx:%d, value:%d\n", i, tmp-v)
			break
		}
	}
}

// 测试 redis float 精度问题
func TestZFloat64(t *testing.T) {
	val, _ := RedisDB.ZRevRangeWithScores(Gctx, "z", 0, -1).Result()
	for _, rdsZ := range val {
		fmt.Printf("%+v\n", rdsZ.Member)
		fmt.Printf("%.10f\n", rdsZ.Score)
		fmt.Println(cast.ToInt64(rdsZ.Score))
	}
}
