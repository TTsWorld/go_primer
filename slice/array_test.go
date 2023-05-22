package slice

import (
	"fmt"
	"testing"
)

func TestArr01(t *testing.T) {

	//声明并初始化
	a1 := [3]int{1, 2, 3}
	fmt.Printf("%+v", a1) //[1 2 3]

	//声明，数组会被初始化为 0 值
	var a2 [3]int
	fmt.Printf("%+v", a2) //[0 0 0]
}
