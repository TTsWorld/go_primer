package query_test_test

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"regexp"
	"strings"
	"testing"
)

func TestContent2(t *testing.T) {
	// 这里是你的 HTML 内容
	htmlContent := ``
	// 按行分割文本
	re := regexp.MustCompile(`[\p{P}\n\r]+`)

	// 使用正则表达式拆分字符串
	// 使用 goquery 解析 HTML 内容
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		fmt.Println("Error loading HTML: ", err)
		return
	}

	// 查找所有的 <p> 标签，并打印其中的文本内容
	doc.Find("p").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		fmt.Println(text)
		replaced := re.ReplaceAll([]byte(text), []byte(""))
		fmt.Println(string(replaced))
		fmt.Println("=======")
	})
}
