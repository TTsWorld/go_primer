package goroutine

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

// runnext, runq, global runq

func Test0301(t *testing.T) {
	runtime.GOMAXPROCS(1)
	for i := 0; i < 10; i++ { //主协程一直创建协程，go 创建的协程还来不及执行，所有写成创建完后，开始执行 go 启的协程
		i := i //先打印 9，是因为9是最后创建的协程，会被防止到 run-next, run-next有值，则先执行run-next, 再执行 runq
		go func() {
			fmt.Print("A: \t", i)
		}()
	}
	fmt.Println()

	time.Sleep(time.Minute)
}
