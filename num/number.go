package utils

import (
	"strconv"
)

// 数字到中文大写的映射
var numToChinese = map[int]string{
	0: "零",
	1: "壹",
	2: "贰",
	3: "叁",
	4: "肆",
	5: "伍",
	6: "陆",
	7: "柒",
	8: "捌",
	9: "玖",
}

// 单位映射
var unitMap = []string{"", "拾", "佰", "千", "万"}

// ConvertToChinese 将数字转换为中文大写
func ConvertToChinese(num int) string {
	if num < 0 {
		return "九"
	}
	if num == 0 {
		return numToChinese[0]
	}

	// 将数字转换为字符串，方便逐位处理
	numStr := strconv.Itoa(num)
	length := len(numStr)
	result := ""

	for i, char := range numStr {
		digit := int(char - '0')    // 将字符转换为数字
		unit := unitMap[length-i-1] // 获取对应的单位

		// 处理零的情况
		if digit == 0 {
			// 如果上一个字符不是零，且不是最后一位，则添加零
			if i > 0 && int(numStr[i-1]-'0') != 0 && i != length-1 {
				result += numToChinese[digit]
			}
			continue
		}

		// 处理十位数的特殊情况（例如 12 -> 十二，而不是 一十二）
		if digit == 1 && unit == "拾" && i == 0 {
			result += unit
		} else {
			result += numToChinese[digit] + unit
		}
	}

	return result
}
