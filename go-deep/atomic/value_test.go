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
	v.Store(x)
	fmt.Println(v)
	old := v.Swap(y)
	fmt.Println(v)       // {{{4 5 6}} []}
	fmt.Println(old.(T)) // {1 2 3}

	swapped := v.CompareAndSwap(x, y)
	fmt.Println(swapped, v) // false {{{4 5 6}} []}
	swapped = v.CompareAndSwap(y, z)
	fmt.Println(swapped, v) // true {{{7 8 9}} []}

	nowT := v.Load().(T)
	fmt.Println(nowT) // {7 8 9}

	v.Store(a)
	fmt.Println(v) // {7 8 9}

	fmt.Println(v.Load().(T))

}
