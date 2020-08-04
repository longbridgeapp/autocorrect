package autocorrect

import (
	"bytes"
	"io"
	"regexp"
	"strings"

	"github.com/pkg/errors"
	"github.com/tdewolff/parse/v2/html"
	// "golang.org/x/net/html"
)

var (
	ignoreTagsRe = regexp.MustCompile("(?mi)<(pre|script|style|textarea)")
)

// FormatHTML format HTML content
func FormatHTML(body string) (out string, err error) {
	return processHTML(body, Format)
}

// UnformatHTML cleanup spaces for HTML
func UnformatHTML(body string) (out string, err error) {
	return processHTML(body, Unformat)
}

func processHTML(body string, fn func(plainText string) string) (out string, err error) {
	w := &bytes.Buffer{}
	lex := html.NewLexer(strings.NewReader(body))
	defer lex.Restore()
	out = body

	ignoreTag := false

	for {
		t, data := lex.Next()

		switch t {
		case html.ErrorToken:
			if lex.Err() == io.EOF {
				return w.String(), nil
			}

			err = errors.Errorf("Error on line %d, %v", lex.Offset(), lex.Err())
			return
		case html.TextToken:
			if ignoreTag {
				if _, err := w.Write(data); err != nil {
					return out, err
				}

				ignoreTag = false
				continue
			}

			formated := fn(string(data))
			if _, err := w.Write([]byte(formated)); err != nil {
				return out, err
			}
		case html.StartTagToken:
			if ignoreTagsRe.Match(data) {
				ignoreTag = true
			}

			if _, err := w.Write(data); err != nil {
				return out, err
			}

		default:
			if _, err := w.Write(data); err != nil {
				return out, err
			}
		}
	}
}
