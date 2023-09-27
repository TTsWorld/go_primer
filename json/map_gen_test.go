package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

var s1 = "{\"api_param\":\"{\\\"uid\\\":\\\"15205579\\\"}\",\"appsflyer-id\":\"1684734461278-4911811637343970480\",\"caller_addr\":\"27.97.64.62:40230\",\"channel\":\"Google Play\",\"country\":\"IN\",\"did\":\"0c0b9af0dd0141ec\",\"elapsed\":480,\"event_id\":\"/proto.user.UserService/GetUserBasic\",\"gender\":1,\"idfa\":\"50f359f0-1be7-489e-b862-96787d291479\",\"ip\":\"27.97.64.62\",\"lang\":\"en\",\"locale\":\"en_US\",\"mcc\":404,\"register_time\":\"2023-05-15 03:53:04.000\",\"result\":0,\"show-id\":20227886,\"sys_version\":\"android-11-V2037\",\"time\":\"2023-05-24 07:38:22.248\",\"tz\":5,\"uid\":20227886,\"vc\":1430095,\"vn\":\"1.0.0.2\"}"
var m1 = map[string]interface{}{}

func init() {
	err := json.Unmarshal([]byte(s1), &m1)
	fmt.Printf("%#v, err:%v \n", m1, err)
}

func TestMapGen(t *testing.T) {
	//s, _ := mapGen(m1)
	//fmt.Printf("%s\n", s)

}
