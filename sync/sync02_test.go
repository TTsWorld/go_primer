package sync

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// data race
//对于多 goroutine 访问共享变量，需要使用锁货 chan 对变量进行保护

var Wait sync.WaitGroup
var Count int

func Test0201(t *testing.T) {

	for i := 0; i < 2; i++ {
		Wait.Add(1)
		go run()
	}
	Wait.Wait()
	fmt.Println(Count)

}
func run() {
	for i := 0; i < 2; i++ {
		//Count++
		Count = Count + 1
		time.Sleep(time.Nanosecond)
	}

	Wait.Done()
}

// ==== data race 02
// 接口赋值的时候，可能存在iface 中 type赋值成功，但value 未赋值成功的情况
type IceCreamMaker interface {
	Hello()
}
type Ben struct {
	id   int
	name string
}

func (b *Ben) Hello() {
	fmt.Printf("Ben says, hello, myname is %s\n", b.name)
}

type Jerry struct {
	name string
}

func (j *Jerry) Hello() {
	fmt.Printf("jerry says, hello, myname is %s\n", j.name)
}

func Test0202(t *testing.T) {
	var ben = &Ben{id: 10, name: "Ben"}
	var jerry = &Jerry{name: "jerry"}
	var maker IceCreamMaker = ben
	var loop0, loop1 func()

	loop0 = func() {
		maker = ben
		go loop1()
	}
	loop1 = func() {
		maker = jerry
		go loop0()
	}

	go loop0()

	for {
		maker.Hello()
	}

}
