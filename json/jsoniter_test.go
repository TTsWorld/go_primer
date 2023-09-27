package json

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/bytedance/sonic"
	jsoniter "github.com/json-iterator/go"
)

var s = "{\"api_param\":\"{\\\"uid\\\":\\\"15205579\\\"}\",\"appsflyer-id\":\"1684734461278-4911811637343970480\",\"caller_addr\":\"27.97.64.62:40230\",\"channel\":\"Google Play\",\"country\":\"IN\",\"did\":\"0c0b9af0dd0141ec\",\"elapsed\":480,\"event_id\":\"/proto.user.UserService/GetUserBasic\",\"gender\":1,\"idfa\":\"50f359f0-1be7-489e-b862-96787d291479\",\"ip\":\"27.97.64.62\",\"lang\":\"en\",\"locale\":\"en_US\",\"mcc\":404,\"register_time\":\"2023-05-15 03:53:04.000\",\"result\":0,\"show-id\":20227886,\"sys_version\":\"android-11-V2037\",\"time\":\"2023-05-24 07:38:22.248\",\"tz\":5,\"uid\":20227886,\"vc\":1430095,\"vn\":\"1.0.0.2\"}"

var json2 = jsoniter.ConfigCompatibleWithStandardLibrary
var fastJson = jsoniter.ConfigFastest
var m = map[string]interface{}{}

func init() {
	json.Unmarshal([]byte(s), &m)
}

func TestUnmarshal(t *testing.T) {
	var m1 = map[string]interface{}{}
	var m2 = map[string]interface{}{}
	err := json.Unmarshal([]byte(s), &m1)
	fmt.Printf("%#v, err:%v \n", m1, err)
	jsoniter.UnmarshalFromString(m1["api_param"].(string), &m2)
	s, err := jsoniter.Marshal(m1)
	fmt.Printf("-------%s, err:%v \n", string(s), err)
}

func Testmarshal(t *testing.T) {
	err := json.Unmarshal([]byte(s), &m)
	fmt.Printf("%#v, err:%v \n", m, err)
	s, err := jsoniter.Marshal(m)
	fmt.Printf("-------%s, err:%v \n", string(s), err)
}

func BenchmarkJsoniterMarshal(b *testing.B) {

	for i := 0; i < b.N; i++ {
		marshalByJsoniter(m)
	}
}

func BenchmarkJsonMarshal(b *testing.B) {

	for i := 0; i < b.N; i++ {
		marshalByJson(m)
	}
}

func BenchmarkSonicarshal(b *testing.B) {

	for i := 0; i < b.N; i++ {
		marshalBySonic(m)
	}
}

//func BenchmarkMapGen(b *testing.B) {
//
//	for i := 0; i < b.N; i++ {
//		marshalByMapGen(m)
//	}
//}

func BenchmarkJsoniterFast(b *testing.B) {

	for i := 0; i < b.N; i++ {
		marshalByJsoniterFast(m)
	}
}

func marshalByJsoniter(a any) {
	jsoniter.Marshal(a)

}

func marshalByJson(a any) {
	json.Marshal(a)
}

func marshalBySonic(a any) {
	sonic.Marshal(&a)
}

//func marshalByMapGen(a map[string]interface{}) {
//	mapGen(a)
//}

func marshalByJsoniterFast(a map[string]interface{}) {
	fastJson.Marshal(&a)
}
