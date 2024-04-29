package goroutine

//
//import (
//	"fmt"
//	"sync"
//	"testing"
//	"time"
//)
//
//// 定义一个任务类型
//type Task func()
//
//// 协程池结构体
//type Pool struct {
//	// 存储任务的通道
//	taskChan chan Task
//	// 控制最大协程数的信号量
//	semaphore *sync.WaitGroup
//	// 关闭通道的信号量
//	doneChan chan struct{}
//}
//
//// 新建协程池
//func NewPool(taskChan chan Task) *Pool {
//	return &Pool{
//		taskChan:  taskChan,
//		semaphore: &sync.WaitGroup{},
//		doneChan:  make(chan struct{}),
//	}
//}
//
//// 执行任务的协程
//func (p *Pool) worker() {
//	defer p.semaphore.Done() // 当协程执行完毕，减少信号量计数
//
//	for task := range p.taskChan {
//		task() // 执行任务
//	}
//}
//
//// 将任务添加到协程池
//func (p *Pool) AddTask(task Task) {
//	p.taskChan <- task
//	p.semaphore.Done() // 增加信号量计数，允许新的协程开始
//}
//
//// 关闭协程池
//func (p *Pool) Close() {
//	close(p.taskChan)  // 关闭任务通道
//	p.semaphore.Wait() // 等待所有协程执行完毕
//	close(p.doneChan)  // 关闭完成通道
//}
//
//// 协程池任务执行函数
//func workerTask(id int, taskChan chan Task, wg *sync.WaitGroup) {
//	defer wg.Done()
//	for task := range taskChan {
//		fmt.Printf("Worker %d is starting...\n", id)
//		task() // 执行任务
//		fmt.Printf("Worker %d is done.\n", id)
//	}
//}
//
//func TestGoPool(t *testing.T) {
//
//	// 定义协程池大小
//	poolSize := 5
//	// 创建任务通道
//	taskChan := make(chan Task, poolSize)
//
//	// 创建协程池
//	pool := NewPool(poolSize, taskChan)
//
//	// 启动协程池中的协程
//	for i := 0; i < poolSize; i++ {
//		go pool.worker()
//	}
//
//	// 添加任务到协程池
//	for i := 0; i < 10; i++ {
//		task := func() {
//			fmt.Printf("Task %d is running...\n", i)
//			time.Sleep(1 * time.Second)
//		}
//		pool.AddTask(task)
//	}
//
//	// 等待所有任务执行完毕
//	pool.Close()
//
//	// 阻塞主协程，直到协程池关闭
//	<-pool.doneChan
//}
