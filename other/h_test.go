package other

import (
	"sync"
	"testing"
	"time"
)

//go 实现协程池 - 地平线
//最大并发数
//成员方法  addTask 提交任务
// 排队
// Run()
// initPool()

type task func()

type Pool struct {
	Size       int //最大并发数
	taskChan   chan task
	resizeChan chan struct{}
	wg         *sync.WaitGroup
}

func (p *Pool) AddTask(t task) {
	p.taskChan <- t
}

func (p *Pool) InitWorker() {
	defer p.wg.Done()
	for {
		select {
		case t := <-p.taskChan:
			t()
		}

	}

}

func InitPool(size int) *Pool {
	return &Pool{
		Size:     size,
		taskChan: make(chan task),
		wg:       &sync.WaitGroup{},
	}
}

func (p *Pool) ResizePool(size int) {
	if size < p.Size {
		for i := 0; i < p.Size-size; i++ {
			p.Size--
			p.resizeChan <- struct{}{}
		}
	}
}

func TestGoroutinePool(t *testing.T) {
	//1 init pool
	pool := InitPool(5)
	pool.wg.Add(pool.Size)

	// 2 init worker
	for i := 0; i < pool.Size; i++ {
		go pool.InitWorker()
	}
	time.Sleep(time.Second)

	// 3 add task
	for i := 0; i < 5; i++ {
		pool.AddTask(func() {
			println("hello")
		})
	}

	pool.wg.Wait()
}
