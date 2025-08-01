package json

import (
	"fmt"
	"github.com/goccy/go-json"
	"testing"
)

func TestMain(t *testing.T) {
	s := ""
	m := make(map[string]string)
	err := json.Unmarshal([]byte(s), &m)
	fmt.Printf("%v", err)
	fmt.Printf("%v", m)

}
