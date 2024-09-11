package query_test_test

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"go.uber.org/atomic"
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func TestContent3(t *testing.T) {
	// 打开 HTML 文件
	file, err := os.Open("test.html")
	if err != nil {
		return
	}
	defer file.Close()
	//
	var builder strings.Builder
	_, err = io.Copy(&builder, file)
	if err != nil {
		fmt.Printf("无法读取文件内容: %v", err)
		return
	}
	content := builder.String()
	// 使用 goquery 解析 HTML 文件
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	if err != nil {
		fmt.Printf("无法读取文件内容: %v", err)
		return
	}

	redMap := make(map[string]bool)
	doc.Find("p").Each(func(index int, element *goquery.Selection) {
		// 获取 p 标签的文本内容
		text := element.Text()
		html, _ := element.Html()
		if element.Text() == "" {
			return
		}
		if !strings.Contains(text, "保健") {
			return
		}

		// 为每一行添加红色样式
		redMap[html] = true
	})
	fmt.Printf("%+v", redMap)

	for k, _ := range redMap {
		content = strings.ReplaceAll(content, k, "<span style='color:red'>"+k+"</span>")
	}

	fmt.Println("=====")
	fmt.Println(content)
	return
}

func TestContent4(t *testing.T) {
	// 打开 HTML 文件
	file, err := os.Open("test2.html")
	if err != nil {
		return
	}
	defer file.Close()
	//
	var builder strings.Builder
	_, err = io.Copy(&builder, file)
	if err != nil {
		fmt.Printf("无法读取文件内容: %v", err)
		return
	}
	content := builder.String()
	// 使用 goquery 解析 HTML 文件
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	if err != nil {
		fmt.Printf("无法读取文件内容: %v", err)
		return
	}
	hitCount := atomic.NewInt64(0)
	wg := sync.WaitGroup{}
	doc.Find("p").Each(func(index int, element *goquery.Selection) {

		go func() {
			wg.Add(1)
			defer wg.Done()
			// 获取 p 标签的文本内容
			text := element.Text()
			if !strings.Contains(text, "西瓜土豆大鸭梨") {
				return
			}
			hitCount.Add(1)
			element.SetAttr("style", "color:red")
		}()

	})
	wg.Wait()
	// 生成新的 HTML 内容
	modifiedHTML, _ := doc.Html()

	// 创建新文件以保存修改后的 HTML
	newFile, err := os.Create("modified_test3.html")
	if err != nil {
		t.Errorf("无法创建新文件: %v", err)
		return
	}
	defer newFile.Close()

	// 将修改后的 HTML 写入新文件
	_, err = newFile.WriteString(modifiedHTML)
	if err != nil {
		t.Errorf("无法写入文件: %v", err)
		return
	}

	fmt.Println(hitCount.Load())
	return
}
