package err

import (
	"fmt"
	"golang.org/x/xerrors"
	"testing"
)

func BenchmarkXerrors(b *testing.B) {
	for n := 0; n < b.N; n++ {
		xerrors.Errorf("format %s", "test")
	}

}

func BenchmarkErrorf(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fmt.Errorf("format %s", "test")
	}
}
