package go_redis

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/spf13/cast"

	"github.com/go-redis/redis/v8"
)

func TestPipeLine(t *testing.T) {
	RedisDB.Set(Gctx, "a", 1, time.Minute)
	RedisDB.Set(Gctx, "b", 1, time.Minute)
	RedisDB.Set(Gctx, "c", 1, time.Minute)
	RedisDB.Set(Gctx, "d", 1, time.Minute)
	RedisDB.Set(Gctx, "e", 1, time.Minute)

	cmds, err := RedisDB.Pipelined(Gctx, func(pipe redis.Pipeliner) error {
		for _, v := range []string{"a", "b", "", "d", "e"} {
			pipe.Get(Gctx, fmt.Sprintf("%s", v))
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
	RedisDB.ZAdd(Gctx, "a11", &redis.Z{Score: 1, Member: 1})
	RedisDB.ZAdd(Gctx, "a11", &redis.Z{Score: 2, Member: 2})
	RedisDB.ZAdd(Gctx, "a11", &redis.Z{Score: 3, Member: 3})
	RedisDB.ZAdd(Gctx, "a11", &redis.Z{Score: 4, Member: 4})
	keys := []string{"a11"}

	cmds, err := RedisDB.Pipelined(Gctx, func(pipe redis.Pipeliner) error {
		for _, key := range keys {
			pipe.ZRevRangeWithScores(Gctx, key, 0, -1)
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

func TestPipeLineGet(t *testing.T) {

	ctx := context.Background()
	uids := []int64{1, 2, 4}
	pipeLine := RedisDB.Pipeline()

	for _, uid := range uids {
		key := fmt.Sprintf("aa:%d", uid)
		pipeLine.Get(ctx, key)
	}

	res, err := pipeLine.Exec(ctx)
	if err != nil && err != redis.Nil {
		fmt.Printf("err:%v\n", err)
	}
	valList := make([]int64, len(uids))
	for idx, cmdResTmp := range res {
		cmdRes, ok := cmdResTmp.(*redis.StringCmd)
		if ok {
			if result, err := cmdRes.Result(); err == nil {
				valList[idx] = cast.ToInt64(result)
				continue
			}
		}
		valList[idx] = -1
	}
}
