package autocorrect

import (
	"bytes"
	"io"
	"regexp"
	"strings"

	"github.com/pkg/errors"
	"github.com/tdewolff/parse/v2"
	"github.com/tdewolff/parse/v2/html"
	// "golang.org/x/net/html"
)

var (
	ignoreTagsRe = regexp.MustCompile("(?mi)<(pre|script|style|textarea)")
)

// FormatHTML format HTML content
func FormatHTML(body string, options ...Option) (out string, err error) {
	return processHTML(body, func(text string) string {
		return Format(text, options...)
	})
}

// UnformatHTML cleanup spaces for HTML
func UnformatHTML(body string, options ...UnformatOption) (out string, err error) {
	return processHTML(body, func(text string) string {
		return Unformat(text, options...)
	})
}

func processHTML(body string, fn func(plainText string) string) (out string, err error) {
	w := &bytes.Buffer{}
	i := parse.NewInput(strings.NewReader(body))
	lex := html.NewLexer(i)

	out = body

	ignoreTag := false

	for {
		t, data := lex.Next()

		switch t {
		case html.ErrorToken:
			if lex.Err() == io.EOF {
				return w.String(), nil
			}

			err = errors.Errorf("Error on line %d, %v", i.Offset(), lex.Err())
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
