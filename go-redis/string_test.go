package go_redis

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/go-redis/redis/v8"
)

var redisDB *redis.Client
var ctx context.Context
var str1, str2, str3, str4 = "str1", "str2", "str3", "str4"

func TestPing(t *testing.T) {
	println(redisDB.Ping(ctx).Result())
}

func TestSet01(t *testing.T) {
	ctx := context.Background()
	val, err := redisDB.Set(ctx, str1, str1, 0).Result()
	if err != nil {
		panic(err)
	}
	println(val, reflect.TypeOf(val).Name())

	println("===============")
	//get no exists key
	err = redisDB.Del(ctx, str2).Err()
	fmt.Printf("%v\n", err)
	val, err = redisDB.Get(ctx, str2).Result()
	fmt.Printf("%v\n", val)
	fmt.Printf("%v\n", err == nil)
	fmt.Printf("%v\n", err == redis.Nil)
	if err == redis.Nil {
		println("==xxxxxxx=========")
	}

	println("===============")

	//getset && setnx
	//redisDB.GetSet(str2, str2)
}
func TestMget(t *testing.T) {
	redisDB.Set(ctx, str3, str3, 0)
	fmt.Printf("%#v\n", redisDB.MGet(ctx, str1, str2, str3, str4).String())
	fmt.Printf("%#v\n", redisDB.MGet(ctx, str1, str2, str3, str4).Name())
	fmt.Printf("%#v\n", redisDB.MGet(ctx, str1, str2, str3, str4).Val()[1])
	fmt.Printf("%#v\n", redisDB.MGet(ctx, str1, str2, str3, str4).Args())
	println("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	val, err := redisDB.MGet(ctx, str1, str2, str3, str4).Result()
	fmt.Printf("%#v\n %#v\n", val, err == nil)
	fmt.Printf("%#v\n %#v\n", val, err == redis.Nil)
}
