package go_redis

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/shopspring/decimal"
	"github.com/spf13/cast"
	mtime "go_primer/time"
	"strconv"
	"time"
)

func Redis() (*redis.Client, context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*10)
	return RedisDB, ctx, cancel

}

func RedisIncrByAndExpire(key string, amount int64, duration time.Duration) (int64, error) {
	client, ctx, cancel := Redis()
	defer cancel()

	lua := ""
	expTmsStr := strconv.FormatInt(duration.Milliseconds()/1000, 10)
	amountStr := strconv.FormatInt(amount, 10)
	if res, err := client.Eval(ctx, lua, []string{key}, []string{amountStr, expTmsStr}).Int64(); err != nil {
		return 0, err
	} else {
		if res < 0 {
			return 0, errors.New("RedisIncrByAndExpire failed")
		} else {
			return res, nil
		}
	}
}

func RedisGet(key string) (string, error) {
	client, ctx, cancel := Redis()
	defer cancel()

	result, err := client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}

func RedisGetInt64(key string) (int64, error) {
	client, ctx, cancel := Redis()
	defer cancel()
	result, err := client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return 0, nil
		}
		return 0, err
	}

	return cast.ToInt64(result), nil
}

func RedisExpire(key string, expir time.Duration) (bool, error) {
	client, ctx, cancel := Redis()
	defer cancel()
	result, err := client.Expire(ctx, key, expir).Result()
	if err != nil {
		return false, err
	}
	return result, nil
}

func RedisSetNx(key string, expir time.Duration) (bool, error) {
	client, ctx, cancel := Redis()
	defer cancel()
	ok, err := client.SetNX(ctx, key, 1, expir).Result()
	if err != nil {
		return false, err
	}
	client.Expire(ctx, key, expir)
	return ok, nil
}

func RedisSet(key, value string, expir time.Duration) (string, error) {
	client, ctx, cancel := Redis()
	defer cancel()
	res, err := client.Set(ctx, key, value, expir).Result()
	if err != nil {
		return "", err
	}
	client.Expire(ctx, key, expir)
	return res, nil
}

func RedisLPushAndPop(key, value string, expire time.Duration) (bool, error) {
	client, ctx, cancel := Redis()
	defer cancel()

	pipeLine := client.Pipeline()
	pipeLine.RPop(ctx, key)
	pipeLine.LPush(ctx, key, value)
	pipeLine.Expire(ctx, key, expire)
	_, err := pipeLine.Exec(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}

// FIFO, redis，list，当超出 limited 后，会被pop
func RedisLPushLimited(key, value string, limited int32, expire time.Duration) (int64, error) {
	client, ctx, cancel := Redis()
	defer cancel()

	lua := "local count = redis.call('lpush', KEYS[1], ARGV[1]); " +
		"if count > ARGV[2] then" +
		"    redis.call('rpop', KEYS[1])" +
		"end " +
		"redis.call('expire', KEYS[1], ARGV[3]); " +
		"return count; "

	expTms := expire.Seconds()
	retV, err := EvalInt(client, ctx, lua, []string{key}, []string{value, cast.ToString(limited), cast.ToString(expTms)})
	if err != nil {
		return 0, err
	}
	return retV, nil
}

// FIFO, redis，list，当超出 limited 后，会被ltrim
func RedisLPushLimitedV2(key, value string, limited int32, expire time.Duration) (int64, error) {
	client, ctx, cancel := Redis()
	defer cancel()

	lua := "local count = redis.call('lpush', KEYS[1], ARGV[1]); " +
		"if count > tonumber(ARGV[2]) then" +
		"    redis.call('ltrim', KEYS[1], 0, ARGV[2])" +
		"end " +
		"redis.call('expire', KEYS[1], ARGV[3]); " +
		"return count; "

	expTms := expire.Seconds()
	retV, err := EvalInt(client, ctx, lua, []string{key}, []string{value, cast.ToString(limited - 1), cast.ToString(expTms)})
	if err != nil {
		return 0, err
	}
	return retV, nil
}

func RedisPushAndExpire(key string, value []interface{}, expire time.Duration) (bool, error) {
	client, ctx, cancel := Redis()
	defer cancel()

	pipeLine := client.Pipeline()
	pipeLine.LPush(ctx, key, value)
	pipeLine.Expire(ctx, key, expire)
	_, err := pipeLine.Exec(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}

func RedisLRange(key string, start, end int64) ([]string, error) {
	client, ctx, cancel := Redis()
	defer cancel()

	result, err := client.LRange(ctx, key, start, end).Result()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func RedisLRangeInt64(key string, start, end int64) ([]int64, error) {
	client, ctx, cancel := Redis()
	defer cancel()

	result, err := client.LRange(ctx, key, start, end).Result()
	if err != nil {
		return nil, err
	}
	arr := make([]int64, len(result))
	for i, v := range result {
		arr[i] = cast.ToInt64(v)
	}
	return arr, nil
}

func RedisLLen(key string) (int64, error) {
	client, ctx, cancel := Redis()
	defer cancel()

	count, err := client.LLen(ctx, key).Result()
	if err != nil && err != redis.Nil {
		return 0, err
	}
	if err == redis.Nil {
		return 0, nil
	}
	return count, nil
}

func GetRankList(key string, start, stop int64) ([]string, error) {
	client, ctx, cancel := Redis()
	defer cancel()

	result, err := client.LRange(ctx, key, start, stop).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}
	return result, err
}

func RedisZRevRangeWithScore(key string, start, stop int64) ([]redis.Z, error) {
	client, ctx, cancel := Redis()
	defer cancel()

	result, err := client.ZRevRangeWithScores(ctx, key, start, stop).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}
	return result, err
}

type redisZ struct {
	Score  int64
	Member int64
}

func RedisZRevRangeWithScoreMemberInt64(key string, start, stop int64) ([]int64, []redisZ, error) {
	client, ctx, cancel := Redis()
	defer cancel()

	result, err := client.ZRevRangeWithScores(ctx, key, start, stop).Result()
	if err != nil && err != redis.Nil {
		return nil, nil, err
	}
	arr := make([]redisZ, len(result))
	memberArr := make([]int64, len(result))
	for i, rdsz := range result {
		arr[i] = redisZ{
			Score:  cast.ToInt64(rdsz.Score),
			Member: cast.ToInt64(rdsz.Member),
		}
		memberArr[i] = cast.ToInt64(rdsz.Member)
	}
	return memberArr, arr, err
}

func RedisZRevRangeByScoresWithScores(key string, offset, count int64) ([]redis.Z, error) {
	client, ctx, cancel := Redis()
	defer cancel()

	result, err := client.ZRevRangeByScoreWithScores(ctx, key, &redis.ZRangeBy{
		Min:    "-inf",
		Max:    "+inf",
		Offset: offset,
		Count:  count,
	}).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}
	return result, err
}

func GuardRedisZRevRange(key string, start, stop int64) ([]string, error) {
	client, ctx, cancel := Redis()
	defer cancel()

	result, err := client.ZRevRange(ctx, key, start, stop).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}
	return result, err
}

func RedisHGetAll(key string) (map[string]string, error) {
	client, ctx, cancel := Redis()
	defer cancel()

	result, err := client.HGetAll(ctx, key).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}
	return result, err
}

func RedisHGet(key, member string) (string, error) {
	client, ctx, cancel := Redis()
	defer cancel()

	result, err := client.HGet(ctx, key, member).Result()
	if err != nil && err != redis.Nil {
		return "", err
	}
	return result, err
}

func RedisHMSet(key string, values []string, expire time.Duration) error {
	client, ctx, cancel := Redis()
	defer cancel()

	_, err := client.HMSet(ctx, key, values).Result()
	if err != nil {
		return err
	}
	client.Expire(ctx, key, expire)
	return nil
}

func RedisHDel(key string, values []string, expTime time.Duration) error {
	client, ctx, cancel := Redis()
	defer cancel()

	_, err := client.HDel(ctx, key, values...).Result()
	if err != nil {
		return err
	}
	client.Expire(ctx, key, expTime)
	return nil
}

func RedisDel(keys []string) error {
	client, ctx, cancel := Redis()
	defer cancel()

	_, err := client.Del(ctx, keys...).Result()
	if err != nil {
		return err
	}
	return nil
}

func RedisRename(key, newKey string) error {
	client, ctx, cancel := Redis()
	defer cancel()

	_, err := client.Rename(ctx, key, newKey).Result()
	if err != nil {
		return err
	}
	return nil
}

func RedisZscore(key string, member string) (float64, error) {
	client, ctx, cancel := Redis()
	defer cancel()

	result, err := client.ZScore(ctx, key, member).Result()
	if err != nil && err != redis.Nil {
		return 0, err
	}
	return result, err
}

func RedisInZset(key string, member string) (bool, error) {
	client, ctx, cancel := Redis()
	defer cancel()

	_, err := client.ZScore(ctx, key, member).Result()
	if err != nil {
		if err == redis.Nil {
			return false, nil
		}
		return false, err
	}
	return true, err
}

func RedisZRevRank(key string, member string) (int64, error) {
	client, ctx, cancel := Redis()
	defer cancel()

	rank, err := client.ZRevRank(ctx, key, member).Result()
	if err != nil {
		if err == redis.Nil {
			return -1, nil
		}
		return -1, err
	}
	return rank, err
}

func RedisZIncryBy(key string, member string, increment float64) (float64, error) {
	client, ctx, cancel := Redis()
	defer cancel()

	result, err := client.ZIncrBy(ctx, key, increment, member).Result()
	if err != nil {
		return 0, err
	}
	return result, err
}

func RedisZIncryByAndExpire(key, member string, increment int64, duration time.Duration) (int64, error) {
	client, ctx, cancel := Redis()
	defer cancel()
	lua := `
		local num = redis.call('zincrby',KEYS[1], ARGV[1], ARGV[2]); 
		redis.call('expire',KEYS[1],ARGV[3]); 
		return num`

	expTmsStr := cast.ToString(duration.Milliseconds() / 1000)
	res, err := client.Eval(ctx, lua, []string{key}, []string{strconv.FormatInt(increment, 10), member, expTmsStr}).Int64()
	if err != nil {
		return 0, err
	}
	return res, nil
}

func RedisHIncryByAndExpire(key, field string, increment int64, expire time.Duration) (int64, error) {
	client, ctx, cancel := Redis()
	defer cancel()
	lua := `
		local num = redis.call('hincrby',KEYS[1], ARGV[1], ARGV[2]); 
		redis.call('expire',KEYS[1],ARGV[3]); 
		return num`

	expTmsStr := strconv.FormatInt(expire.Milliseconds()/1000, 10)
	res, err := client.Eval(ctx, lua, []string{key}, []string{field, strconv.FormatInt(increment, 10), expTmsStr}).Int64()
	if err != nil {
		return 0, err
	}
	return res, nil
}

func RedisHSetAndExpire(key, field string, value string, expire time.Duration) (int64, error) {
	client, ctx, cancel := Redis()
	defer cancel()
	lua := `
		local num = redis.call('hset',KEYS[1], ARGV[1], ARGV[2]); 
		redis.call('expire',KEYS[1],ARGV[3]); 
		return num`

	expTmsStr := strconv.FormatInt(expire.Milliseconds()/1000, 10)
	res, err := client.Eval(ctx, lua, []string{key}, []string{field, value, expTmsStr}).Int64()
	if err != nil {
		return 0, err
	}
	return res, nil
}

func RedisZAdd(key string, members []*redis.Z, duration time.Duration) (int64, error) {
	client, ctx, cancel := Redis()

	defer cancel()

	result, err := client.ZAdd(ctx, key, members...).Result()
	if err != nil && err != redis.Nil {
		return 0, err
	}
	client.Expire(ctx, key, duration)
	return result, nil
}

func RedisZRem(key string, members []string) (int64, error) {
	client, ctx, cancel := Redis()
	defer cancel()

	result, err := client.ZRem(ctx, key, members).Result()
	if err != nil && err != redis.Nil {
		return 0, err
	}
	return result, nil
}

func GetMyRank(key string, uid int64) (int64, int64, error) {
	client, ctx, cancel := Redis()
	defer cancel()

	myRank, err := client.ZRevRank(ctx, key, cast.ToString(uid)).Result()
	if err != nil {
		if err == redis.Nil { // 没有排名, 取最后一名的
			z, err := client.ZRangeByScoreWithScores(ctx, key, &redis.ZRangeBy{
				Min:   "-inf",
				Max:   "+inf",
				Count: 1,
			}).Result()
			if err != nil {
				return 0, 0, err
			}

			if len(z) > 0 { // 存在最后一名
				return -1, int64(z[0].Score), nil
			}
			return -1, -1, nil
		} else {
			return 0, 0, err
		}
	}

	var lastRankScore int64
	if myRank == 0 { // 排名第一
		lastRankScore = int64(client.ZScore(ctx, key, cast.ToString(uid)).Val())
	} else { // 排名非第一
		myScore := int64(client.ZScore(ctx, key, cast.ToString(uid)).Val())
		z, err := client.ZRangeByScoreWithScores(ctx, key, &redis.ZRangeBy{
			Min:    cast.ToString(myScore + 1),
			Max:    "+inf",
			Offset: 0,
			Count:  1,
		}).Result()
		if err != nil && err != redis.Nil {
			return 0, 0, err
		}

		if len(z) == 1 {
			lastRankScore = int64(z[0].Score) - myScore
		}
	}

	return myRank + 1, lastRankScore, nil
}

// RedisZIncryByFIFO 相同时，先到达的排在前
func RedisZIncryByFIFO(key, member string, increment, tmms int64, duration time.Duration) (int64, error) {
	client, ctx, cancel := Redis()
	defer cancel()
	lua := `
		local num = redis.call('zincrby', KEYS[1], ARGV[1], ARGV[2]);
		local new_score = string.format("%d",num) .. "." .. ARGV[3]
		redis.call('zadd', KEYS[1], new_score, ARGV[2]);
		redis.call('expire', KEYS[1], ARGV[4]);
		return new_score 
	`
	expTmsStr := cast.ToString(duration.Milliseconds() / 1000)
	fraction := cast.ToString(mtime.MaxUnixTimeStamp - tmms)

	res, err := client.Eval(ctx, lua, []string{key}, []string{cast.ToString(increment), member, fraction, expTmsStr}).Result()
	if err != nil {
		return 0, err
	}
	f, err := decimal.NewFromString(cast.ToString(res))
	if err != nil {
		return 0, err
	}
	fmt.Println(f.IntPart())
	return f.IntPart(), nil
}

func RedisZIncryByFIFO2(key, member string, increment, tmms int64, duration time.Duration) (int64, error) {
	client, ctx, cancel := Redis()
	defer cancel()
	lua := `
		local num = redis.call('zincrby', KEYS[1], ARGV[1], ARGV[2]);
		local new_score = string.format("%d",num) .. "." .. ARGV[3]
		redis.call('zadd', KEYS[1], new_score, ARGV[2]);
		redis.call('expire', KEYS[1], ARGV[4]);
		return new_score 
	`
	expTmsStr := cast.ToString(duration.Milliseconds() / 1000)
	fraction := cast.ToString(mtime.MaxUnixTimeStamp - tmms)

	res, err := client.Eval(ctx, lua, []string{key}, []string{cast.ToString(increment), member, fraction, expTmsStr}).Result()
	if err != nil {
		return 0, err
	}
	f, err := decimal.NewFromString(cast.ToString(res))
	if err != nil {
		return 0, err
	}
	return f.IntPart(), nil
}
