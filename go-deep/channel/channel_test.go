package channel

import (
	"context"
	"fmt"
	"sync"
	"testing"
)

func Test01(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
	}()

	for {
		select {
		case <-ch1: // 当 ch1 被设置为 nil 后，将不会到达此分支了。
			func() {
				println("ch11111")
			}()
		case <-ch2:
			func() {
				println("ch222222")
			}()
		}
	}

}

var ch = make(chan int, 100)

func TestACh02(t *testing.T) {
	wg := sync.WaitGroup{}
	ctx := context.Background()
	wg.Add(2)
	handle := func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("done done done")
				return
			case task, ok := <-ch:
				if !ok {
					fmt.Println("read ch error")
				} else {
					doTask(task)
				}
			}
		}
	}
	handle()
	handle()
	ch <- 15
	ch <- 15
	ch <- 15

}

func doTask(a int) {
	println("dotask|| a a a %v", a)
}
