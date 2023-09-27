package concurrent

import (
	"fmt"
	cmap "github.com/orcaman/concurrent-map"
	"testing"
)

// sync.Map不同的是，cmap使用了分段锁的方式，先看下cmap主要的数据结构
func TestConcurrent01(t *testing.T) {
	m := cmap.New()
	fmt.Printf("m:%v \n", m)
	m.Set("a", 1)
	fmt.Printf("m:%v \n", m)

	if tmp, ok := m.Get("a"); ok {
		fmt.Printf("tmp:%v, ok:%v \n", tmp, ok)
	}
	m.Remove("b")
	fmt.Printf("m:%#v \n", m)

}
