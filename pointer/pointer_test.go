package pointer

import (
	"encoding/json"
	"fmt"
	"testing"
)

type P struct {
	A int
}

func TestP1(t *testing.T) {
	var p *P
	ss, err := json.Marshal(p)
	fmt.Println(ss, err)
	p2 := new(P)
	p = p2
	p.A = 10
	fmt.Println(p.A)
}
