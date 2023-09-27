package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Resource struct {
	ID   int `json:"id"`
	Type int `json:"type"`
}

type Data struct {
	Resources []Resource `json:"resources"`
}

func TestJson02(t *testing.T) {
	jsonData := `
         {"1":{"resources":[{"id":1,"type":1},{"id":2,"type":1},{"id":3,"type":1}]},"2":{"resources":[{"id":1,"type":1},{"id":2,"type":1},{"id":3,"type":1}]}}
        `

	// Create a map to hold the parsed data
	dataMap := make(map[string]Data)

	// Unmarshal the JSON data into the map
	err := json.Unmarshal([]byte(jsonData), &dataMap)
	if err != nil {
		fmt.Println("Error while unmarshaling JSON:", err)
		return
	}

	// Now you can access the data using the keys "1", "2", "3", "4"
	fmt.Println("Data for key '1':", dataMap["1"])
	fmt.Println("Data for key '2':", dataMap["2"])
	// ... and so on
}

func Test03(t *testing.T) {
	res := make([]Resource, 0)
	res = append(res, Resource{ID: 1, Type: 1})
	res = append(res, Resource{ID: 2, Type: 1})
	res = append(res, Resource{ID: 3, Type: 1})
	dataMap := make(map[string]Data)
	dataMap["1"] = Data{Resources: res}
	dataMap["2"] = Data{Resources: res}
	v, _ := json.Marshal(dataMap)
	fmt.Println(string(v))
}
