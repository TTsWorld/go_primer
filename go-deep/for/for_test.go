package _for

import (
	"fmt"
	"testing"
	"time"
)

func TestFor(t *testing.T) {
	for i := 0; i < 10; i++ {
		println(i)
		time.Sleep(1 * time.Second)
		continue
		i = 10

	}
}
func TestA(t *testing.T) {
	aMap := make(map[int]int)
	aMap[1] = 10
	aMap[2] = 20
	for k := range aMap {
		fmt.Println(k)
	}

}
