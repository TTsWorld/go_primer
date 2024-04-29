package other

import (
	"testing"
)

/*
给一个字符串 str ，包含多个单词， 单词之间以及字符串前后可能有多个#。要求：
	1	将字符串中的单词提取出来，按输入串的中出现的次序，做一个逆序排列， 单词之间用空格分割
	2	请不要使用库函数来分割字符串
例如， 输入：str = "##Please##show####me#the####code" 输出：str = "code the me show Please"
*/

func solution10(param string) string {
	if len(param) == 0 {
		return ""
	}

	ans := ""
	l := len(param) - 1
	r := l
	for l > 0 {
		if param[l] == '#' && param[l+1] != '#' {
			ans = ans + param[l+1:r+1] + " "
			r = l
		}
		if param[l] == '#' && r == l {
			r--
			l--
			continue
		}
		l--
	}

	return ans
}

func TestSplit(t *testing.T) {
	in := "##Please##show####me#the####code"
	println(solution10(in))

}

func TestString(t *testing.T) {
	var a chan int
	<-a
	//a := string(10000)
	//fmt.Print(a)
}
