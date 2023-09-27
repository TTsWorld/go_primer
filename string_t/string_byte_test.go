package string_t

import (
	"reflect"
	"testing"
	"unsafe"
)

func BenchmarkStringByteCommon(b *testing.B) {
	x := []byte("hello worldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworld")
	for i := 0; i < b.N; i++ {
		byte2StringCommon(x)
	}

}

func BenchmarkStringByteReflect(b *testing.B) {
	x := []byte("hello worldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworld")
	for i := 0; i < b.N; i++ {
		byte2StringReflect(x)
	}

}
func BenchmarkStringByteReflectSimple(b *testing.B) {
	x := []byte("hello worldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworld")
	for i := 0; i < b.N; i++ {
		byte2StringReflectSimple(x)
	}

}
func byte2StringCommon(b []byte) string {
	return string(b)
}

func byte2StringReflect(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{bh.Data, bh.Len}
	return *(*string)(unsafe.Pointer(&sh))
}

func byte2StringReflectSimple(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
