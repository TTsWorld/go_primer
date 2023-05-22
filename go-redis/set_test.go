package go_redis

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-redis/redis"
)

var set1 = "set1"

func TestSet(t *testing.T) {
	RedisDB.SAdd(Gctx, set1, 1)
	RedisDB.SAdd(Gctx, set1, 2)
	RedisDB.SAdd(Gctx, set1, 3)

	val, err := RedisDB.SMembers(Gctx, set1).Result()
	fmt.Printf("%#v :::: %#v\n", val, err == nil)
	//[]string{"1", "2", "3"} :::: true

	val1, err := RedisDB.SIsMember(Gctx, set1, 2).Result()
	fmt.Printf("%#v :::: %#v ::: %#v\n", val1, err == nil, err == redis.Nil)
	//true :::: true ::: false

	val2, err := RedisDB.SCard(Gctx, set1).Result()
	fmt.Printf("%#v :::: %#v ::: %#v\n", val2, err == nil, err == redis.Nil)
	// 3 :::: true ::: false

	val3, err := RedisDB.SMembersMap(Gctx, set1).Result()
	fmt.Printf("%#v :::: %#v ::: %#v\n", val3, err == nil, err == redis.Nil)
	//map_t[string]struct {}{"1":struct {}{}, "2":struct {}{}, "3":struct {}{}} :::: true ::: false

	val4, err := RedisDB.SRem(Gctx, set1, 3).Result()
	fmt.Printf("%#v :::: %#v ::: %#v\n", val4, err == nil, err == redis.Nil)
	//1 :::: true ::: false

	val5, err := RedisDB.SMembers(Gctx, set1).Result()
	fmt.Printf("%#v :::: %#v\n", val5, err == nil)
	//[]string{"1", "2"} :::: true

	// key not is a set1
	val6, err := RedisDB.SRem(Gctx, str1, 3).Result()
	fmt.Printf("%#v :::: %#v :::%#v ::: %#v\n", val6, err, err == nil, err == redis.Nil)
	//1 :::: true ::: false

	// key  exists
	val7, err := RedisDB.SCard(Gctx, set1).Result()
	fmt.Printf("%#v :::: %#v :::%#v ::: %#v\n", val7, err, err == nil, err == redis.Nil)
	//1 :::: true ::: false

	// key not exists
	val8, err := RedisDB.SCard(Gctx, "keynotexists").Result()
	fmt.Printf("%#v :::: %#v :::%#v ::: %#v\n", val8, err, err == nil, err == redis.Nil)
	//1 :::: true ::: false

	ok, err1 := RedisDB.SetNX(Gctx, "setnxtest2", "1", time.Minute).Result()
	fmt.Printf("%#v :::: %#v \n", ok, err1)
	ok, err2 := RedisDB.SetNX(Gctx, "setnxtest2", "1", time.Minute).Result()
	fmt.Printf("%#v :::: %#v \n", ok, err2)

}
