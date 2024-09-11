package _go

import (
	"testing"
)

func TestSafeGo(t *testing.T) {
	SafeGo(func() {
		panic("panic")
	})
}
