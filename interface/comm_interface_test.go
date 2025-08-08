package _interface

import (
	"fmt"
	"testing"

	"github.com/daozhonglee/go-util/json"
)

type CommMsg struct {
	Action string                 `json:"action"` // 请求 Request_Similarity 响应 Response_Similarity
	TaskID string                 `json:"task_id"`
	Code   int                    `json:"code"` //0 成功
	Data   interface{}            `json:"data"`
	Extra  map[string]interface{} `json:"extra"`
}

type SimilarityCommDataInterface interface {
	Unmarshal(data []byte) error
}

type DataImpl struct {
	Content   string `json:"content"`
	Threshold int32  `json:"threshold"`
}

type DataImpl2 struct {
	Subs []DataImpl2Sub `json:"subs"`
}
type DataImpl2Sub struct {
	Key   string `json:"content"`
	Value int    `json:"threshold"`
}

func Test01(t *testing.T) {
	msg := CommMsg{
		Action: "Request_Similarity",
		TaskID: "1234567890",
		Data:   &DataImpl2{Subs: []DataImpl2Sub{{Key: "sub1", Value: 1}, {Key: "sub2", Value: 2}}},
	}
	s := json.MarshalFailSafe(msg)
	fmt.Println(s)

	comm := &CommMsg{}
	err := json.Unmarshal([]byte(s), comm)
	if err != nil {
		fmt.Println(err)
	}

	impl := &DataImpl2{}
	err = json.Unmarshal([]byte(json.MarshalFailSafe(comm.Data)), &impl)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(json.MarshalFailSafe(impl))
}
