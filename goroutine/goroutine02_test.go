package goroutine

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// ==========0201, from geektime course
// 空的 select 将永远阻塞
// case 启示:如果你的 goroutine 在从另一个 goroutine 获得结果之前无法取得进展，那么通常情况下，你自己去做这项工作比委托它（go func(){}）更简单

type Tracker struct {
	ch   chan string
	stop chan struct{} //通过 stop 知道 Run()函数何时退出，然后 shutdown 退出
}

func NewTracker() *Tracker {
	return &Tracker{
		ch: make(chan string, 10),
	}

}

func (t *Tracker) Event(ctx context.Context, data string) error {
	select {
	case t.ch <- data:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}

}
func (t *Tracker) Run() {
	for data := range t.ch {
		time.Sleep(1 * time.Second)
		fmt.Println(data)
	}
	t.stop <- struct{}{}
}

func (t *Tracker) Shutdown(ctx context.Context) {
	close(t.ch)
	select {
	case <-t.stop:
	case <-ctx.Done():

	}

}

func Test0201(t *testing.T) {
	tr := NewTracker()
	_ = tr.Event(context.Background(), "event01")
	_ = tr.Event(context.Background(), "event01")
	_ = tr.Event(context.Background(), "event01")
	time.Sleep(3 * time.Second)
	_, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	defer cancel()

}
