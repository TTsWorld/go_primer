package defer_t

import (
	"fmt"
	"testing"
	"time"
)

func TestDefer(t *testing.T) {
	var w [3]struct{}

	for i := range w {
		defer func() {
			fmt.Println(i)
		}()
	}
}

// return 后面的 defer 由于没有注册，不会被执行
func TestDefer2(t *testing.T) {
	defer func() {
		time.Sleep(time.Second * 3)
		println("before return")
	}()

	if true {
		println("during return")
		return
	}

	defer func() {
		println("after return")
	}()
}

// 测试 recover 用法
func TestDefer3(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered")
		}
	}()

	var a []int32
	a = append(a, 1)
	fmt.Println(a[2])
	println(1111) //因为 panic 不打印这行
}

// 因为再 defer 执行前 panic 了，所以 defer 不会执行
func TestDefer4(t *testing.T) {
	defer func() {
		fmt.Println("defer not execute")
	}()

	var a []int32
	a = append(a, 1)
	fmt.Println(a[2])
	println(1111) //因为 panic 不打印这行
}
