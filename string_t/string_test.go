package string_t

import (
	"fmt"
	"strconv"
	"strings"
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

func TestSplit2(t *testing.T) {
	s := ""
	aa := strings.Split(s, ",")
	fmt.Printf("%v, %v\n", aa, len(aa))
}
