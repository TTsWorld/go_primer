package concurrent

import (
	"testing"
	"time"
)

func for1() {
	for i := 0; i < 5; i++ {
		go func(aa int) {
			println(aa)
		}(i)
	}
	time.Sleep(time.Second)
}

func for2() {
	for i := 0; i < 5; i++ {
		go func() {
			println(i)
		}()
	}
	time.Sleep(time.Second)
}
func TestFor(t *testing.T) {
	for1()
	time.Sleep(time.Second)
	println("============")
	for2()
	a := []int{}
	a = append(a, 3)

}

func TestSlice(t *testing.T) {

	var s []*struct{}
	s = append(s, &struct{}{})
	s = append(s, &struct{}{})
	s = append(s, &struct{}{})
	s = append(s, &struct{}{})

	var res []interface{}
	for _, a := range s {
		res = append(res, a)
		println(res)
	}
}
