package other

import (
	"fmt"
	"testing"
)

/*
生成全排列 考虑去重
input : [1, 2, 3]
output:[
    [1, 2, 3],
    [1, 3, 2],
    [2, 1, 3],
...
]

input:
    [1, 1]
output:
    [[1, 1]]
*/

func TestName2(t *testing.T) {
	var arr [3]int
	brr := arr[1:2]
	fmt.Printf("%v\n", arr)
	fmt.Printf("%v\n", brr)
	brr = append(brr, 9)
	fmt.Printf("%v\n", brr)
	brr = append(brr, 9)
	brr[1] = 8
	fmt.Printf("%v\n", brr)
	fmt.Printf("%v\n", arr)
}

/*
给定两个字符串str1和str2，再给定三个整数ic，dc和rc，分别代表插入、删除和替换一个字符的代价，请输出将str1编辑成str2的最小代价。
给定一个非负整数数组nums，假定最开始处于下标为0的位置，数组里面的每个元素代表下一跳能够跳跃的最大长度，如果可以跳到数组最后一个位置，请你求出跳跃路径中所能获得的最多的积分。
1.如果能够跳到数组最后一个位置，才能计算所获得的积分，否则积分值为-1
2.如果无法跳跃(即数组长度为0)，也请返回-1
3.数据保证返回的结果不会超过整形范围，即不会超过
*/
