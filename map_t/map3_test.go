package map_t

import (
	"fmt"
	"testing"
)

func TestMap0301(t *testing.T) {
	m := make(map[int]int)
	m[1] = 1
	m[2] = 3
	m[3] = 3
	fmt.Println(m[1])
	fmt.Println(m[5])

	m2, ok2 := m[2]
	m6, ok6 := m[6]
	fmt.Println(m2, ok2)
	fmt.Println(m6, ok6)

}
