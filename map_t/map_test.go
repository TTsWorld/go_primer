package map_t

import (
	"encoding/json"
	"fmt"
	"math/rand"
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

func TestMapAppend(t *testing.T) {
	a := make([]map[string]interface{}, 0)
	for i := 0; i < 10; i++ {
		t := map[string]interface{}{"a": i}
		a = append(a, t)
	}
	fmt.Println(a)
}

var EcsInstanseId = []string{
	"i-t4n317o78w5nmej9jvgu",
	"i-t4n317o78w5nmej9jvgt",
	"i-t4n7bqi0f5qntqwd09so",
	"i-t4n7bqi0f5qntqwd09sp",
	"i-t4nb2htfd1lvjs335e1c",
	"i-t4n0fketd6fmatdcvdrc",
	"i-t4n0pyibci9cyq59njka",
	"i-t4n4d6gh513ac69lxynj",
	"i-t4n0fketd6fmatdcvdra",
	"i-t4n57y7u94v29li2u9by",
	"i-t4n0fketd6fmatdcvdrb",
	"i-t4n1mp121g05m0ohlf0q",
}

func TestMapAppend2(t *testing.T) {
	arr := make([]map[string]string, 0)

	for _, instanceId := range EcsInstanseId {
		t := map[string]string{"instanceId": instanceId}
		arr = append(arr, t)
	}
	v, _ := json.Marshal(arr)
	fmt.Println(string(v))
}

func TestGetUnExistsElement(t *testing.T) {
	arr := make(map[string]string)

	fmt.Println(arr["a"])
}

// 测试 map[string]map[string]int 结构添加过程中是否需要创建
func TestMapAdd(t *testing.T) {
	println("23:58:59.000" > "24:58:39.000")
	for i := 0; i < 10; i++ {
		randN := rand.Intn(1000)
		println(randN)

	}
}

func TestMapAdd2(t *testing.T) {
	m := make(map[string][]int)
	for i := 0; i < 10; i++ {
		m["a"] = append(m["a"], i)
	}
	for i := 0; i < 10; i++ {
		m["b"] = append(m["b"], i)
	}
	fmt.Printf("%+v", m)

}
