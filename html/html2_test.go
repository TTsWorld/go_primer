package html

import (
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/anaskhan96/soup"
	"github.com/spf13/cast"

	"golang.org/x/net/html"
)

func TestSplitContent2(t *testing.T) {
	sections := SplitHTMLByWordCount(text, 2500)

	// 验证分段结果
	for i, section := range sections {
		fmt.Printf("Section %d length: %d \n", i+1, GetContentWordCount(section))
		if len(section) < 2000 || len(section) > 3000 {
			//t.Logf("Warning: Section %d has irregular length: %d", i+1, len(section))
		}
	}
}

func splitHTMLByWordCount(htmlContent string, wordCountLimit int) []string {
	var result []string
	var currentHTML string
	var wordCount int

	z := strings.NewReader(htmlContent)
	tokenizer := html.NewTokenizer(z)

	for {
		tokenType := tokenizer.Next()
		token := tokenizer.Token()
		tData := token.Data
		tType := token.Type
		lentAttribute := len(token.Attr)
		fmt.Printf("tokenType: %v, token: %v, tData: %v, tType: %v, lentAttribute: %v\n", tokenType, token, tData, tType, lentAttribute)

		switch tokenType {
		case html.TextToken:
			// 统计文本中的单词数（这里简单地按空格分割）
			wordCount += len([]rune(token.Data))

			// 如果当前 HTML 文本的字数超过限制，则将其添加到结果列表中，并重置计数器
			if wordCount >= wordCountLimit {
				result = append(result, currentHTML)
				currentHTML = ""
				wordCount = 0
			}

			// 将当前文本添加到当前 HTML 文本
			currentHTML += token.Data
		case html.StartTagToken, html.EndTagToken, html.SelfClosingTagToken:
			// 添加 HTML 标签
			currentHTML += token.String()
		case html.ErrorToken:
			if err := tokenizer.Err(); err != io.EOF {
				fmt.Println("error:", err)
			}
			return result
		}
	}

	// 将最后一个 HTML 文本添加到结果列表中
	if currentHTML != "" {
		result = append(result, currentHTML)
	}

	return result
}

func SplitHTMLByWordCount(htmlContent string, wordCountLimit int) []string {
	var result []string
	var currentHTML string
	var wordCount int

	z := strings.NewReader(htmlContent)
	tokenizer := html.NewTokenizer(z)
	var canAppend bool

	for {
		tokenType := tokenizer.Next()
		token := tokenizer.Token()

		switch tokenType {
		case html.TextToken:
			// 统计文本中的单词数（这里简单地按空格分割）
			content := strings.ReplaceAll(token.Data, string(rune(160)), "")
			content = strings.ReplaceAll(content, string(rune(32)), "")
			content = strings.ReplaceAll(content, "\n", "")
			content = strings.ReplaceAll(content, "\t", "")
			content = strings.ReplaceAll(content, "\r", "")
			wordCount += len([]rune(content))

			// 如果当前 HTML 文本的字数超过限制，则将其添加到结果列表中，并重置计数器
			if wordCount >= wordCountLimit {
				canAppend = true
			}

			// 将当前文本添加到当前 HTML 文本
			currentHTML += token.Data
		case html.EndTagToken:
			currentHTML += token.String()
			if token.Data != "p" {
				continue
			}
			if canAppend {
				result = append(result, currentHTML)
				currentHTML = ""
				wordCount = 0
				canAppend = false
			}
		case html.StartTagToken, html.SelfClosingTagToken:
			currentHTML += token.String()
		case html.ErrorToken:
			if err := tokenizer.Err(); err != io.EOF {
				fmt.Println("error:", err)
			} else {
				if currentHTML != "" {
					result[len(result)-1] = result[len(result)-1] + currentHTML
				}
			}
			return result
		default:
			continue
		}
	}
}

func GetContentWordCount(content string) int64 {
	doc := soup.HTMLParse(content)
	fullText := doc.FullText()
	text := strings.ReplaceAll(fullText, string(rune(160)), "")
	text = strings.ReplaceAll(text, string(rune(32)), "")
	text = strings.ReplaceAll(text, "\n", "")
	text = strings.ReplaceAll(text, "\t", "")
	text = strings.ReplaceAll(text, "\r", "")
	return cast.ToInt64(len([]rune(text)))
}
