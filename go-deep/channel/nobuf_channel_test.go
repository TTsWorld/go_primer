package channel

import (
	"fmt"
	"testing"
)

func TestNoBufChannel01(t *testing.T) {
	ch := make(chan int)
	ch <- 1
	test01(ch)
}
func test01(ch chan int) {
	a := <-ch
	fmt.Println(a)

}
