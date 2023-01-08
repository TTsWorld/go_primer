package os

import (
	"fmt"
	"os"
	"testing"
)

func TestOpenFile(t *testing.T) {
	f, err := os.OpenFile("/data/log/aaa/bbb", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Printf("%+v", err)
	}
	defer f.Close()
}
