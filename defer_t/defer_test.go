package defer_t

import (
	"fmt"
	"testing"
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

func TestDefer3(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered")
		}
	}()

	var a []int32
	a = append(a, 1)
	fmt.Println(a[2])

}
