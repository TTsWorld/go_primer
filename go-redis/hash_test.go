package go_redis

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/go-redis/redis/v8"
)

var hash1, hash2 = "hash1", "hash2"

func TestHash(t *testing.T) {
	//add
	RedisDB.HSet(Gctx, hash1, "a", 1)
	RedisDB.HSet(Gctx, hash1, "b", 2)
	RedisDB.HSet(Gctx, hash1, "c", 3)
	// if exists, update the field's value
	RedisDB.HSet(Gctx, hash1, "b", 4)

	//get all
	val, err := RedisDB.HGetAll(Gctx, hash1).Result()
	fmt.Printf("%#v :::: %#v :::%#v ::: %#v\n", val, err, err == nil, err == redis.Nil)
	//map_t[string]string{"a":"1", "b":"4", "c":"3", "e":"5", "z":"5"} :::: <nil> :::true ::: false
	val2 := RedisDB.HGetAll(Gctx, hash1).String()
	fmt.Printf("%#v :::: \n", val2)
	//"hgetall hash1: map_t[a:1 b:4 c:3 e:5 z:5]" ::::

	//get  exists
	val3, err := RedisDB.HGet(Gctx, hash1, "c").Result()
	fmt.Printf("%#v :::: %#v :::%#v ::: %#v\n", val3, err, err == nil, err == redis.Nil)
	//3 :::: true ::: false

	//get not exists
	val4, err := RedisDB.HGet(Gctx, hash1, "f").Result()
	fmt.Printf("%#v :::: %#v :::%#v ::: %#v\n", val4, err, err == nil, err == redis.Nil)
	//"" :::: "redis: nil" :::false ::: true

	//incr exists
	val5, err := RedisDB.HIncrBy(Gctx, hash1, "c", 5).Result()
	fmt.Printf("%#v :::: %#v :::%#v ::: %#v\n", val5, err, err == nil, err == redis.Nil)
	//8 :::: <nil> :::true ::: false

	//incr not exists, will add to hash and set the value to 5
	val6, err := RedisDB.HIncrBy(Gctx, hash1, "z", 5).Result()
	fmt.Printf("%#v :::: %#v :::%#v ::: %#v\n", val6, err, err == nil, err == redis.Nil)
	//20 :::: <nil> :::true ::: false

	//Hdel
	val7, err := RedisDB.HDel(Gctx, hash1, "z").Result()
	fmt.Printf("%#v :::: %#v :::%#v ::: %#v\n", val7, err, err == nil, err == redis.Nil)
	//1 :::: <nil> :::true ::: false

	// get HDEL field
	val8, err := RedisDB.HGet(Gctx, hash1, "z").Result()
	fmt.Printf("%#v :::: %#v :::%#v ::: %#v\n", val8, err, err == nil, err == redis.Nil)
	//"" :::: "redis: nil" :::false ::: true
}

func TestMHash(t *testing.T) {
	RedisDB.HSet(Gctx, hash2, "a", 1)
	RedisDB.HSet(Gctx, hash2, "b", 2)
	RedisDB.HSet(Gctx, hash2, "c", 3)
	RedisDB.HSet(Gctx, hash2, "d", 4)

	// HMSet, has exists and not exists
	val1, err := RedisDB.HMSet(Gctx, hash2, map[string]interface{}{"a": "a", "b": 4, "e": "e"}).Result()
	fmt.Printf("%#v :::: %#v :::%#v ::: %#v\n", val1, err, err == nil, err == redis.Nil)
	//"OK" :::: <nil> :::true ::: false

	//get all
	val, err := RedisDB.HGetAll(Gctx, hash2).Result()
	fmt.Printf("%#v :::: %#v :::%#v ::: %#v\n", val, err, err == nil, err == redis.Nil)
	//map_t[string]string{"a":"a", "b":"4", "c":"3", "d":"4", "e":"e"} :::: <nil> :::true ::: false

	//HMGet
	val2, err := RedisDB.HMGet(Gctx, hash2, "a", "b", "z", "f").Result()
	fmt.Printf("%#v :::: %#v :::%#v ::: %#v\n", val2, err, err == nil, err == redis.Nil)
	fmt.Printf("%#v :::: %#v :::: %#v:::: %#v:::%#v ::: %#v\n", val2, val2[0] == nil, val2[3] == nil, err, err == nil, err == redis.Nil)
	//[]interface {}{"a", "4", interface {}(nil), interface {}(nil)} :::: <nil> :::true ::: false
}

func TestHget(t *testing.T) {
	if err := RedisDB.HSet(Gctx, "keyisnotexist", "aa", 1).Err(); err != nil {
		fmt.Printf("HSet failed key:%s , value:%s v\n", "keyisnotexist", "aaa")
	}

	if err := RedisDB.HSet(Gctx, "keyexistvalnotexists", "aa", 1).Err(); err != nil {
		fmt.Printf("HSet failed key:%s , value:%s v\n", "keyexistvalnotexists", "aaa")
	}

	val := RedisDB.HGet(Gctx, "keyisnotexist0", "aaa").String()
	valInt, _ := strconv.Atoi(val)
	println(valInt)
	//fmt.Printf("%#v :::: %#v :::%#v ::: %#v\n", val, err, err == nil, err == redis.Nil)
	//"" :::: "redis: nil" :::false ::: true

	val2, err := RedisDB.HGet(Gctx, "keyexistvalnotexists", "bbb").Result()
	fmt.Printf("%#v :::: %#v :::: %#v:::: %#v:::%#v ::: %#v\n", val2, val2 == "0", val2 == "1", err, err == nil, err == redis.Nil)
	//"" :::: false :::: false:::: "redis: nil":::false ::: true

	RedisDB.HSet(Gctx, "keyexistvalexists", "aaa", 1)
	val3, err := RedisDB.HGet(Gctx, "keyexistvalexists", "aaa").Result()
	fmt.Printf("%#v :::: %#v :::: %#v:::: %#v:::%#v ::: %#v\n", val3, val3 == "0", val3 == "1", err, err == nil, err == redis.Nil)
	//"1" :::: false :::: true:::: <nil>:::true ::: false
}

func TestHMget(t *testing.T) {
	if err := RedisDB.HSet(Gctx, "keyexistvalnotexists1", "aa", 1, "bb", 2).Err(); err != nil {
		fmt.Printf("HSet failed key:%s , value:%s v\n", "keyexistvalnotexists1", "aaa")
	}
	result, err := RedisDB.HMGet(Gctx, "keyexistvalnotexists1", "aa", "cc").Result()
	fmt.Printf("result[0]: %#v :::: %#v :::: %#v:::: %#v:::%#v ::: %#v\n", result[0], result[0] == "0", result[0] == "1", err, err == nil, err == redis.Nil)
	fmt.Printf("result[1]: %#v :::: %#v :::: %#v:::: %#v:::%#v ::: %#v\n", result[1], result[1] == "0", result[1] == "1", err, result[1] == nil, result[1] == redis.Nil)
	//"" :::: false :::: false:::: "redis: nil":::false ::: true
}

type A struct {
	A int `redis:"a"`
	B int `redis:"b"`
	C int `redis:"c"`
	D int `redis:"d"`
}

func TestScan(t *testing.T) {
	a := A{}

	Gctx := context.Background()
	err := RedisDB.HGetAll(Gctx, "hh000").Scan(&a)
	fmt.Printf("%#v :::: %#v :::: %#v:::: %#v\n", a, err, err == nil, err == redis.Nil)
}

func TestHGet(t *testing.T) {
	RedisDB.HSet(Gctx, "aaa", "a", "1").Result()
	res, err := RedisDB.HGet(Gctx, "aaa", "a").Result()
	res2, err2 := RedisDB.HGet(Gctx, "aaa", "b").Result()
	fmt.Printf("HSet res:%s , err:%v", res, err)
	fmt.Printf("HSet res:%s , err:%v", res2, err2)
}

func TestHMset(t *testing.T) {
	setarr := []string{"1", "1", "2", "2", "3", "3", "4", "4"}
	setMap := map[string]string{"1": "1", "2": "2", "3": "3", "4": "4"}
	res1, err1 := RedisDB.HSet(Gctx, "ac2a", setarr).Result()
	res2, err2 := RedisDB.HSet(Gctx, "ac1a", setMap).Result()
	fmt.Printf("HSet res:%+v , err:%+v", res1, err1)
	fmt.Printf("HSet res:%+v , err:%+v", res2, err2)

}
