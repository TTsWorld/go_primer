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
	htmlContent := `
    <p>嫖娼被抓有啥后果<span class=" fw-cl "><span>？</span></span> - 知乎</p>
    <p><a class="Versatile-Link" href="https://www.zhihu.com/question/306119126"><u>https://www.zhihu.com/question/306119126</u></a></p>
    <p>男人去嫖娼被抓了有多丢人<span class=" fw-cl "><span>？</span></span> - 知乎</p>
    <p><a class="Versatile-Link" href="https://www.zhihu.com/question/652885153"><u>https://www.zhihu.com/question/652885153</u></a></p>
    <p>人大本硕<span class=" fw-cl "><span>，</span></span>天之骄子<span class=" fw-cl "><span>，</span></span>孩子刚出生两周<span class=" fw-cl "><span>，</span></span>出门接老家的亲人<span class=" fw-cl "><span>。</span></span></p>
    <p>啪<span class=" fw-cl "><span>，</span></span>人没了<span class=" fw-cl "><span>。</span></span></p>
    <p>妻子又惊又怒<span class=" fw-cl "><span>：</span></span>我这么大个老公呢<span class=" fw-cl "><span>？</span></span></p>
    <p>直到警察接了电话<span class=" fw-cl "><span>，</span></span>家人才知道<span class=" fw-cl "><span>：</span></span></p>
    <p>这个人中龙凤<span class=" fw-cl "><span>，</span></span>居然在去接亲人的路上<span class=" fw-cl "><span>，</span></span>做<span class="fw-op  "><span>「</span></span>大保健<span class=" fw-cl fw--collapsed"><span>」</span></span><span class=" fw-cl "><span>，</span></span>遇上便衣扫黄<span class=" fw-cl "><span>。</span></span></p>
    <p>更离谱的是<span class=" fw-cl "><span>，</span></span>他拒捕跳车<span class=" fw-cl "><span>，</span></span>把自己给<span class="fw-op  "><span>「</span></span>摔死了<span class=" fw-cl fw--collapsed"><span>」</span></span><span class=" fw-cl "><span>。</span></span></p>
    <p>然而<span class=" fw-cl "><span>，</span></span>看到尸体的那一刻<span class=" fw-cl "><span>，</span></span>家人起了疑心<span class=" fw-cl "><span>。</span></span></p>
    <p>不对<span class=" fw-cl "><span>，</span></span>他的死<span class=" fw-cl "><span>，</span></span>绝对没有那么简单<span class=" fw-cl "><span>！</span></span></p>
    `
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
