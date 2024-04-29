package channel

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

// 无缓冲队列的基本操作
func Test0101(t *testing.T) {
	ch := make(chan int, 1)
	ch <- 1
	go func() {
		ch <- 1
		println("xxx")
	}()
	time.Sleep(time.Second * 2)
}

// ====缓冲队列的基本操作
func Test0102(t *testing.T) {
	var ch chan int
	ch = make(chan int, 4)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		ch <- 1
		ch <- 2
		ch <- 3
		ch <- 4
		ch <- 5
		close(ch)
	}()

	fmt.Printf("cap(ch)=%+v\n", cap(ch)) //获取 ch 的容量
	fmt.Printf("len(ch)=%+v\n", len(ch)) //获取 ch 的有效元素个数
	for {
		select {
		case a, ok := <-ch:
			fmt.Println(a)
			fmt.Println(ok)
			if !ok {
				return
			}
		default:
			fmt.Println("dddd")
		}

	}
	wg.Wait()
	fmt.Printf("%+v\n", <-ch) //1
	fmt.Printf("%+v\n", <-ch) //2
	fmt.Printf("%+v\n", <-ch) //3
	fmt.Printf("%+v\n", <-ch) //4
	fmt.Printf("%+v\n", <-ch) //0 //如果 channel 已关闭，则会消费到 0 值
	fmt.Printf("cap(ch)=%+v\n", cap(ch))
	fmt.Printf("len(ch)=%+v\n", len(ch))
}

// ==== channel 关闭的判断
// channel被关闭，则使用多返回值接受 channel 的值时，第二个值为 true
func Test0103(t *testing.T) {
	var ch chan int
	ch = make(chan int, 4)

	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	close(ch)
	_, ok := <-ch
	fmt.Println(ok)
}

// ==== channel关闭后，for-range会自动退出
func Test0104(t *testing.T) {
	var ch chan int
	ch = make(chan int, 4)

	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	go func() {
		for val := range ch {
			fmt.Println(val)
		}
		fmt.Println("goroutine exit")
	}()
	time.Sleep(3 * time.Second)
	close(ch)
	time.Sleep(time.Second) //等待一下 goroutine 中打印退出后的信息，否则直接退出了
}

// ==== channel - select 基本用法
func Test0105(t *testing.T) {
	var ch chan int
	ch = make(chan int, 4)

	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	select {
	case a := <-ch:
		fmt.Println(a)
	}
	go func() {
		for val := range ch {
			fmt.Println(val)
		}
		fmt.Println("goroutine exit")
	}()
	time.Sleep(3 * time.Second)
	close(ch)
	time.Sleep(time.Second) //等待一下 goroutine 中打印退出后的信息，否则直接退出了
}

// ==== channel - select -default 行为
func Test0106(t *testing.T) {
	var ch chan int
	ch = make(chan int, 4)

	select {
	case a := <-ch:
		fmt.Println(a)
	default:
		fmt.Println("xxx")
	}
}

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
