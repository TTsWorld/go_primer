package go_redis

import (
	"fmt"
	"testing"
)

func TestZsetExists(t *testing.T) {
	val, err := redisDB.ZScore(ctx, "a", "b").Result()
	fmt.Printf("%+v\n", val)
	fmt.Printf("%+v\n", err)
}

func TestZRevRank(t *testing.T) {
	val, err := redisDB.ZRevRank(ctx, "a", "a").Result()
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
