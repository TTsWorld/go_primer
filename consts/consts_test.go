package consts

import (
	"fmt"
	"testing"
)

type Status int

const (
	Status01 Status = 0
	Status02 Status = 1
	Status03 Status = 2
	Status04 Status = 3
)

func TestConsts(t *testing.T) {
	s := Status02
	fmt.Printf("%d", int(s))

	s2 := 1
	fmt.Print(s2)
}
