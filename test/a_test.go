package test

import (
	"fmt"
	"testing"

	"github.com/mohae/deepcopy"
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
	a := &A{1, 1, C{1}}
	b := deepcopy.Copy(a).(*A)
	b.A = 2
	b.C1.C = 2
	fmt.Println(a)
	fmt.Println(b)

}
