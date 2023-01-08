package main

import (
	"runtime"
	"time"

	"go.uber.org/atomic"
)

var GV atomic.Value

//func Test(t *testing.T) {
//
//	m := make(map[int]map[int]int)
//
//	m = nil
//	GV.Store(m)
//	println(GV.Load())
//
//	m1 := GV.Load().(map[int]map[int]int)
//	m1[0] = make(map[int]int)
//
//	println(m1[0])
//}
type M map[int]map[int]int

func main() {
	GV.Store(M{})

	go func() {
		for {
			GV.Store(M{})
		}
	}()

	go func() {
		time.Sleep(1 * time.Second)
		for {
			a := GV.Load()
			b := GV.Load().(M)
			a = a
			b = b
		}
	}()

	go func() {
		time.Sleep(10)
		runtime.GC()
	}()
	//	println(m == nil)
	//	GV.Store(m)
	//	println(GV.Load())
	//
	//	m0 := GV.Load()
	//	m1 := m0.(map[int]map[int]int)
	//
	//	println(m0 == nil)
	//	println(m1 == nil)
	time.Sleep(10 * time.Minute)

}
