package go_redis

import (
	"github.com/spf13/cast"
	"golang.org/x/xerrors"
	"strconv"

	"context"
	"github.com/go-redis/redis/v8"
)

type REval interface {
	Eval(ctx context.Context, script string, keys []string, args ...interface{}) *redis.Cmd
	SScan(ctx context.Context, key string, cursor uint64, match string, count int64) *redis.ScanCmd
	Scan(ctx context.Context, cursor uint64, match string, count int64) *redis.ScanCmd
}

func EvalStringArray(r REval, ctx context.Context, lua string, keys []string, args []string) ([]string, error) {
	ret := []string{}
	res2, err := r.Eval(ctx, lua, keys, args).Result()
	if err != nil {
		return ret, err
	} else {
		switch res2.(type) {
		case []interface{}:
			x := res2.([]interface{})
			for i := 0; i < len(x); i++ {
				switch x[i].(type) {
				case string:
					ret = append(ret, x[i].(string))
				case int64:
					ret = append(ret, strconv.FormatInt(x[i].(int64), 10))
				case int:
					ret = append(ret, strconv.FormatInt(int64(x[i].(int)), 10))
				default:
					return ret, xerrors.Errorf("Invalid eval interface:%v", x[i])
				}
			}
			return ret, nil
		default:
			return ret, xerrors.Errorf("Invalid eval result %v", res2)
		}
	}
}

func EvalInt(r REval, ctx context.Context, lua string, keys []string, args []string) (int64, error) {
	res2, err := r.Eval(ctx, lua, keys, args).Result()
	if err != nil {
		return 0, err
	} else {
		switch res2.(type) {
		case string:
			ival, e := strconv.ParseInt(res2.(string), 10, 64)
			if e != nil {
				return 0, xerrors.Errorf("Invalid eval result:%v", res2)
			}
			return ival, nil
		case int64:
			return res2.(int64), nil
		case int:
			return int64(res2.(int)), nil
		default:
			return 0, xerrors.Errorf("Invalid eval result %v", res2)
		}
	}
}

func EvalIntArray(r REval, ctx context.Context, lua string, keys []string, args []string) ([]int64, error) {
	ret := []int64{}
	res2, err := r.Eval(ctx, lua, keys, args).Result()
	if err != nil {
		return ret, err
	} else {
		switch res2.(type) {
		case []interface{}:
			x := res2.([]interface{})
			for i := 0; i < len(x); i++ {
				switch x[i].(type) {
				case string:
					f64, e := strconv.ParseFloat(x[i].(string), 64)
					if e != nil {
						return ret, xerrors.Errorf("Invalid eval result:%v, idx:%d", x[i], i)
					}
					ret = append(ret, int64(f64))
				case int64:
					ret = append(ret, x[i].(int64))
				case int:
					ret = append(ret, int64(x[i].(int)))
				default:
					return ret, xerrors.Errorf("Invalid eval interface:%v", x[i])
				}
			}
			return ret, nil
		default:
			return ret, xerrors.Errorf("Invalid eval result %v", res2)
		}
	}
}

type RZSetCmd interface {
	ZRevRangeWithScores(ctx context.Context, key string, start int64, end int64) *redis.ZSliceCmd
	ZRevRangeByScoreWithScores(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.ZSliceCmd
	ZRangeWithScores(ctx context.Context, key string, start, stop int64) *redis.ZSliceCmd
	ZRangeByScoreWithScores(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.ZSliceCmd
}

type ZIntScore struct {
	Member int64
	Score  float64
}

func ZSliceToMap_I64(z []redis.Z) ([]ZIntScore, error) {
	ret := []ZIntScore{}
	for i := 0; i < len(z); i++ {
		switch z[i].Member.(type) {
		case string:
			if ival := cast.ToInt64(z[i].Member); ival != 0 {
				ret = append(ret, ZIntScore{Member: ival, Score: z[i].Score})
			}
		default:
			return ret, xerrors.Errorf("Invalid ZRevRangeWithScores_I64 member:%v", z[i])
		}
	}
	return ret, nil
}

func ZRevRangeWithScores_I64(r RZSetCmd, ctx context.Context, key string, start int64, end int64) ([]ZIntScore, error) {
	if results, err := r.ZRevRangeWithScores(ctx, key, start, end).Result(); err != nil {
		return []ZIntScore{}, err
	} else {
		return ZSliceToMap_I64(results)
	}
}

func ZRevRangeByScoreWithScores_I64(r RZSetCmd, ctx context.Context, key string, min string, max string, start int64, end int64) ([]ZIntScore, error) {
	op := redis.ZRangeBy{
		Min:    min,   // 最小分数
		Max:    max,   // 最大分数
		Offset: start, // 类似sql的limit, 表示开始偏移量
		Count:  end,   // 一次返回多少数据
	}
	if results, err := r.ZRevRangeByScoreWithScores(ctx, key, &op).Result(); err != nil {
		return []ZIntScore{}, err
	} else {
		return ZSliceToMap_I64(results)
	}
}
