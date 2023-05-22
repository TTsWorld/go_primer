package deep_copy

import (
	"fmt"
	"github.com/mohae/deepcopy"
	"testing"
)

func Test01(t *testing.T) {

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := deepcopy.Copy(a).([]int)
	b[2] = 10
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)

}
