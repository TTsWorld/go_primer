package byte

import (
	"bytes"
	"fmt"
	"testing"
)

const (
	separator = "::"
	trimText  = "event.log|ta|"
)

var separatorB = []byte(separator)

var b = []byte("event.log|ta|{\"user_id\": \"111\", \"distinct_id\": \"123\"} ")

func Test01(t *testing.T) {
	b1 := []byte("event.log|ta|{\"user_id\": \"abc\", \"distinct_id\": \"123\"} ::event.log|ta|{\"user_id\": \"def\", \"distinct_id\": \"456\"}")
	preProcess(b1)
}

func preProcess(b []byte) []byte {
	if b == nil || len(b) == 0 {
		return nil
	}

	rawText := bytes.Split(b, []byte(separator))
	rawCnt := len(rawText)
	result := make([]byte, 0, rawCnt)
	for _, r := range rawText {

		a := bytes.TrimLeft(r, trimText)
		fmt.Printf("%v \n", string(a))
		//b := map[string]string{}
		//json.Unmarshal(a, &b)
		//fmt.Printf("%v \n", b)
		result = append(result, a...)
		result = append(result, separatorB...)
	}
	result = bytes.TrimSuffix(result, separatorB)
	fmt.Printf("%v \n", string(result))
	return result
}

func Test22(t *testing.T) {
	var a any
	fmt.Printf("%v", a == "")
}
