package os

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestProcs(t *testing.T) {
	fmt.Printf("os:%v", filepath.Base(os.Args[0]))
}
