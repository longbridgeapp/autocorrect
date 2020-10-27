package autocorrect

import (
	"regexp"
)

var (
	fullwidthMaps = map[string]string{
		",": "，",
		".": "。",
		";": "；",
		":": "：",
		"!": "！",
		"?": "？",
		"~": "～",
		// "(": "（",
		// ")": "）",
	}

	spcialPunctuations = `[.:]`
	normalPunctuations = `[,;\!\?~]`

	punctuationWithLeftCJKRe        = regexp.MustCompile(normalPunctuations + `[` + cjk + `]+`)
	punctuationWithRightCJKRe       = regexp.MustCompile(`[` + cjk + `]+` + normalPunctuations)
	punctuationWithSpeicalCJKRe     = regexp.MustCompile(`[` + cjk + `]+` + spcialPunctuations + `[` + cjk + `]+`)
	punctuationWithSpeicalLastCJKRe = regexp.MustCompile(`[` + cjk + `]+` + spcialPunctuations + "$")
	punctuationsRe                  = regexp.MustCompile(`(` + spcialPunctuations + `|` + normalPunctuations + `)`)
)

// fullwidth correct punctuations near the CJK chars
func fullwidth(text string) (out string) {
	out = text

	out = punctuationWithLeftCJKRe.ReplaceAllStringFunc(out, fullwidthReplacePart)
	out = punctuationWithRightCJKRe.ReplaceAllStringFunc(out, fullwidthReplacePart)
	out = punctuationWithSpeicalCJKRe.ReplaceAllStringFunc(out, fullwidthReplacePart)
	out = punctuationWithSpeicalLastCJKRe.ReplaceAllStringFunc(out, fullwidthReplacePart)

	return
}

func fullwidthReplacePart(part string) string {
	part = punctuationsRe.ReplaceAllStringFunc(part, func(str string) string {
		str = fullwidthMaps[str]
		return str
	})

	return part

}
