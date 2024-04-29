package concurrent

import (
	"fmt"
	cmap "github.com/orcaman/concurrent-map"
	"sync"
	"testing"
	"time"
)

// sync.Map不同的是，cmap使用了分段锁的方式，先看下cmap主要的数据结构
func TestConcurrent01(t *testing.T) {
	m := cmap.New()
	fmt.Printf("m:%v \n", m)
	m.Set("a", 1)
	fmt.Printf("m:%v \n", m)

	if tmp, ok := m.Get("a"); ok {
		fmt.Printf("tmp:%v, ok:%v \n", tmp, ok)
	}
	m.Remove("b")
	fmt.Printf("m:%#v \n", m)

}

func A(a int) int {
	a++
	return a
}

func B(b int) int {
	b++
	return b
}

func testConcurrent(param int) {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func(p int) int {
		defer wg.Done()
		p++
		return p
	}(param)
	//time.Sleep(time.Second)
	go func(p int) int {
		defer wg.Done()
		p++
		return p
	}(param)
	//time.Sleep(time.Second)

	println(param)

	wg.Wait()
}

func testConcurrent2(param int) {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		param++
	}()
	//time.Sleep(time.Second)
	go func() {
		defer wg.Done()
		param++
	}()
	//time.Sleep(time.Second)

	println(param)

	wg.Wait()
}

func testConcurrent3(param int) {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		param++
	}()
	time.Sleep(time.Second)
	go func() {
		defer wg.Done()
		param++
	}()
	//time.Sleep(time.Second)

	println(param)
	wg.Wait()
}

func testConcurrent4(param *int) {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		*param++
	}()

	go func() {
		defer wg.Done()
		*param++
	}()

	println(*param)
	wg.Wait()
}

func TestC2(t *testing.T) {

	var a = 2
	testConcurrent4(&a)
	testConcurrent(2)
	testConcurrent2(2)
	testConcurrent3(2)

}
