package diff

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

// ExtractText 提取 html 的文本
func ExtractText(content string) (string, error) {

	content = strings.Replace(content, "\n", "", -1)
	buffer := bytes.NewBufferString(content)
	doc, err := goquery.NewDocumentFromReader(buffer)
	if err != nil {
		return "", err
	}
	t := doc.Text()
	t = strings.Replace(t, "\\n", "\n", -1)
	return t, nil
}
