package go_redis

import (
	"bufio"
	"fmt"
	"github.com/spf13/cast"
	"io"
	"os"
	"strings"
	"testing"
	"time"
)

func TestMRedis01(t *testing.T) {

	after, err := RedisLPushLimitedV2("l", cast.ToString(time.Now().Unix()), 10, time.Minute*10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(after)

}

func Test2(t *testing.T) {
	doConvert()

}

func doConvert() {
	client, ctx, cancel := Redis()
	defer cancel()
	rdbFile, err := os.Open("/tmp/rdb.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	reader := bufio.NewReader(rdbFile)
	cnt := 0

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		key := strings.TrimSpace(string(line))
		cnt++
		if key == "" {
			continue
		}

		var year, uid int
		_, err = fmt.Sscanf(key, "birthdaywall_has_birthday:{%d}_{%d}", &year, &uid)
		if err != nil {
			fmt.Printf("doConvert||Sscanf failed, key:%s, err:%v", key, err)
			continue
		}
		newKey := fmt.Sprintf("birthdaywall_has_birthday:{%d}_{%d}", uid, year)
		client.Del(ctx, key)
		if err != nil {
			fmt.Printf("doConvert||Sscanf failed, key:%s, err:%v", key, err)
			continue
		}
		client.Set(ctx, newKey, "1", time.Hour*24*30*365)
		if err != nil {
			fmt.Printf("doConvert||Sscanf failed, key:%s, err:%v", key, err)
			continue
		}
	}
	fmt.Printf("doConvert||process cnt:%d", cnt)
}
