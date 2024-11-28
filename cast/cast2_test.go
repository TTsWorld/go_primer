package cast

import (
	"fmt"
	"github.com/spf13/cast"
	"testing"
)

func TestBool(t *testing.T) {
	b := string("0")
	b2 := string("1")
	s1 := cast.ToBool(b)
	s2 := cast.ToBool(b2)
	fmt.Printf("%+v\n", s1)
	fmt.Printf("%+v\n", s2)

}
