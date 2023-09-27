package hex

import "testing"

// 在json-iterator中，valueTypes是用来表示JSON数据的值类型的枚举类型。通过使用这个枚举类型，我们可以很容易地判断JSON数据的值类型，并据此执行不同的操作

func TestHex01(t *testing.T) {
	hexDigits := make([]byte, 256)
	for i := 0; i < len(hexDigits); i++ {
		hexDigits[i] = 255
	}
	for i := '0'; i <= '9'; i++ {
		hexDigits[i] = byte(i - '0')
	}
	for i := 'a'; i <= 'f'; i++ {
		hexDigits[i] = byte((i - 'a') + 10)
	}
	for i := 'A'; i <= 'F'; i++ {
		hexDigits[i] = byte((i - 'A') + 10)
	}
}
