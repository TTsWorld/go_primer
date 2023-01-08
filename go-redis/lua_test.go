package go_redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"testing"
)

func init() {
	redisDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	ctx = context.Background()
}

func TestLua1(t *testing.T) {
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
	eval_result, err := redisDB.Eval(ctx, lua, []string{"a"}, []string{"7"}).Float64Slice()
	fmt.Printf("%v \n %v \n n2 == '-1':%v", eval_result, err, eval_result[0] == -1)
}

func TestLua2(t *testing.T) {
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
	eval_result, err := redisDB.Eval(ctx, lua, []string{"a"}, []string{"1"}).Float64Slice()
	fmt.Printf("%v \n %v", eval_result, err)
}
