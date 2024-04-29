package other

import (
	"math"
	"testing"
	"time"
)

func shunzi(args []int) bool {
	m := make(map[int]struct{})
	max := math.MinInt
	min := math.MaxInt

	for _, arg := range args {
		if arg == 0 {
			continue
		}

		if _, ok := m[arg]; ok {
			return false
		}
		m[arg] = struct{}{}
		if arg > max {
			max = arg
		}
		if arg < min {
			min = arg
		}
	}
	if max-min < 5 {
		return true
	}

	return false
}

func TestA(t *testing.T) {
	/*
	   1,2,3,4,5,6,7,8,9,10,11,12,13
	*/
	input := []int{1, 2, 3, 4, 5}
	println(shunzi(input))
	input2 := []int{1, 2, 2, 4, 5}
	println(shunzi(input2))

}

func solution() int {
	randFunc := func() int { return 0 } // 0: p, 1: (1-p)

	/*
		0-0 = p*p
		0-1 = p(1-p)
		1-1 = (1-p)(1-p)
		1-0 = (1-p)p
	*/

	for {
		r1 := randFunc()
		r2 := randFunc()
		if r1 == 0 && r2 == 1 {
			return 0
		} else if r1 == 1 && r2 == 0 {
			return 1
		} else {
			time.Sleep(time.Millisecond)
			continue
		}
	}

}

func TestB(t *testing.T) {
	solution()

}
