package xss

import (
	"fmt"
	"strings"
	"testing"

	"github.com/microcosm-cc/bluemonday"
)

func TestXss0(t *testing.T) {
	p := bluemonday.UGCPolicy()
	p.AllowElements("p", "html", "head", "body", "h1", "h2", "li", "ul", "br", "ol", "blockquote", "code", "em", "figure", "hr", "pre", "span", "sup", "sub", "figcaption", "table", "tr", "td", "th", "tbody", "strong", "video", "audio")
	p.AllowAttrs("class", "style", "id").Globally()
	p.AllowAttrs("href", "name", "target").OnElements("a")
	p.AllowAttrs("src", "srcset", "alt", "title", "width", "height", "loading").OnElements("img")
	p.AllowDataAttributes()
	p.AllowStandardAttributes()
	p.AllowStyles()

	o71 = strings.ReplaceAll(o030301, "&nbsp;", " ")
	//o2 := o1
	//o2 = o2

	html := p.Sanitize(o71)
	html = strings.ReplaceAll(html, "&gt;", ">")
	html = strings.ReplaceAll(html, "&lt;", "<")
	fmt.Println(html)

	fmt.Println(o1 == html)
}
func TestOnline(t *testing.T) {
	p := bluemonday.NewPolicy()
	p.AllowElements("p", "html", "head", "body", "h1", "h2", "li", "ul", "br", "ol", "blockquote", "code", "em", "figure", "hr", "pre", "span", "sup", "sub", "figcaption", "table", "tr", "td", "th", "tbody", "strong", "video", "audio")
	p.AllowAttrs("class", "style", "id").Globally()
	p.AllowAttrs("href", "name", "target").OnElements("a")
	p.AllowAttrs("src", "srcset", "alt", "title", "width", "height", "loading").OnElements("img")
	p.AllowDataAttributes()
	p.AllowStandardAttributes()

	content := o0701a

	afterHtml := p.Sanitize(content)
	fmt.Println(content == afterHtml)
	//println(content)
	println("=======")
	println(afterHtml)
}
