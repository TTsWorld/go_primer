package test

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	probs := []int64{1499, 99, 990, 1400}
	probSum := int64(0)
	for i := 0; i < len(probs); i++ {
		prob := probs[i]
		if i == len(probs)-1 {
			prob = 40000 - prob
		}
		probSum = probSum + prob
	}
	fmt.Printf("probSum:%v, probs:%v", probSum, probs)
	res := int32(float64(int64(len(probs))-probSum) / float64(probSum) * 10000)
	fmt.Printf("res:%v", res)
}
