package json

import (
	"fmt"
	"github.com/bytedance/sonic"
	"testing"
)

func TestSonic(t *testing.T) {
	v, err := sonic.Marshal(m)
	fmt.Printf("result:%s, err:%v", v, err)
}
