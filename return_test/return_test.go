package return_test

import (
	"fmt"
	"testing"
)

func TestReturn(t *testing.T) {
	r1 := return01()
	fmt.Println(r1.A)

	r2 := return02()
	fmt.Println(r2.A)

	r3 := return03()
	fmt.Println(r3)

	r4 := return04()
	fmt.Println(r4)
}

type R struct {
	A int
	B int
}

func return01() (r R) {
	return
}

func return02() (r *R) {
	r = &R{}
	return
}

func return03() (r []R) {
	return
}

func return04() (r []*R) {
	return
}
