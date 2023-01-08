package map_t

import (
	"fmt"
	"testing"
)

type A struct {
	a int
}

func TestMap(t *testing.T) {
	b := []A{{1}, A{2}, A{3}, {4}}
	bMap := make(map[int]*A)
	for k, bi := range b {
		println(&bi)
		println(&k)
		bMap[k] = &bi
	}
	fmt.Printf("%+v", bMap)

}

func TestMapForRandom(t *testing.T) {
	aMap := map[int]struct{}{}
	aMap[1] = struct{}{}
	aMap[2] = struct{}{}
	aMap[3] = struct{}{}
	aMap[4] = struct{}{}
	aMap[5] = struct{}{}
	aMap[6] = struct{}{}
	aMap[7] = struct{}{}
	for k, _ := range aMap {
		println(k)
	}

}
