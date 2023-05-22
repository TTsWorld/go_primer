package htt_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

//使用 HTTP 读取请求中的各种基本信息

func queryParams(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	fmt.Fprintf(w, "query is %v\n", values)
}

func wholeUrl(w http.ResponseWriter, r *http.Request) {
	data, _ := json.Marshal(r.URL)
	fmt.Fprintf(w, string(data))
}

func header(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "header is %v\n", r.Header)
}

func form(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "before parse form %v\n", r.Form)
	err := r.ParseForm() //使用 form 前需要先 ParseForm
	if err != nil {
		fmt.Fprintf(w, "parse form error %v\n", r.Form)
	}
	fmt.Fprintf(w, "before parse form %v\n", r.Form)
}

func TestHttp02(t *testing.T) {
	http.HandleFunc("/queryParams", queryParams)
	http.HandleFunc("/wholeUrl", wholeUrl)
	http.HandleFunc("/header", header)
	http.HandleFunc("/form", form)

	http.ListenAndServe(":8090", nil)
}
