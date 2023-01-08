package unsafe

import (
	"fmt"
	"testing"
	"unsafe"
)

//unsafe.Pointer
func TestName(t *testing.T) {

	f1 := float64(1.343535)
	f1p := (unsafe.Pointer(&f1))
	i1p := *(*uint64)(f1p)
	fmt.Printf("%#v", i1p)

}
