package ants_pool

// 示例代码：使用 ants.Pool 处理任务

import (
	"github.com/panjf2000/ants/v2"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {

	pool, _ := ants.NewPool(1000) // 创建容量为 1000 的池
	defer pool.Release()          // 程序退出时释放资源

	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		_ = pool.Submit(func() { // 提交任务
			time.Sleep(time.Second * 1)
			defer wg.Done()
			// 执行具体任务（如处理请求）
		})
	}
	wg.Wait()
}
