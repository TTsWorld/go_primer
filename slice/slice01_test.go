package slice

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	var s []int
	fmt.Println(len(s))
	s = append(s, 1)
	fmt.Println(s)
	fmt.Println(len(s))
}

func TestSlice01(t *testing.T) {
	//声明并初始化
	s1 := make([]int, 10)                      // 这种形式10 是 capacity，slice 会被初始化为含有10 个默认值0的 slice。
	fmt.Printf("%+v, len(s1)=%d", s1, len(s1)) //[0 0 0 0 0 0 0 0 0 0], 10
	s2 := make([]int, 0, 10)                   // 这种形式len传 0，则不写入值则slice 不会有值
	fmt.Printf("%+v", s2)                      //[]
}

// slice 天坑: 切片和子切片共享底层数组, 不要使用子切片修改切片的数据
func TestSlice02(t *testing.T) {

	//声明并初始化
	s1 := make([]int, 10)    // 这种形式10 是 capacity，slice 会被初始化为含有10 个默认值0的 slice。
	fmt.Printf("%+v", s1)    //[0 0 0 0 0 0 0 0 0 0]
	s2 := make([]int, 0, 10) // 这种形式len传 0，则不写入值则slice 不会有值
	fmt.Printf("%+v", s2)    //[]
}

// == slice作为函数参数的 case
// 传递 slice
func modifySlice(s []int) {
	s = append(s, 4, 5, 6)
}

func TestFunc01(t *testing.T) {
	originalSlice := []int{1, 2, 3}
	modifySlice(originalSlice)
	fmt.Println(originalSlice) // 输出: [1 2 3]
}

// 传递 slice 的指针
func modifySlicePtr(s *[]int) {
	*s = append(*s, 4)
}

func TestFunc02(t *testing.T) {
	originalSlice := []int{1, 2, 3}
	fmt.Printf("%p\n", originalSlice)     //0x140000b2018
	fmt.Printf("%p\n", &originalSlice)    //0x140000a20c0
	fmt.Printf("%p\n", &originalSlice[0]) //0x140000b2018
	fmt.Printf("%v\n", len(originalSlice))
	fmt.Printf("%v\n", cap(originalSlice))
	modifySlicePtr(&originalSlice)
	//如果你传递的是一个slice的指针，在函数内部使用append操作会影响原始的slice，因为你直接修改了底层数组。
	fmt.Println(originalSlice) // 输出: [1 2 3 4 5 6]
}

// 传递数组
func modifyArray(arr [3]int) {
	arr[0] = 99
}

func TestFunc04(t *testing.T) {
	originalArray := [3]int{1, 2, 3}
	modifyArray(originalArray)
	//当你传递一个数组给函数时，会发生数组的值拷贝，函数内部对数组的修改不会影响原始数组。
	fmt.Println(originalArray) // 输出: [1 2 3]
}
