package autocorrect

import "regexp"

var (
	removeSpaceRe = regexp.MustCompile(`(` + spaceRe + `+)?(` + cjkRe + `)(` + spaceRe + `+)?`)
)

// Unformat to remove all spaces
func Unformat(text string) string {
	text = removeSpaceRe.ReplaceAllString(text, "$2")
	return text
}
