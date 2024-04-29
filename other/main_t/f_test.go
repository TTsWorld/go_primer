package main

import (
	"fmt"
	"sync"
	"testing"
)

// 用两个 channel。协程A 从 chan 1 里面读， 协程A读完以后 往 chan 2 里面写， 协程B 再从 chan 2 里面读，读完再往 chan1 里面写。
// 这里 chan1 一定要用带缓冲的 channel, 因为 chan1 被读以后不会立马有 协程来写，而是要等待协程 A 打印以后 写入 chan2, 协程B 读完 chan2 以后打印, 才会去写 chan2

func TestName(t *testing.T) {

	wg := sync.WaitGroup{}
	c1 := make(chan int, 1)
	c2 := make(chan int)
	c3 := make(chan int)
	wg.Add(3)

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			<-c1
			fmt.Println("A")
			c2 <- 1
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			<-c2
			fmt.Println("B")
			c3 <- 1
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			<-c3
			fmt.Println("C")
			c1 <- 1
		}
	}()

	c1 <- 1
	wg.Wait()
}
