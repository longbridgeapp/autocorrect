package autocorrect

import "regexp"

var (
	removeSpaceRe = regexp.MustCompile(`(` + spaceRe + `+)?(` + cjk + `)(` + spaceRe + `+)?`)
)

// UnformatOption addition unformat options
type UnformatOption interface {
	Unformat(text string) string
}

// Unformat to remove all spaces
func Unformat(text string, options ...UnformatOption) string {
	text = removeSpaceRe.ReplaceAllString(text, "$2")

	for _, opt := range options {
		text = opt.Unformat(text)
	}

	return text
}
