package diff

import (
	"strings"
	"testing"
)

func TestM2(t *testing.T) {
	origin, _ = ExtractText(C1)
	modify, _ = ExtractText(C2)
	a := DiffTextOnlie(origin, modify)
	print(a)
}

func TestExtract(t *testing.T) {
	origin, _ = ExtractText(C1)
	modify, _ = ExtractText(C2)
	println(origin)
	for i, s := range strings.Split(origin, "\n") {
		println(i, s)
	}
}
