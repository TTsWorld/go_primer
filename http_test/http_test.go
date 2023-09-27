package htt_test

import (
	"fmt"
	"io"
	"net/http"
	"testing"
)

// A fundamental concept in `net/http` servers is
// *handlers*. A handler is an object implementing the
// `http.Handler` interface. A common way to write
// a handler is by using the `http.HandlerFunc` adapter
// on functions with the appropriate signature.
func hello(w http.ResponseWriter, req *http.Request) {

	// Functions serving as handlers take a
	// `http.ResponseWriter` and a `http.Request` as
	// arguments. The response writer is used to fill in the
	// HTTP response. Here our simple response is just
	// "hello\n".
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	// This handler does something a little more
	// sophisticated by reading all the HTTP request
	// headers and echoing them into the response body.
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

// http天坑，body 只能读一次，再次读不回报错但读不到
func readBodyOnce(w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(w, "read body failed, err:%v \n", err)
		return
	}
	fmt.Fprintf(w, "read the data: %v\n", string(body))

	//再次读 body，啥也读不到，但是也不会报错
	body, err = io.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(w, "read body failed 222, err:%v \n", err)
		return
	}
	fmt.Fprintf(w, "read the data 222: %v\n", string(body))
}

func readBodyOnceResolve(w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(w, "read body failed, err:%v \n", err)
		return
	}
	fmt.Fprintf(w, "read the data: %v\n", string(body))

	//再次读 body，啥也读不到，但是也不会报错
	body, err = io.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(w, "read body failed 222, err:%v \n", err)
		return
	}
	fmt.Fprintf(w, "read the data 222: %v\n", string(body))
}

// GetBody默认为 nil，中间件可以考虑注入这个方法
func getBody(w http.ResponseWriter, req *http.Request) {
	body, _ := req.GetBody()
	if body == nil {
		fmt.Fprintf(w, "getBody is Nil \n")
	} else {
		fmt.Fprintf(w, "getBody is not Nil \n")
	}
}

func TestHttp(t *testing.T) {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/read", readBodyOnce)
	http.HandleFunc("/getbody", getBody)

	http.ListenAndServe(":8090", nil)
}
