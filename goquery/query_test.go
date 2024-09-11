package query_test_test

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestGoquery(t *testing.T) {
	// 打开 HTML 文件
	file, err := os.Open("test.html")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 使用 goquery 解析 HTML 文件
	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		log.Fatal(err)
	}

	// 遍历所有元素节点
	doc.Find("*").Each(func(index int, element *goquery.Selection) {
		// 打印元素标签
		fmt.Printf("Index:%d, Element: %s\n", index, goquery.NodeName(element))
		cnt := 0
		for i := 0; i < len(element.Text()); i++ {
			if element.Text()[i] == '\n' {
				cnt += 1
			}
		}
		if cnt > 0 {
			fmt.Printf("cnt: %d\n", cnt)
		}
		//统计 element.Text 中的换行符

		// 打印元素的属性
		element.Each(func(i int, s *goquery.Selection) {
			for _, attr := range s.Nodes[0].Attr {
				fmt.Printf("  Attribute: %s=%s\n", attr.Key, attr.Val)
			}
		})
	})
}

func TestModify(t *testing.T) {
	// 打开 HTML 文件
	file, err := os.Open("test.html")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	//
	var builder strings.Builder
	_, err = io.Copy(&builder, file)
	if err != nil {
		log.Fatalf("无法读取文件内容: %v", err)
	}
	content := builder.String()

	// 使用 goquery 解析 HTML 文件
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	if err != nil {
		log.Fatal(err)
	}
	redMap := make(map[string]bool)

	// 遍历所有元素节点
	doc.Find("p").Each(func(index int, s *goquery.Selection) {
		// 获取 p 标签的文本内容
		text := s.Text()
		if s.Text() == "" {
			return
		}

		// 按行分割文本
		lines := strings.Split(text, "\n")

		// 为每一行添加红色样式
		for _, line := range lines {
			trimmedLine := strings.TrimSpace(line)
			if strings.Contains(trimmedLine, "保健") {
				//fmt.Printf("line:%v\n", line)
				redMap[line] = true
			}
		}
		if len(lines) == 0 {
			//fmt.Printf("no:%+v", s.Text())
		}
		//fmt.Printf("no2:%+v", s.Text())
	})
	fmt.Println("======")
	fmt.Printf("%v", redMap)
	for k, _ := range redMap {
		content = strings.ReplaceAll(content, k, "bbbbbb")
	}
	fmt.Printf(content)
}

func TestModify2(t *testing.T) {
	// 打开 HTML 文件
	file, err := os.Open("test.html")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	//
	var builder strings.Builder
	_, err = io.Copy(&builder, file)
	if err != nil {
		log.Fatalf("无法读取文件内容: %v", err)
	}
	content := builder.String()

	// 使用 goquery 解析 HTML 文件
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	if err != nil {
		log.Fatal(err)
	}
	redMap := make(map[string]bool)

	// 遍历所有元素节点
	doc.Find("p").Each(func(index int, s *goquery.Selection) {
		// 获取 p 标签的文本内容
		text := s.Text()
		if s.Text() == "" {
			return
		}

		fmt.Println("111======")
		fmt.Println(text)
		fmt.Println("222======")

		// 按行分割文本
		re := regexp.MustCompile(`[\p{P}\n\r]+`)

		// 使用正则表达式拆分字符串
		parts := re.Split(text, -1)

		// 为每一行添加红色样式
		for _, line := range parts {
			trimmedLine := strings.TrimSpace(line)
			if strings.Contains(trimmedLine, "保健") {
				//fmt.Printf("line:%v\n", line)
				redMap[line] = true
			}
		}
		if len(parts) == 0 {
			//fmt.Printf("no:%+v", s.Text())
		}
		//fmt.Printf("no2:%+v", s.Text())
	})
	//fmt.Println("======")
	//fmt.Printf("%v", redMap)
	for k, _ := range redMap {
		content = strings.ReplaceAll(content, k, "bbbbbb")
	}
	//fmt.Printf(content)
}

func TestModifyOK(t *testing.T) {
	// 打开 HTML 文件
	file, err := os.Open("test.html")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	//
	var builder strings.Builder
	_, err = io.Copy(&builder, file)
	if err != nil {
		log.Fatalf("无法读取文件内容: %v", err)
	}
	content := builder.String()

	// 使用 goquery 解析 HTML 文件
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	if err != nil {
		log.Fatal(err)
	}
	redMap := make(map[string]bool)

	// 遍历所有元素节点
	doc.Find("p").Each(func(index int, s *goquery.Selection) {
		// 获取 p 标签的文本内容
		text := s.Text()
		if s.Text() == "" {
			return
		}

		// 按行分割文本
		re := regexp.MustCompile(`[\p{P}\n\r]+`)

		// 使用正则表达式拆分字符串
		parts := re.Split(text, -1)

		// 为每一行添加红色样式
		for _, line := range parts {
			trimmedLine := strings.TrimSpace(line)
			if strings.Contains(trimmedLine, "保健") {
				//fmt.Printf("line:%v\n", line)
				redMap[line] = true
			}
		}
		if len(parts) == 0 {
			//fmt.Printf("no:%+v", s.Text())
		}
		//fmt.Printf("no2:%+v", s.Text())
	})
	fmt.Println("======")
	fmt.Printf("%v", redMap)
	for k, _ := range redMap {
		content = strings.ReplaceAll(content, k, "<span style='color:red'>"+k+"</span>")

	}
	fmt.Printf(content)
}

// printElementHierarchy 用于递归打印元素及其子元素的树状结构。
func printElementHierarchy(indent int, element *goquery.Selection) {
	// 计算缩进字符串
	indentStr := ""
	for i := 0; i < indent; i++ {
		indentStr += "  "
	}

	// 打印元素标签
	fmt.Printf("%s- Index: %d, Element: %s\n", indentStr, element.Index(), goquery.NodeName(element))

	// 打印元素的属性
	element.Each(func(i int, s *goquery.Selection) {
		for _, attr := range s.Nodes[0].Attr {
			fmt.Printf("%s  Attribute: %s=%s\n", indentStr, attr.Key, attr.Val)
		}
	})

	// 递归打印子元素
	element.Children().Each(func(i int, child *goquery.Selection) {
		printElementHierarchy(indent+1, child)
	})
}

func TestGoquery2(t *testing.T) {
	// 打开 HTML 文件
	file, err := os.Open("test.html")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 使用 goquery 解析 HTML 文件
	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		log.Fatal(err)
	}

	// 遍历所有元素节点
	doc.Find("*").Each(func(index int, element *goquery.Selection) {
		printElementHierarchy(0, element)
	})
}
