package _interface

import (
	"bytes"
	"fmt"
	"testing"
)

//golang 接口转换的原理

func Test0301(t *testing.T) {
	var i int = 9

	var f float64
	f = float64(i)
	fmt.Printf("%T, %v\n", f, f)

	f = 10.8
	a := int(f)
	fmt.Printf("%T, %v\n", a, a)
}

type Test03Interface interface{}

func Test0302(t *testing.T) {
	var buf interface{} = bytes.Buffer{}
	fmt.Printf("%v \n", buf)        // {[] 0 0}
	fmt.Printf("%v \n", buf == nil) // false, 当前 buf 是一个 iface 的 interface，其 value ==nil ,  但是_type !=nil,  所以此时 buf!=nil.
}

// ================== error01和 error02 的区别主要是 New 返回值是否是指针
type error01 interface {
	Error01() string
}

func New(text string) error01 { return errorString{text: text} }

type errorString struct {
	text string
}

func (e errorString) Error01() string {
	return e.text
}

type error02 interface {
	Error02() string
}

func New2(text string) error02 { return &errorString{text: text} }

func (e *errorString) Error02() string {
	return e.text
}

func Test0303(t *testing.T) {
	w1 := New("ERR")
	w2 := New("ERR")
	fmt.Println(w1 == w2) // w1, w2相等是由于
	fmt.Println(w1.Error01())
	fmt.Println(w2.Error01())

	w3 := New2("ERR1")
	w4 := New2("ERR1")
	fmt.Println(w3 == w4) //结构体相等是结构体里所有的值相等
	fmt.Println(w3.Error02())
	fmt.Println(w4.Error02())

}
