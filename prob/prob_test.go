package prob

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func SamplingWithoutReplacement(arr []int64, times int) []int {
	lenA := len(arr)
	rand.Seed(time.Now().UnixNano())
	result := []int{}

	findTmp := func(tmp int, arr []int) bool {
		for _, v := range arr {
			if tmp == v {
				return true
			}
		}
		return false
	}

	for {
		tmp := rand.Intn(lenA)
		if findTmp(tmp, result) {
			continue
		} else {
			result = append(result, tmp)
		}
		if len(result) == times {
			return result
		}
	}
}

func Test(t *testing.T) {
	a := make([]int64, 50)
	b := SamplingWithoutReplacement(a, 20)
	fmt.Println(len(b))

}
