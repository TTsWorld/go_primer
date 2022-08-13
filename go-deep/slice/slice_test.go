package slice

import (
	"fmt"
	"testing"
)

// =========================slice传参
func Test01(t *testing.T) {
	a := make([]int, 10, 20)
	b := a[5:]
	println(len(b), cap(b)) // 5 15

	//b [0]  = a[5] ,切片后的 slice 和原 slice 指向同一内存
	b[0] = 1
	println(a[5]) // 1

}

func doAppend(a []int) {
	_ = append(a, 0)
}

//函数传参会影响入参切片，因为底层数组是同一个，
func Test02(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	doAppend(a[0:2])
	fmt.Println(a) // [1 2 0 4 5]

	//如何避免切片后的操作影响原 slice
	b := []int{1, 2, 3, 4, 5}
	doAppend(b[0:2:2]) //指定切片时的cap 为2，此时函数内 append 会出发内存分配，则 b 会指向新的内存地址，不会影响 a
	fmt.Println(b)     // [1 2 0 4 5]
}

// =========================slice append
func Test03(t *testing.T) {
	slice1 := make([]int, 0, 4)
	slice1 = append(slice1, 1, 2, 3)

	slice2 := append(slice1, 4)
	slice2[0] = 10

	fmt.Println(slice1) //[10,2,3]  ,因为 slice1 的 len 为 3，所以只能打出 3 个值
	fmt.Println(slice2) //[10,2,3,4]

}
func Test04(t *testing.T) {
	slice1 := make([]int, 0, 4)
	slice1 = append(slice1, 1, 2, 3)

	slice2 := append(slice1, 4, 5) //append 完后发生了扩容，导致 slice2 指向了自己的空间
	slice2[0] = 10

	fmt.Println(slice1) // [1,2,3]
	fmt.Println(slice2) // [10,2,3,4,5]
}

// =========================slice 切片
func Test05(t *testing.T) {
	slice1 := make([]int, 0, 4)
	slice1 = append(slice1, 1, 2, 3)

	slice2 := slice1[:len(slice1)-1]
	slice2[0] = 10

	fmt.Println(slice1)
	fmt.Println(slice2)
}

func Test06(t *testing.T) {
	slice1 := make([]int, 0, 4)
	slice1 = append(slice1, 1, 2, 3)

	slice2 := slice1[:len(slice1)-1]
	slice2 = append(slice2, 11, 12, 13, 14, 15) //append 完后发生了扩容，导致 slice2 指向了自己的空间
	slice2[0] = 10

	fmt.Println(slice1)
	fmt.Println(slice2)

}

//golang的slice操作默认都是浅拷贝。触发发生扩容才会让两个slice指向不同的数组。在实际业务中，很多场景是需要深拷贝的，这个时候可以使用copy函数

// =========================slice for-range

//错误使用 循环迭代器的变量
func TestSliceForRange(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6}
	b := []*int{}
	for i, v := range a {
		fmt.Printf("%d -- %d\n", i, v)
		b = append(b, &i)
	}
	fmt.Println(*b[0], *b[1], *b[2])
	fmt.Println(b)
}

func TestSliceForRange02(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6}
	b := []*int{}
	for i, v := range a {
		fmt.Printf("%d -- %d\n", i, v)
		b[i] = &v
	}
	fmt.Println(*b[0], *b[1], *b[2])
	fmt.Println(b)

}

// =========================slice clear

func dump(letters []string) {
	fmt.Println("letters = ", letters)
	fmt.Println(cap(letters))
	fmt.Println(len(letters))
	for i := range letters {
		fmt.Println(i, letters[i])
	}
}
func TestSliceClear(t *testing.T) {
	letters := []string{"a", "b", "c", "d"}
	dump(letters)
	// clear the slice
	letters = nil
	dump(letters)
	// add stuff back to it
	letters = append(letters, "e")
	dump(letters)
}
