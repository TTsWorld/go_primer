package sync

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestCond(t *testing.T) {
	var mu sync.Mutex
	var cond sync.Cond
	cond.L = &mu

	ready := false

	// 等待条件的 goroutine
	go func() {
		mu.Lock()
		for !ready {
			cond.Wait() // 释放锁并等待条件变量
		}
		fmt.Println("Condition met, proceeding...")
		mu.Unlock()
	}()

	// 改变条件的 goroutine
	go func() {
		mu.Lock()
		defer mu.Unlock()
		// 模拟一些工作
		time.Sleep(time.Second)
		ready = true
		cond.Signal() // 通知等待的 goroutine
		fmt.Println("Condition is now true.")
	}()

	// 模拟主 goroutine 继续执行
	fmt.Println("Main goroutine is doing other work.")
	time.Sleep(2 * time.Second)
}
