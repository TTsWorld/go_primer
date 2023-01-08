package string_t

import (
	"fmt"
	"strconv"
	"testing"
)

func TestString(t *testing.T) {
	a := 4
	if a == 4 {
		a, b := 3, 1
		a = a
		b = b
		println(a)
	}
	println(a)
	return
	a, err := strconv.Atoi("xxx1")
	fmt.Printf("%+v\n", err)
	println(a)

}
