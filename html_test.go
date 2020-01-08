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
	for i := 0; i < b.N; i++ {
		// about 1.4ms/op
		FormatHTML(readFile("example.txt"))
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
