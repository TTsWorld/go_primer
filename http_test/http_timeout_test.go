package htt_test

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

//使用 HTTP 读取请求中的各种基本信息

func TestHttpTimeout(t *testing.T) {
	//写个函数，统计该函数执行时间
	tm := time.Now()

	client := http.Client{Timeout: time.Second * 5}
	resp, err := client.Get("http://localhost:8090/sleep")
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	fmt.Printf("Response: %s\n", body)
	fmt.Printf("cost: %s\n", time.Now().Sub(tm))
}

func TestHttpTimeout2(t *testing.T) {
	//写个函数，统计该函数执行时间
	tm := time.Now()

	ctx, fn := context.WithTimeout(context.Background(), 10*time.Second)
	defer fn()

	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8090/sleep", nil)
	if err != nil {
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	fmt.Printf("Response: %s\n", body)
	fmt.Printf("cost: %s\n", time.Now().Sub(tm))
}
