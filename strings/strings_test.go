package strings

import (
	"fmt"
	"strings"
	"testing"
)

func Test01(t *testing.T) {
	a := "a|b|c|d"
	bb := strings.Split(a, "|")
	fmt.Printf("%+v", bb)
}

func Test02(t *testing.T) {
}
