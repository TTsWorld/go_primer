package html

import (
	"bytes"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/anaskhan96/soup"
)

// 存储段落信息的结构
type Paragraph struct {
	HTML    string
	TextLen int
}

// 存储分段结果的结构
type Section struct {
	Content string
	TextLen int
}

func TestGoquery(t *testing.T) {
	buffer := bytes.NewBufferString(text)
	doc, err := goquery.NewDocumentFromReader(buffer)
	if err != nil {
		return
	}
	println(doc)
}

func TestSoup(t *testing.T) {
	sp := soup.HTMLParse(text).FullText()
	println(sp)
}

func TestSplitContent(t *testing.T) {
	sections := splitContent(text)

	// 验证分段结果
	for i, section := range sections {
		t.Logf("Section %d length: %d", i+1, section.TextLen)
		if section.TextLen < 2000 || section.TextLen > 3000 {
			t.Logf("Warning: Section %d has irregular length: %d", i+1, section.TextLen)
		}
	}
}

func splitContent(content string) []Section {
	// 解析HTML内容
	paragraphs := extractParagraphs(content)

	// 按照字数规则分组
	var sections []Section
	currentSection := Section{}

	for _, p := range paragraphs {
		// 如果当前段落加入后会超过2500字，且当前section已经有内容
		if currentSection.TextLen > 0 && currentSection.TextLen+p.TextLen > 2500 {
			sections = append(sections, currentSection)
			currentSection = Section{}
		}

		currentSection.Content += p.HTML
		currentSection.TextLen += p.TextLen
	}

	// 添加最后一个section
	if currentSection.TextLen > 0 {
		sections = append(sections, currentSection)
	}

	// 处理最后一个section如果少于2000字的情况
	if len(sections) > 1 && sections[len(sections)-1].TextLen < 2000 {
		lastSection := sections[len(sections)-1]
		sections = sections[:len(sections)-1]
		sections[len(sections)-1].Content += lastSection.Content
		sections[len(sections)-1].TextLen += lastSection.TextLen
	}

	return sections
}

func extractParagraphs(content string) []Paragraph {
	buffer := bytes.NewBufferString(content)
	doc, err := goquery.NewDocumentFromReader(buffer)
	if err != nil {
		return nil
	}

	var paragraphs []Paragraph

	// 选择所有HTML元素
	doc.Find("*").Each(func(i int, s *goquery.Selection) {

		sText := s.Text()
		println(sText)
		// 跳过script、style等无意义标签
		if tag := goquery.NodeName(s); tag == "script" || tag == "style" {
			return
		}

		// 只有当前元素直接包含文本，才将其作为段落
		// 这样可以避免重复计算嵌套元素的文本
		if containsDirectText(s) {
			html, _ := s.Html()
			text := strings.TrimSpace(s.Text())
			if len(text) > 0 {
				paragraphs = append(paragraphs, Paragraph{
					HTML:    html,
					TextLen: len([]rune(text)),
				})
			}
		}
	})

	return paragraphs
}

// 检查元素是否直接包含文本节点
func containsDirectText(s *goquery.Selection) bool {
	hasText := false
	// 遍历所有直接子节点
	s.Contents().Each(func(i int, sel *goquery.Selection) {
		if sel.Get(0).Type == 3 { // 3 是文本节点的类型
			text := strings.TrimSpace(sel.Text())
			if len(text) > 0 {
				hasText = true
			}
		}
	})
	return hasText
}
