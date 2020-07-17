package autocorrect

import (
	"io"
	"strings"

	"github.com/pkg/errors"
	"github.com/tdewolff/parse/v2/html"
	// "golang.org/x/net/html"
)

// FormatHTML format HTML content
func FormatHTML(body string) (out string, err error) {
	lex := html.NewLexer(strings.NewReader(body))
	out = body
	for {
		tt, data := lex.Next()
		switch tt {
		case html.TextToken:
			raw := string(data)
			formated := Format(raw)
			out = strings.Replace(out, raw, formated, -1)
		case html.ErrorToken:
			if lex.Err() == io.EOF {
				return out, nil
			}

			err = errors.Errorf("Error on line %d, %v", lex.Offset(), lex.Err())
			return
		}
	}
}
