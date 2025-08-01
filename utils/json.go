package utils

import "encoding/json"

func MarshalString(i interface{}) string {
	s, err := json.Marshal(i)
	if err != nil {
		return ""
	}
	return string(s)
}
