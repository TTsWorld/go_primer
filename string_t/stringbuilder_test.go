package string_t

import (
	"fmt"
	"strings"
	"testing"
)

func TestStringBuilder(t *testing.T) {
	var b1 strings.Builder
	b1.WriteString("ABC")
	fmt.Printf("len:%v, cap:%v \n", b1.Len(), b1.Cap())
	//使用 Grow 预定义 slice 的容量, 可以避免扩容而新建 slice
	b1.Grow(10)
	fmt.Printf("len:%v, cap:%v \n", b1.Len(), b1.Cap())
	b1.WriteString("ABC")
	fmt.Printf("len:%v, cap:%v \n", b1.Len(), b1.Cap())
	b1.WriteString("ABC")
	fmt.Printf("len:%v, cap:%v \n", b1.Len(), b1.Cap())
	fmt.Printf("%v \n", b1)
	fmt.Printf("%s \n", b1.String())
}
