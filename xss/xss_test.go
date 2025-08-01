package xss

import (
	"fmt"
	"strings"
	"testing"

	"github.com/microcosm-cc/bluemonday"

	"golang.org/x/net/html"
)

var o1 = ""
var o3 = `<p data-block-key="1uarp">现</p>\n<p data-block-key="5cd3a">分割线她在个月前就用了同样的方式侮辱造谣了妹妹和其他几个女孩<span class=" fw-cl "><span>。</span></span>&nbsp;</p>\n<p data-block-key="776cl"><br></p>\n<p data-block-key="7a8de"><br></p>\n<p data-block-key="4qcom"><br></p>\n<p data-block-key="6c82a"><br></p>\n<p data-block-key="63hvq"><br></p>\n<p data-block-key="dtcl3"><br></p>\n<p data-block-key="5om4r"><br></p>\n<p data-block-key="fpdv2"><br></p>`
var o4 = `<p data-block-key="9kuj1">现&nbsp;</p>`
var o5 = `<p data-block-key="9kuj1">花里>胡「」哨[的]测试文件</p>
<p data-block-key="7e7kl"><br></p>`

func TestXss(t *testing.T) {
	p := bluemonday.UGCPolicy() // 这个策略会给 某些标签添加 rel="nofollow" 属性
	p.AllowElements("p", "html", "head", "body", "h1", "h2", "li", "ul", "ol", "blockquote", "code", "em", "figure", "hr", "pre", "span", "sup", "sub", "figcaption", "table", "tr", "td", "th", "tbody", "strong", "video", "audio", "br")
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

func TestAnalyzeHtml(t *testing.T) {
	// 使用 bluemonday 的 parser 来解析 HTML
	reader := strings.NewReader(o1)
	doc, err := html.Parse(reader)
	if err != nil {
		t.Fatal(err)
	}

	// 使用 map 来存储唯一的标签和属性
	tags := make(map[string]struct{})
	attrs := make(map[string]struct{})

	// 递归遍历 HTML 节点
	var traverse func(*html.Node)
	traverse = func(n *html.Node) {
		if n.Type == html.ElementNode {
			tags[n.Data] = struct{}{}
			for _, attr := range n.Attr {
				attrs[attr.Key] = struct{}{}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}
	}
	traverse(doc)

	// 打印结果
	fmt.Println("唯一标签:")
	for tag := range tags {
		fmt.Printf("- %s\n", tag)
	}

	fmt.Println("\n唯一属性:")
	for attr := range attrs {
		fmt.Printf("- %s\n", attr)
	}
}

func TestName(t *testing.T) {
	m := make(map[string]interface{})
	m["a"] = "1872680804360003585"
	var b, _ = m["a"].(float64)
	fmt.Println(b)
}

func TestEqu(t *testing.T) {
	fmt.Println(o0701a == o0701b)

}
func TestZtest(t *testing.T) {
	content := o0302
	// parse the web page
	//node, _ := ztext.NewZOutputText(content, ztext.AllowGif(true), ztext.StrictBMP(false))
	node, _ := ztext.NewZOutputText(content)
	str, _ := ztext.Render(node)
	fmt.Println(str == content)
}
