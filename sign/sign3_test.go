package sign

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"google.golang.org/appengine/log"
	"sort"
	"strings"
)

// @param body 请求体
// @param appSecret string 密钥
// @return string
func sign2(body []byte, appSecret string) string {
	// 1. 对请求体中的所有参数，根据参数名作升序排序。
	rawMsgs := make(map[string]json.RawMessage)
	err := json.Unmarshal(body, &rawMsgs)
	if err != nil {
		fmt.Printf("[CheckRequestSignValid] json parse raw fail, body = %s, err:%v \n", string(body), err)
		return err.Error()
	}

	fieldMap := make(map[string]string)
	fieldKeyList := make([]string, 0)
	for key, msg := range rawMsgs {
		vals, err := msg.MarshalJSON()
		if err != nil {
			fmt.Printf("[CheckRequestSignValid] msge marshal raw fail, key = %s, msg = %s, err: %v \n", key, string(msg), err)
			return err.Error()
		}
		fmt.Println(string(vals))
		fieldMap[key] = string(vals)
		fieldKeyList = append(fieldKeyList, key)
	}

	sort.Slice(fieldKeyList, func(i, j int) bool {
		return fieldKeyList[i] < fieldKeyList[j]
	})

	// 2，将排序好的参数名与参数值拼接在一起（格式：key1value1key2value2...keyNvalueN），得到字符串 s1。
	s1 := ""
	for _, key := range fieldKeyList {
		val := strings.Trim(fieldMap[key], `"`)
		s1 += key + val
	}
	fmt.Printf("[CheckRequestSignValid] s1: %s \n", s1)

	// 3. 将app_secret拼接到从步骤2获得的字符串s1的结尾。得到字符串：x1。
	x1 := s1 + appSecret
	log.Infof(ctx, "[CheckRequestSignValid] x1: %s", x1)

	// 4.对步骤3获得的字符串x1作HmacSha256(x1,[]byte(app_secret))签名。得到字节序列：b1。
	h := hmac.New(sha256.New, []byte(appSecret))
	h.Write([]byte(x1))

	// 5.对步骤4获得对字节序列b1作base64编码，base64(b1),得到字符串：s2。
	s2 := base64.StdEncoding.EncodeToString(h.Sum(nil))
	fmt.Printf("[AuthenticationCtrlImpl:CheckRequestSignValid] s2: %s \n", s2)

	// 6.将步骤5获得对字符串s2全部转为大写，得到最终签名
	finalSign := strings.ToUpper(s2)
	fmt.Printf("[CheckRequestSignValid] xsign = %s,  body: %s \n", finalSign, string(body))
	return ""
}
