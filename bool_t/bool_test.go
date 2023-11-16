package bool_t

import (
	"fmt"
	"testing"

	"github.com/spf13/cast"
)

func TestBool01(t *testing.T) {
	fmt.Println(cast.ToBool(int32(1)))
	fmt.Println(cast.ToBool(int32(0)))
	fmt.Println(cast.ToBool(int32(2)))
}
