package autocorrect

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"testing"
)

var (
	htmlSpaceRe = regexp.MustCompile(`>[\s]+<`)
)

func assertHTMLEqual(t *testing.T, exptected, actual string) {
	if htmlSpaceRe.ReplaceAllString(exptected, "><") != htmlSpaceRe.ReplaceAllString(actual, "><") {
		t.Errorf("\nexptected:\n%s\nactual   :\n%s", exptected, actual)
	}
}

func readFile(filename string) (out string) {
	data, err := ioutil.ReadFile(fmt.Sprintf("./_fixtures/%s", filename))
	if err != nil {
		panic(err)
	}
	return string(data)
}

func BenchmarkFormatHTML(b *testing.B) {
	raw := readFile("example.txt")
	for i := 0; i < b.N; i++ {
		// about 1.1ms/op
		FormatHTML(raw)
	}
}
func TestFormatHTMLWithFixtuires(t *testing.T) {
	expected := readFile("example.expected.txt")
	out, err := FormatHTML(readFile("example.txt"))
	if err != nil {
		t.Error(err)
	}
	assertHTMLEqual(t, expected, out)
}

func TestFormatHTMLWithEscapedHTML(t *testing.T) {
	html := `<p>据2019年12月27日，三生制药获JP Morgan Chase &amp; Co.每股均价9.582港元，增持270.3万股</p>`
	expected := `<p>据2019年12月27日，三生制药获 JP Morgan Chase &amp; Co.每股均价 9.582 港元，增持 270.3 万股</p>`
	out, err := FormatHTML(html)
	if err != nil {
		t.Error(err)
	}
	assertHTMLEqual(t, expected, out)

	html = `<p>据2019年12月27日，三生制药获JP Morgan Chase & Co.每股均价9.582港元，增持270.3万股</p>`
	expected = `<p>据2019年12月27日，三生制药获 JP Morgan Chase &amp; Co.每股均价 9.582 港元，增持 270.3 万股</p>`
	out, err = FormatHTML(html)
	if err != nil {
		t.Error(err)
	}
	assertHTMLEqual(t, expected, out)
}

func TestFormatHTML_HalfWidth(t *testing.T) {
	html := `<p>自动转换全角“字符、数字”：我们将在（１６：３２）出发去ＣＢＤ中心。</p>`

	out, err := FormatHTML(html)
	if err != nil {
		t.Error(err)
	}

	assertEqual(t, "<p>自动转换全角 “字符、数字”：我们将在（16:32）出发去 CBD 中心。</p>", out)
}
