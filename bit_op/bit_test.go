package bit_op

import (
	"fmt"
	"testing"
)

func TestBit01(t *testing.T) {
	a := 10
	b := 0x10
	fmt.Printf("%v \t %b \t %x\n", a, a, a)
	fmt.Printf("%v \t %b \t %x\n", b, b, b)
	fmt.Printf("%v \t %b \t %x\n", ^b, ^b, ^b)
	fmt.Printf("%v \t %b \t %x\n", ^b, ^b, ^b)

}
