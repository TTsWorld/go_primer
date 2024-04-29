package other

import (
	"fmt"
	"sync"
	"testing"
)

// 2个协程, 1 个打印 1, 3, 5, 7, 9 ;1 个打印 2, 4, 6, 8, 10
// 最终输出 1，2，3，4，5，6，7，8，9，10
func TestK0(t *testing.T) {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)

	f1 := func(arr []int) {
		defer wg.Done()
		for _, v := range arr {
			ch <- v
			println(v)
		}
	}
	f2 := func(arr []int) {
		defer wg.Done()
		for _, v := range arr {
			<-ch
			println(v)
		}
	}

	go f1([]int{1, 3, 5})
	go f2([]int{2, 4, 6})

	wg.Wait()
}

// 2个协程, 1 个打印 1, 3, 5, 7, 9 ;1 个打印 2, 4, 6, 8, 10
// 最终输出 1，2，3，4，5，6，7，8，9，10
func TestK(t *testing.T) {
	ch := make(chan int)
	ch2 := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)

	f1 := func(arr []int) {
		defer wg.Done()
		for _, v := range arr {
			ch <- v
			println(v)
			<-ch2 //
		}
	}
	f2 := func(arr []int) {
		defer wg.Done()
		for _, v := range arr {
			<-ch
			println(v)
			ch2 <- v
		}
	}

	go f1([]int{1, 3, 5, 7, 9})
	go f2([]int{2, 4, 6, 8, 10})

	wg.Wait()
}

// 2个协程, 1 个打印 1, 3, 5, 7, 9 ;1 个打印 2, 4, 6, 8, 10
// 最终输出 1，2，3，4，5，6，7，8，9，10
func TestK1(t *testing.T) {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)

	f1 := func(arr []int) {
		defer wg.Done()
		for _, v := range arr {
			ch <- v
			println(v)
		}
	}
	f2 := func(arr []int) {
		defer wg.Done()
		for _, v := range arr {
			<-ch
			println(v)
		}
	}

	go f1([]int{1, 3, 5})
	go f2([]int{2, 4, 6})

	wg.Wait()
}

//上面这个题，考虑下如何合适的 close 协程

func TestKK(t *testing.T) {
	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{1, 2, 3}
	b2 := (arr1 == arr2) //能否运行，如果能，b2韵值是什么
	fmt.Println(b2)

	slice1 := []int{1, 2, 3}
	slice2 := []int{1, 2, 3}
	//b3 := (slice1 == slice2) //能否运行，如果能，b3的值是什么
	fmt.Println(slice1, slice2) //Invalid operation: slice1 == slice2 (the operator == is not defined on []int)

	var m1 map[string]int
	m1["one"] = 1 //panic: assignment to entry in nil map [recovered]
	m2 := map[string]int{"one": 1}
	//b1 := (m1 == m2) //能否运行，如果能，b1的值是什么

	fmt.Println(m2) //Invalid operation: m1 == m2 (the operator == is not defined on map[string]int)

	m3 := make(map[string]string)
	m3["one"] = "1"
	m3["two"] = "2"
	m3["three"] = "3"
	for key := range m3 {
		fmt.Println(key) //能否运行，如果能，输出什么, 乱序输出
	}
}
