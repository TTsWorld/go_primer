package channel

import (
	"fmt"
	"testing"
	"time"
)

func TestNoBufChannel01(t *testing.T) {
	ch := make(chan int)
	go test01(ch)
	ch <- 1
	time.Sleep(time.Second * 2)

}
func test01(ch chan int) {
	a := <-ch
	fmt.Println(a)

}
