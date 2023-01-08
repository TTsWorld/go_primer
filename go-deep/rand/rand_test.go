package rand

import (
	"math/rand"
	"testing"
	"time"
)

func TestA(t *testing.T) {
	ret := []int{1, 2, 3, 4, 5, 6, 7, 8}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(ret), func(i, j int) {
		ret[i], ret[j] = ret[j], ret[i]
	})
	println(ret[0])
}
