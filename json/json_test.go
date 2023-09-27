package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type User struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

func TestName(t *testing.T) {

	arr := []*User{}
	arr = append(arr, &User{Name: "aa", Age: "aa"})
	arr = append(arr, &User{Name: "bb", Age: "bb"})
	arr = append(arr, &User{Name: "cc", Age: "cc"})
	v, _ := json.Marshal(arr)
	fmt.Println(string(v))
}
