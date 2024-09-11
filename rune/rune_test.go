package rune

import (
	"fmt"
	"testing"
)

func TestRune(t *testing.T) {
	// 定义一个字符串
	str := "Hello, 世界"

	// 将字符串转换为 rune 切片
	runes := []rune(str)
	fmt.Println("字符串转换为 rune 切片:", runes)

	// 遍历字符串中的每个 rune
	for i, r := range str {
		fmt.Printf("字符 %c 在位置 %d\n", r, i)
	}

	// 将 rune 转换回字符串
	newStr := string(runes)
	fmt.Println("rune 切片转换回字符串:", newStr)
}
