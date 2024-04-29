package other

import (
	"fmt"
	"testing"
)

/*
编写一个函数来查找字符串数组中的最长公共前缀0。
如果不存在公共前缀，返回空字符串 ""。
*/
func solution7(arr []string) string {
	if len(arr) == 0 {
		return ""
	}

	// arr := {"c", "ab", "a"}
	prefix := arr[0]
	for i := 1; i < len(arr); i++ {
		length := 0
		if len(prefix) < len(arr[i]) {
			length = len(prefix)
		} else {
			length = len(arr[i])
		}

		index := 0
		for j := 0; j < length; j++ {
			if prefix[j] == arr[i][j] {
				index++
			} else {
				break
			}
		}
		prefix = prefix[:index]
	}

	return prefix
}

func TestName7(t *testing.T) {
	//arr := []string{"abc", "ab", "a"}
	//arr := []string{"c", "ab", "ab"}
	//fmt.Printf("%v \n", solution7(arr))

	//a := []int{1, 2, 3}
	a := make([]int, 0, 5)
	fmt.Printf("a:%p\n", a)
	b := append(a, 4)
	fmt.Printf("b:%p\n", b)
	c := append(a, 5)
	fmt.Printf("c:%p\n", c)
	fmt.Println(a, b, c) //如果没有触发扩容，c 会直接覆盖 b 写在 a 内存中的值

	//扩容
	aa := []int{1, 2, 3}
	fmt.Printf("aa:%p, aa[0]:%p, len(aa):%d, cap:%d \n", aa, &aa[0], len(aa), cap(aa))
	//aa:0x140000b2018, aa[0]:0x140000b2018, len(aa):3, cap:3
	bb := append(aa, 4)
	fmt.Printf("bb:%p, bb[0]:%p, len(aa):%d, cap:%d\n", bb, &bb[0], len(bb), cap(bb))
	//bb:0x140000be090, bb[0]:0x140000be090, len(aa):4, cap:6cap:3
	cc := append(aa, 5)
	fmt.Printf("cc:%p, cc[0]:%p, len(aa):%d, cap:%d\n", cc, &cc[0], len(cc), cap(cc))
	//cc:0x140000be0c0, cc[0]:0x140000be0c0, len(aa):4, cap:6
	fmt.Println(aa, bb, cc) // [1 2 3] [1 2 3 4] [1 2 3 5]
	//上面每次都会触发扩容，所以，bb,cc都是在新的内存中 append 的，所以不会互相影响

}
