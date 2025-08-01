package cast

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	"strings"
	"testing"
)

func TestCast(t *testing.T) {
	s := []string{"1", "2", "3"}
	s1 := cast.ToSlice(s)
	fmt.Printf("%+v", s1)

	strSlice := []string{"hello", "world"}
	// 使用 spf13/cast 包将 []string 转换为 []interface{}
	ifaceSlice, err := cast.ToSliceE(strSlice)
	if err != nil {
		fmt.Println("转换失败:", err)
		return
	}

	fmt.Println(ifaceSlice) // 输出：[]interface {}{"hello", "world"}
}

func Marshal(value interface{}) string {
	result, err := json.Marshal(value)
	if err != nil {
		return ""
	}
	return string(result)
}

func Unmarshal(data string, v interface{}) {
	err := json.Unmarshal([]byte(data), v)
	if err != nil {
	}

}

func Test22(t *testing.T) {
	m := make(map[string]interface{})
	m["id"] = 1873763436359471105
	a := Marshal(m)
	println(a)

	var b map[string]interface{}
	decoder := json.NewDecoder(strings.NewReader(a))
	decoder.UseNumber()
	err := decoder.Decode(&b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cast.ToString(b["id"]))
}

type Binlog struct {
	Database string                 `json:"database"`
	Table    string                 `json:"table"`
	LogType  string                 `json:"type"`
	Ts       int                    `json:"ts"`
	Xid      int                    `json:"xid"`
	Data     map[string]interface{} `json:"data"`
	OldData  map[string]interface{} `json:"old"`
}

func TestName3(t *testing.T) {
	//var binlog Binlog
	//err := json.Unmarshal(text, &binlog)
	//if err != nil {
	//	return
	//}
	//fmt.Println(binlog)
	//
	text := ""
	var binlog Binlog
	decoder := json.NewDecoder(strings.NewReader(text))
	decoder.UseNumber() // 启用 json.Number
	err := decoder.Decode(&binlog)
	if err != nil {
		return
	}

}
