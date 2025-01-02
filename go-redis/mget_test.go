package go_redis

import (
	"fmt"
	"testing"
)

func TestMGet2(t *testing.T) {
	val, err := RedisDB.MGet(Gctx, "a", "b", "d").Result()
	fmt.Printf("%+v\n", err)
	fmt.Printf("%+v\n", val)
}
