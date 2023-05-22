package _interface

import (
	"fmt"
	"testing"
	"unsafe"
)

//go 大坑 :interface{} 与 nil 的比较

func interfaceIsNil(x interface{}) {
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("non-empty interface")
}

// 在 go 中，非空接口由 iface 表示，空接口由 eface 表示，都是结构体，所以 interface 是否为空取决于结构体内的
// 一个接口包括动态类型和动态值, 如果一个接口的动态类型和动态值都为空，则这个接口为空的。
func Test0101(t *testing.T) {
	var x interface{} = nil
	var y *int = nil
	fmt.Printf("%+v \n", x == nil) //true
	fmt.Printf("%+v \n", y == nil) //true

	interfaceIsNil(x) // empty interface, x 的动态类型和值都为空，所以通过该函数判断 Nil 时为空成立
	interfaceIsNil(y) // non-empty interface ,y的动态值为空，但是动态类型不为空, 则不为 Nil
}

func Test0102(t *testing.T) {
	var a interface{} = nil         // 动态类型 = nil, data = nil
	var b interface{} = (*int)(nil) // 动态类型 包含 *int 类型信息, data = nil

	fmt.Println(a == nil) //true
	fmt.Println(b == nil) //false
}

type iface struct {
	tab  *itab
	data unsafe.Pointer
}
type itab struct {
	inter uintptr
	_type uintptr
	link  uintptr
	hash  uint32
	_     [4]byte
	fun   [1]uintptr
}

func Test0103(t *testing.T) {
	var a interface{} = nil         // 动态类型 = nil, data = nil
	var b interface{} = (*int)(nil) // 动态类型 包含 *int 类型信息, data = nil

	ap := (*iface)(unsafe.Pointer(&a))
	bp := (*iface)(unsafe.Pointer(&b))
	//打印接口的 hash 值
	fmt.Printf("interface hash:%+v \n", ap.tab.hash) //true
	fmt.Printf("interface hash:%+v \n", bp.tab.hash) //true
}
