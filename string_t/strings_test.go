package string_t

import (
	"fmt"
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	s := "a"
	ss := strings.Split(s, "|")
	fmt.Printf("ss:%+v, len(ss):%v\n", ss, len(ss))

}
