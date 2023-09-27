package pie

import (
	"fmt"
	"github.com/elliotchance/pie/v2"
	"github.com/spf13/cast"
	"testing"
)

func TestPieJoin(t *testing.T) {
	var a any
	a = []int{1, 2, 3}

	fmt.Println(pie.Join(cast.ToIntSlice(a), ","))

}
