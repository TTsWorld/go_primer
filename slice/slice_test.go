package slice

import (
	"fmt"
	"testing"
)

func TestSlice01(t *testing.T) {

	//声明并初始化
	s1 := make([]int, 10)    // 这种形式10 是 capacity，slice 会被初始化为含有10 个默认值0的 slice。
	fmt.Printf("%+v", s1)    //[0 0 0 0 0 0 0 0 0 0]
	s2 := make([]int, 0, 10) // 这种形式len传 0，则不写入值则slice 不会有值
	fmt.Printf("%+v", s2)    //[]
}

// slice 天坑: 切片和子切片共享底层数组, 不要使用子切片修改切片的数据
func TestSlice02(t *testing.T) {

	//声明并初始化
	s1 := make([]int, 10)    // 这种形式10 是 capacity，slice 会被初始化为含有10 个默认值0的 slice。
	fmt.Printf("%+v", s1)    //[0 0 0 0 0 0 0 0 0 0]
	s2 := make([]int, 0, 10) // 这种形式len传 0，则不写入值则slice 不会有值
	fmt.Printf("%+v", s2)    //[]
}
