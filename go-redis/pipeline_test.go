package go_redis

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
)

func TestPipeLine(t *testing.T) {
	redisDB.Set(ctx, "a", 1, time.Minute)
	redisDB.Set(ctx, "b", 1, time.Minute)
	redisDB.Set(ctx, "c", 1, time.Minute)
	redisDB.Set(ctx, "d", 1, time.Minute)
	redisDB.Set(ctx, "e", 1, time.Minute)

	cmds, err := redisDB.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		for _, v := range []string{"a", "b", "", "d", "e"} {
			pipe.Get(ctx, fmt.Sprintf("%s", v))
		}
		return nil
	})
	//if err != nil {
	//	panic(err)
	//}
	fmt.Printf("%#v\n ", err)

	for _, cmd := range cmds {
		result, err := cmd.(*redis.StringCmd).Result()
		fmt.Printf("%#v\t %#v", result, err)
		fmt.Printf("%#v\t", cmd.(*redis.StringCmd).Args())
		fmt.Printf("%#v\n", cmd.(*redis.StringCmd).Name())
	}
}

func TestPipeLineZset(t *testing.T) {
	redisDB.ZAdd(ctx, "a11", &redis.Z{Score: 1, Member: 1})
	redisDB.ZAdd(ctx, "a11", &redis.Z{Score: 2, Member: 2})
	redisDB.ZAdd(ctx, "a11", &redis.Z{Score: 3, Member: 3})
	redisDB.ZAdd(ctx, "a11", &redis.Z{Score: 4, Member: 4})
	keys := []string{"a11"}

	cmds, err := redisDB.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		for _, key := range keys {
			pipe.ZRevRangeWithScores(ctx, key, 0, -1)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("err:%v \n", err)
	}

	for _, cmd := range cmds {
		//TODO , 不行还得用 string slice cmd 转
		rdsZSlice, err := cmd.(*redis.ZSliceCmd).Result()
		for _, rdsZ := range rdsZSlice {
			member := rdsZ.Member.(string)
			score := rdsZ.Score

			fmt.Printf("member:%v, score:%v, err:%v\n", member, score, err)
		}
	}
}
