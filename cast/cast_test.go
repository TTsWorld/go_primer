package cast

import (
	"fmt"
	"github.com/spf13/cast"
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
