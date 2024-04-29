package main

import (
	"fmt"
	"testing"

	"go.uber.org/atomic"
)

var int01 atomic.Int32
var val01 atomic.Value

func TestAtomicValue(t *testing.T) {
	type T struct{ a, b, c int }
	x := T{1, 2, 3}
	y := T{4, 5, 6}
	z := T{7, 8, 9}
	a := T{7, 7, 7}

	var v atomic.Value
	v.Store(x)           //存入 x
	fmt.Println(v)       //打印 v
	old := v.Swap(y)     //用 y 替换 x，获得老值和存入新值
	fmt.Println(v)       //  新值 :{{{4 5 6}} []}
	fmt.Println(old.(T)) // 老值: {1 2 3}

	swapped := v.CompareAndSwap(x, y) //  看 v 里的值跟 x 是否相等，相等返回 true，并替换为新的值 y ; 如果不相等，不进行替换，并返回 false
	fmt.Println(swapped, v)           // false {{{4 5 6}} []}
	swapped = v.CompareAndSwap(y, z)
	fmt.Println(swapped, v) // true {{{7 8 9}} []}

	nowT := v.Load().(T)
	fmt.Println(nowT) // {7 8 9}

	v.Store(a)
	fmt.Println(v) // {7 8 9}

	fmt.Println(v.Load().(T))

}
