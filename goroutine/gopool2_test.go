package goroutine

import (
	"fmt"
	"sync"
	"testing"
)

// 作业队列
type Job struct {
	Payload func()
}

// 协程池
type Pool struct {
	queue   chan Job
	wg      sync.WaitGroup
	workers []*worker
}

// 工作协程
type worker struct {
	pool *Pool
	id   int
}

// 创建新的协程池
func NewPool(maxWorkers int) *Pool {
	p := &Pool{
		queue:   make(chan Job, maxWorkers),
		workers: make([]*worker, maxWorkers),
	}

	// 启动所有工作协程
	for i := 0; i < maxWorkers; i++ {
		p.workers[i] = &worker{
			pool: p,
			id:   i,
		}
		go p.workers[i].start()
	}
	return p
}

// 工作协程的main函数
func (w *worker) start() {
	for job := range w.pool.queue {
		w.pool.wg.Add(1)
		job.Payload()
		w.pool.wg.Done()
	}
}

// 向协程池提交作业
func (p *Pool) Submit(job Job) {
	p.queue <- job
}

// 阻塞等待所有作业完成
func (p *Pool) Wait() {
	p.wg.Wait()
}

// 关闭协程池
func (p *Pool) Close() {
	close(p.queue)
	p.Wait()
}

// 示例作业
func exampleJob(id int) {
	fmt.Printf("Worker %d executing job\n", id)
}

func TestName(t *testing.T) {
	// 创建协程池
	pool := NewPool(5)

	// 提交多个作业
	for i := 0; i < 20; i++ {
		job := Job{
			Payload: func() {
				exampleJob(i)
			},
		}
		pool.Submit(job)
	}

	// 等待所有作业完成
	pool.Close()
}
