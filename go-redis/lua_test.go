package go_redis

import (
	"context"
	"fmt"
	"github.com/spf13/cast"
	"testing"
	"time"
)

func TestLua1(t *testing.T) {
	client, ctx, cancel := Redis()
	defer cancel()

	lua := `
		local n1 = redis.call('zscore', KEYS[1], ARGV[1]); 
		if (not n1) then
			n1=-1;
		end
		local n2 = redis.call('zrevrank', KEYS[1], ARGV[1]); 
		if (not n2)  then
			n2=-1;
		end
		return {n1, n2};
	`
	eval_result, err := client.Eval(ctx, lua, []string{"a"}, []string{"7"}).Float64Slice()
	fmt.Printf("%v \n %v \n n2 == '-1':%v", eval_result, err, eval_result[0] == -1)
}

func TestLua2(t *testing.T) {
	client, ctx, cancel := Redis()
	defer cancel()
	lua := `
		local n1 = redis.call('zscore', KEYS[1], ARGV[1]); 
		if (not n1) then
			return -1;
		end
		local n2 = redis.call('zrevrank', KEYS[1], ARGV[1]); 
		if (not n2)  then
			return -1;
		end
		return n1 .. "|" .. n2;
	`
	eval_result, err := client.Eval(ctx, lua, []string{"a"}, []string{"1"}).Float64Slice()
	fmt.Printf("%v \n %v", eval_result, err)
}

func TestLpushLua(t *testing.T) {
	//_, _, _ := Redis()

	lua := "local count = redis.call('lpush', KEYS[1], ARGV[1]); " +
		"if count > tonumber(ARGV[2]) then" +
		"    redis.call('rpop', KEYS[1])" +
		"end " +
		"redis.call('expire', KEYS[1], ARGV[3]); " +
		"return count; "

	key := "test_list"
	value := "2"
	limited := 5
	expTms := 1000
	retV, err := EvalInt(RedisDB, context.Background(), lua, []string{key}, []string{value, cast.ToString(limited), cast.ToString(expTms)})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(retV)
}

// 分数相同，排序在前
func TestZincry(t *testing.T) {
	lua := ` 
		local num = redis.call('zincrby', KEYS[1], ARGV[1], ARGV[2]);
		local new_score = string.format("%d",num) .. "." .. ARGV[3]
		redis.call('zadd', KEYS[1], new_score, ARGV[2]);
		redis.call('expire', KEYS[1], ARGV[4]);
		return new_score 
	`
	key := "z"
	argv1 := "1000000"
	argv2 := "2"
	v := cast.ToString(2147483647 - time.Now().Unix())
	argv3 := v
	expTms := "1000"
	retV, err := RedisDB.Eval(context.Background(), lua, []string{key}, []string{argv1, argv2, argv3, expTms}).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(retV)

	res, err := RedisDB.ZRevRangeWithScores(context.Background(), key, 0, -1).Result()
	if err != nil {
		fmt.Println(err)
	}
	for _, rdsZ := range res {
		fmt.Printf("member:%v, ", rdsZ.Member)
		fmt.Printf("score:%v \n", int64(rdsZ.Score))
	}
}
func TestZincryFIFO(t *testing.T) {
	after, err := RedisZIncryByFIFO("z", "2", 100, time.Now().Unix(), time.Minute*10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(after)

}
