package go_redis

import (
	"fmt"
	"testing"

	"github.com/go-redis/redis/v8"
)

func TestBitSet(t *testing.T) {
	key := "bit_set"

	//key offset上没有 key
	val, err := RedisDB.SetBit(Gctx, key, 0, 1).Result()
	fmt.Printf("%#v :::: %#v\n", val, err == nil)
	//0 :::: true

	//key offset上有 key
	val, err = RedisDB.SetBit(Gctx, key, 0, 1).Result()
	fmt.Printf("%#v :::: %#v\n", val, err == nil)
	//1 :::: true

	//key offset上有 key
	val, err = RedisDB.SetBit(Gctx, key, 2, 1).Result()
	fmt.Printf("%#v :::: %#v\n", val, err == nil)
	//0 :::: true

	val, err = RedisDB.GetBit(Gctx, key, 2).Result()
	fmt.Printf("%#v :::: %#v\n", val, err)

	//key offset上没有 key, err=nil, 返回 0
	val, err = RedisDB.GetBit(Gctx, key, 100).Result()
	fmt.Printf("%#v :::: %#v\n", val, err)
	//0 :::: true
}

func TestBitSet2(t *testing.T) {
	key := "set"

	//key offset上没有 key
	val, err := RedisDB.SetBit(Gctx, key, 0, 1).Result()
	val, err = RedisDB.SetBit(Gctx, key, 1, 1).Result()
	val, err = RedisDB.SetBit(Gctx, key, 2, 1).Result()
	val, err = RedisDB.SetBit(Gctx, key, 3, 1).Result()
	val, err = RedisDB.SetBit(Gctx, key, 4, 1).Result()
	val, err = RedisDB.SetBit(Gctx, key, 5, 1).Result()
	val, err = RedisDB.SetBit(Gctx, key, 10, 1).Result()

	val, err = RedisDB.GetBit(Gctx, key, 0).Result()
	fmt.Printf("%#v :::: %#v\n", val, err)
	val, err = RedisDB.GetBit(Gctx, key, 1).Result()
	fmt.Printf("%#v :::: %#v\n", val, err)
	val, err = RedisDB.GetBit(Gctx, key, 2).Result()
	fmt.Printf("%#v :::: %#v\n", val, err)
	val, err = RedisDB.GetBit(Gctx, key, 3).Result()
	fmt.Printf("%#v :::: %#v\n", val, err)
	val, err = RedisDB.GetBit(Gctx, key, 4).Result()
	fmt.Printf("%#v :::: %#v\n", val, err)
	val, err = RedisDB.GetBit(Gctx, key, 5).Result()
	fmt.Printf("%#v :::: %#v\n", val, err)
	val, err = RedisDB.GetBit(Gctx, key, 6).Result()
	fmt.Printf("%#v :::: %#v\n", val, err)
	fmt.Println("===========================")

	val, err = RedisDB.BitPos(Gctx, key, 1, 0, 10).Result()
	fmt.Printf("%#v :::: %#v\n", val, err)

	val, err = RedisDB.BitPos(Gctx, key, 0, 0, 10).Result()
	fmt.Printf("%#v :::: %#v\n", val, err)

	val, err = RedisDB.BitPos(Gctx, key, 1, 6, 10).Result()
	fmt.Printf("%#v :::: %#v\n", val, err)

	val, err = RedisDB.BitCount(Gctx, key, &redis.BitCount{Start: 0, End: -1}).Result()
	fmt.Printf("%#v :::: %#v\n", val, err)

	val1, err := RedisDB.Get(Gctx, key).Result()
	fmt.Printf("%s, %b, %#v :::: %#v\n", val1, []byte(val1), val1, err)

}

func TestBitSet3(t *testing.T) {
	key := "ccc"

	//key offset上没有 key
	val, err := RedisDB.SetBit(Gctx, key, 0, 1).Result()
	val = val
	val1, err := RedisDB.Get(Gctx, key).Result()
	fmt.Printf("%s, %b, %#v :::: %#v\n", val1, []byte(val1), val1, err)

	val, err = RedisDB.SetBit(Gctx, key, 1, 1).Result()
	val1, err = RedisDB.Get(Gctx, key).Result()
	fmt.Printf("%s, %b, %#v :::: %#v\n", val1, []byte(val1), val1, err)

	val, err = RedisDB.SetBit(Gctx, key, 2, 1).Result()
	val1, err = RedisDB.Get(Gctx, key).Result()
	fmt.Printf("%s, %b, %#v :::: %#v\n", val1, []byte(val1), val1, err)

	val, err = RedisDB.SetBit(Gctx, key, 3, 1).Result()
	val1, err = RedisDB.Get(Gctx, key).Result()
	fmt.Printf("%s, %b, %#v :::: %#v\n", val1, []byte(val1), val1, err)

	val, err = RedisDB.SetBit(Gctx, key, 4, 1).Result()
	val1, err = RedisDB.Get(Gctx, key).Result()
	fmt.Printf("%s, %b, %#v :::: %#v\n", val1, []byte(val1), val1, err)

	val, err = RedisDB.SetBit(Gctx, key, 5, 1).Result()
	val1, err = RedisDB.Get(Gctx, key).Result()
	fmt.Printf("%s, %b, %#v :::: %#v\n", val1, []byte(val1), val1, err)

	val, err = RedisDB.SetBit(Gctx, key, 10, 1).Result()
	val2, err := RedisDB.Get(Gctx, key).Bytes()
	fmt.Printf("%v, %b\n", val2, val2)

	val, err = RedisDB.SetBit(Gctx, key, 11, 1).Result()
	val, err = RedisDB.SetBit(Gctx, key, 9, 1).Result()
	val3, err := RedisDB.Get(Gctx, key).Int()
	fmt.Printf("%v, %b\n, err:%v", val3, val3, err)
}

func ByteToBinaryBytes() {

}

func TestBitSet4(t *testing.T) {
	key := "dddd"

	//key offset上没有 key
	val, err := RedisDB.SetBit(Gctx, key, 0, 1).Result()
	val = val
	val1, err := RedisDB.Get(Gctx, key).Result()
	fmt.Printf("%s, %b, %#v :::: %#v\n", val1, []byte(val1), val1, err)

	val, err = RedisDB.SetBit(Gctx, key, 1, 1).Result()
	val1, err = RedisDB.Get(Gctx, key).Result()
	fmt.Printf("%s, %b, %#v :::: %#v\n", val1, []byte(val1), val1, err)

	//val, err = RedisDB.SetBit(Gctx, key, 2, 0).Result()
	//val1, err = RedisDB.Get(Gctx, key).Result()
	//fmt.Printf("%s, %b, %#v :::: %#v\n", val1, []byte(val1), val1, err)
	//
	//val, err = RedisDB.SetBit(Gctx, key, 3, 0).Result()
	//val1, err = RedisDB.Get(Gctx, key).Result()
	//fmt.Printf("%s, %b, %#v :::: %#v\n", val1, []byte(val1), val1, err)

	val, err = RedisDB.SetBit(Gctx, key, 4, 1).Result()
	val1, err = RedisDB.Get(Gctx, key).Result()
	fmt.Printf("%s, %b, %#v :::: %#v\n", val1, []byte(val1), val1, err)

	val, err = RedisDB.SetBit(Gctx, key, 5, 1).Result()
	val1, err = RedisDB.Get(Gctx, key).Result()
	fmt.Printf("%s, %b, %#v :::: %#v\n", val1, []byte(val1), val1, err)

	val2, err := RedisDB.Get(Gctx, key).Bytes()
	fmt.Printf("%v, %b\n", val2, val2)
}
