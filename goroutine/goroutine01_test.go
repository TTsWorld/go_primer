package goroutine

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

// ==========0101, from geektime course
// 空的 select 将永远阻塞
// case 启示:如果你的 goroutine 在从另一个 goroutine 获得结果之前无法取得进展，那么通常情况下，你自己去做这项工作比委托它（go func(){}）更简单
func Test0101(t *testing.T) {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "hello world")
	})

	go func() {
		if err := http.ListenAndServe(":8088", nil); err != nil {
			log.Fatal(err)
		}
	}()
	select {} //空的 select 将永远阻塞

}

// ==========0102, from geektime course
// 让调用者决定是否并发
// case 启示:如果你的 goroutine 在从另一个 goroutine 获得结果之前无法取得进展，那么通常情况下，你自己去做这项工作比委托它（go func(){}）更简单
func Test0102(t *testing.T) {
	ListDirectoryv1 := func(dir string) []string { return nil }
	ListDirectoryv1 = ListDirectoryv1

	ListDirectoryv2 := func(dir string) chan string { return nil }
	ListDirectoryv2 = ListDirectoryv2

	ListDirectoryv3 := func(dir string, fn func(string)) {}
	ListDirectoryv3 = ListDirectoryv3
}

// ==========0103, from geektime course
// 让调用者决定是否并发
// case 启示:如果你的 goroutine 在从另一个 goroutine 获得结果之前无法取得进展，那么通常情况下，你自己去做这项工作比委托它（go func(){}）更简单
func Test0103(t *testing.T) {

}
