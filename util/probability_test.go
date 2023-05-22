package main

import (
	"math/rand"
	"testing"
	"time"
)

type SingleProbInfo struct {
	offset int64
	len    int64
}

func drawOnce(probs []int64) (code int) {
	probsSlice := make([]SingleProbInfo, 0, len(probs))
	var sumLen int64 = 0

	for _, c := range probs {
		a := SingleProbInfo{
			offset: sumLen,
			len:    c,
		}
		probsSlice = append(probsSlice, a)
		sumLen += c
	}

	awardIndex := rand.Int63n(sumLen)
	for idx, u := range probsSlice {
		if u.offset+u.len > awardIndex {
			return idx
		}
	}
	return 0
}

func main() {
	rand.Seed(time.Now().UnixNano())
	t := []int64{10, 10, 10, 70}
	var a, b, c, d int
	for i := 0; i < 10000; i++ {
		res := drawOnce(t)
		switch res {
		case 0:
			a++
		case 1:
			b++
		case 2:
			c++
		case 3:
			d++
		}

	}
	println(a)
	println(b)
	println(c)
	println(d)
}

func calProb(probs []float64) int32 {
	probSum := float64(0)
	for i := 0; i < len(probs); i++ {
		prob := probs[i]
		if i == len(probs)-1 {
			prob = 1 - prob
		}
		probSum += probSum + prob
	}
	res := int32((probSum - float64(len(probs))) / float64(probSum) * 100)
	println(res)
	if res < 50 {
		res = 50
	}

	return res
}
func Test02(t *testing.T) {
	probs := []float64{0.9, 0, 9, 0.9, 0.9}
	println(calProb(probs))

}
