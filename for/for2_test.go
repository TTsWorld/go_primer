package _for

import (
	"fmt"
	"testing"
	"time"
)

func TestA2(t *testing.T) {

	var s []*int
	for i := 0; i < 10; i++ {
		s = append(s, &i)
	}

	m := make(map[int]*int)
	for i, i2 := range s {
		m[i] = i2
	}
	for i, i2 := range m {
		fmt.Printf("%+v,  %+v", i, i2)
	}

	fmt.Println()
	m2 := make([]*int, 0)
	for _, i2 := range s {
		m2 = append(m2, i2)
	}
	for i, i2 := range m2 {
		fmt.Printf("%+v,  %+v", i, i2)
	}
	fmt.Println()
	now := time.Now()
	fmt.Println(now.Unix())
	fmt.Println(now.UTC().Unix())

}
