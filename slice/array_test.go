package slice

import (
	"fmt"
	"testing"
)

func TestArr01(t *testing.T) {

	//声明，数组会被初始化为 0 值
	var a2 [3]int
	fmt.Printf("%+v\n", a2) //[0 0 0]

	//声明初始化-不赋值
	a0 := [3]int{}
	fmt.Printf("%+v\n", a0) //[1 2 3]

	//声明并初始化
	a1 := [3]int{1, 2, 3}
	fmt.Printf("%+v\n", a1) //[1 2 3]

	//数组比较
	a11 := [3]int{1, 2, 3}
	a22 := [3]int{1, 2, 3}
	fmt.Printf("%+v\n", a11 == a22) //true

	//数组比较,不同长度的数组，不是同一个类型，不能进行比较
	a111 := [3]int{1, 2, 3}
	a222 := [4]int{1, 2, 3}
	a111 = a111
	a222 = a222
	//fmt.Printf("%+v\n", a111 == a222) //无法编译，数据类型不同

}
