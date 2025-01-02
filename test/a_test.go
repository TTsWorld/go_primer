package test

import (
	"fmt"
	"testing"
)

func TestAA(t *testing.T) {
	luckyGiftId := 10
	getStageIdx := func(idx int) int {
		if idx < 5 {
			return idx
		}
		luckyGiftId = 20
		return idx % 5
	}
	fmt.Println(getStageIdx((10)))
	fmt.Println(luckyGiftId)
}

type C struct {
	C int
}

type A struct {
	A  int
	B  int
	C1 C
	C2 *C
}

func Test02(t *testing.T) {
	var a = 1
	var b = 2
	c := fmt.Sprintf("aa:%[1]d:%[2]d", a, b)
	fmt.Println(c)

}
