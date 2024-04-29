package slice

import (
	"fmt"
	"testing"
)

func TestArr2Slice01(t *testing.T) {

	//声明，数组会被初始化为 0 值
	var a2 [3]int
	fmt.Printf("addr:%p, val:%+v\n", &a2, a2) //addr:0x140000aa048, val:[0 0 0]

	//声明初始化-不赋值
	s0 := a2[:]
	fmt.Printf("len:%d, cap:%d, addr:%p, val:%+v\n", len(s0), cap(s0), &s0[0], s0[0]) //len:3, cap:3, addr:0x14000018150 (这个地址跟上面数组的地址是同一个地址), val:0

	s0 = append(s0, 1)
	fmt.Printf("len:%d, cap:%d, addr:%p, val:%+v\n", len(s0), cap(s0), &s0[0], s0[0]) //len:4, cap:6, addr:0x140000162a0 (slice 扩容，重新分配内存，底层arr内存地址发生变化), val:0

	//声明并初始化
	a1 := [3]int{1, 2, 3}
	fmt.Printf("%+v\n", a1) //[1 2 3]

	//数组比较
	a11 := [3]int{1, 2, 3}
	a22 := [3]int{1, 2, 3}
	fmt.Printf("%+v\n", a11 == a22) //true

	//数组比较,不同长度的数组，不是同一个类型，不能进行比较
	a111 := [3]int{1, 2, 3}
	a222 := [4]int{1, 2, 3}
	a111 = a111
	a222 = a222
	//fmt.Printf("%+v\n", a111 == a222) //无法编译，数据类型不同

}

// 测试多个 slice 底层引用同一个数组，是否会互相影响
func TestArr2Sclie02(t *testing.T) {
	//不初始化，数组也可以直接使用
	var arr [5]int
	arr[0] = 1
	fmt.Printf("arr == addr:%p, val:%+v\n", &arr[0], arr) //arr == addr:0x1400011a060 (数组的地址就是数组收元素的地址), val:[1 0 0 0 0]

	println("111111111 case 修改底层数组，slice 是否会发生变化")
	s1 := arr[:3]
	fmt.Printf("sclie-s1 == len:%d, cap:%d, addr:%p, val:%+v\n", len(s1), cap(s1), &s1[0], s1) //sclie-s1 == len:3, cap:5, addr:0x1400011a060(与数组的地址相同), val:[2 0 0]
	arr[0] = 3
	fmt.Printf("arr == addr:%p, val:%+v\n", &arr[0], arr)                                      //arr == addr:0x1400011a060 (数组的地址就是数组收元素的地址), val:[2 0 0 0 0]
	fmt.Printf("sclie-s1 == len:%d, cap:%d, addr:%p, val:%+v\n", len(s1), cap(s1), &s1[0], s1) //sclie-s1 == len:3, cap:5, addr:0x1400011a060(与数组的地址相同), val:[2 0 0]
	println("111111111 case 修改底层数组，slice 是否会发生变化")

	s1[0] = 2
	fmt.Printf("arr == addr:%p, val:%+v\n", &arr[0], arr)                                      //arr == addr:0x1400011a060 (数组的地址就是数组收元素的地址), val:[2 0 0 0 0]
	fmt.Printf("sclie-s1 == len:%d, cap:%d, addr:%p, val:%+v\n", len(s1), cap(s1), &s1[0], s1) //sclie-s1 == len:3, cap:5, addr:0x1400011a060(与数组的地址相同), val:[2 0 0]

	s2 := s1
	fmt.Printf("sclie-s2 == len:%d, cap:%d, addr:%p, val:%+v\n", len(s2), cap(s2), &s2[0], s2) //sclie-s2 == len:3, cap:5, addr:0x1400011a060(与数组的地址相同), val:[2 0 0]
	s2[0] = 4                                                                                  // ** 结论: 只要内存中的数组是同一个，那么任何指向这快内存的修改，都会互相影响 **
	fmt.Printf("arr == addr:%p, val:%+v\n", &arr[0], arr)                                      //arr == addr:0x1400011a060 (数组的地址就是数组收元素的地址), val:[4 0 0 0 0]
	fmt.Printf("sclie-s1 == len:%d, cap:%d, addr:%p, val:%+v\n", len(s1), cap(s1), &s1[0], s1) //sclie-s1 == len:3, cap:5, addr:0x1400011a060, val:[4 0 0]
	fmt.Printf("sclie-s2 == len:%d, cap:%d, addr:%p, val:%+v\n", len(s2), cap(s2), &s2[0], s2) //sclie-s2 == len:3, cap:5, addr:0x1400011a060, val:[4 0 0]

	println("aaaaaaaaaaaaaa")
	s2 = append(s2, 5, 5)                                                                      //未扩容, **结论: 只要未扩容，底层数组就不会发生改变**
	fmt.Printf("arr == addr:%p, val:%+v\n", &arr[0], arr)                                      //arr == addr:0x1400011a060, val:[4 0 0 5 5]
	fmt.Printf("sclie-s1 == len:%d, cap:%d, addr:%p, val:%+v\n", len(s1), cap(s1), &s1[0], s1) //sclie-s1 == len:3, cap:5, addr:0x1400011a060, val:[4 0 0]
	// fmt.Printf("sclie-s1[4] == len:%d, cap:%d, val:%+v\n", len(s1), cap(s1), s1[4])            //panic-case sclie-s1[4] panic **结论，如果访问的长度超出数组的 len，会 panic, 不论 cap 是多长 **
	fmt.Printf("sclie-s2 == len:%d, cap:%d, addr:%p, val:%+v\n", len(s2), cap(s2), &s2[0], s2) //sclie-s2 == len:5, cap:5, addr:0x1400011a060, val:[4 0 0 5 5]

	println("bbbbbbbbbbbbbbb")
	s2 = append(s2, 6)                                                                         //发生扩容 **结论:哪个数组发生扩容，就影响哪个slice，另外一个 slice 不受影响 **
	fmt.Printf("arr == addr:%p, val:%+v\n", &arr[0], arr)                                      //arr == addr:0x1400011a060, val:[4 0 0 5 5] 底层数组一直不变
	fmt.Printf("sclie-s1 == len:%d, cap:%d, addr:%p, val:%+v\n", len(s1), cap(s1), &s1[0], s1) //sclie-s1 == len:3, cap:5, addr:0x1400011a060, val:[4 0 0]
	fmt.Printf("sclie-s2 == len:%d, cap:%d, addr:%p, val:%+v\n", len(s2), cap(s2), &s2[0], s2) //sclie-s2 == len:6, cap:10, addr:0x1400014e140, val:[4 0 0 5 5 6]
}

// 不同切片使用同一个底层数组，互相覆盖的情况
func TestSclieConver(t *testing.T) {
	a := make([]int, 0, 5)
	fmt.Printf("a:%p\n", a)
	b := append(a, 4)
	fmt.Printf("sclie-s1 == len:%d, cap:%d, addr:%p, val:%+v\n", len(b), cap(b), &b[0], b) //sclie-s1 == len:3, cap:5, addr:0x1400011a060, val:[4 0 0]
	c := append(a, 5)
	fmt.Printf("sclie-b == len:%d, cap:%d, addr:%p, val:%+v\n", len(b), cap(b), &b[0], b) //sclie-s1 == len:3, cap:5, addr:0x1400011a060, val:[4 0 0]
	fmt.Printf("sclie-c == len:%d, cap:%d, addr:%p, val:%+v\n", len(c), cap(c), &c[0], c) //sclie-s1 == len:3, cap:5, addr:0x1400011a060, val:[4 0 0]
	fmt.Println(a, b, c)
}

// 2 个 slice 互相引用
func TestName7(t *testing.T) {
	//a := []int{1, 2, 3}
	a := make([]int, 0, 5)
	fmt.Printf("a:%p\n", a)
	b := append(a, 4)
	fmt.Printf("b:%p\n", b)
	c := append(a, 5)
	fmt.Printf("c:%p\n", c)
	fmt.Println(a, b, c)

	//扩容
	aa := []int{1, 2, 3}
	fmt.Printf("aa:%p\n", aa)
	bb := append(aa, 4)
	fmt.Printf("bb:%p\n", bb)
	cc := append(aa, 5)
	fmt.Printf("cc:%p\n", cc)
	fmt.Println(aa, bb, cc)
	//a := make([]int, 5)
}
