package gob

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func TestD(t *testing.T) {
	str := ""
	n, _ := fmt.Scan(&str)
	if n == 0 {
		return
	}

}

func solution(arrLen int, arr []int, n int) int {
	m := make(map[int]struct{})
	if arrLen != len(arr) {
		return -1
	}

	newArr := make([]int, 0)
	for _, i2 := range arr {
		if _, ok := m[i2]; ok {
			return -1
		}
		m[i2] = struct{}{}
		newArr = append(newArr, i2)
	}

	sort.Ints(newArr)
	sum := 0
	newArrLen := len(newArr)

	if newArrLen < 2*n {
		return -1
	}

	i, j := 0, newArrLen-1

	for i < newArrLen {
		if i < n {
			sum += newArr[i]
			sum += newArr[j]
			i++
			j--
		} else {
			break
		}
	}

	return sum

}

func TestE(t *testing.T) {
	//arr := []int{3, 2, 3, 4, 2}
	//arr := []int{95, 88, 83, 64, 100}
	//println(solution(5, arr, 2))
	arrLen := 0
	_, err := fmt.Scanf("%d", &arrLen)
	if err != nil {
		println(err.Error())
		return
	}
	arrStr := ""
	fmt.Scanf("%s", &arrStr)
	n := 0
	fmt.Scanf("%n", &n)

	arrSlice := strings.Split(arrStr, " ")
	arr := []int{}
	for _, s := range arrSlice {
		tmp, _ := strconv.Atoi(s)
		arr = append(arr, tmp)
	}
	fmt.Println(arrLen)
	fmt.Printf("%v \n", arr)
	fmt.Println(n)

	println(solution(arrLen, arr, n))

}
