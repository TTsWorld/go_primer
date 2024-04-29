package channel

import "testing"

// 缓冲队列的基本操作
func Test0301(t *testing.T) {

	aMap := map[string]interface{}{
		"a": "a",
	}
	ch := make(chan map[string]interface{}, 1024*1024)
	ch <- aMap
	println(len(ch))
	b := <-ch
	println(b["a"])

}
