package autocorrect

import (
	"regexp"
)

const (
	cjkRe   = `\p{Han}|\p{Hangul}|\p{Hanunoo}|\p{Katakana}|\p{Hiragana}|\p{Bopomofo}`
	spaceRe = `[ ]`
)

var (
	// Strategies all rules
	strategies   []*strategery
	fullDateRe   = regexp.MustCompile(spaceRe + `{0,}\d+` + spaceRe + `{0,}年` + spaceRe + `{0,}\d+` + spaceRe + `{0,}月` + spaceRe + `{0,}\d+` + spaceRe + `{0,}[日号]` + spaceRe + `{0,}`)
	dashHansRe   = regexp.MustCompile(`([` + cjkRe + `）】」》”’])([\-]+)([` + cjkRe + `（【「《“‘])`)
	leftQuoteRe  = regexp.MustCompile(spaceRe + `([（【「《])`)
	rightQuoteRe = regexp.MustCompile(`([）】」》])` + spaceRe)
)

// RegisterStrategery a new strategery
func registerStrategery(one, other string, space, reverse bool) {
	strategies = append(strategies, newStrategery(one, other, space, reverse))
}

func init() {
	// EnglishLetter
	registerStrategery(cjkRe, `[a-zA-Z]`, true, true)

	// Number
	registerStrategery(cjkRe, `[0-9]`, true, true)

	// SpecialSymbol
	registerStrategery(cjkRe, `[\|+*]`, true, true)
	registerStrategery(cjkRe, `[@]`, true, false)
	registerStrategery(cjkRe, `[\[\(‘“]`, true, false)
	registerStrategery(`[’”\]\)!%]`, cjkRe, true, false)
	registerStrategery(`[”\]\)!]`, `[a-zA-Z0-9]+`, true, false)

	// FullwidthPunctuation
	registerStrategery(`[\w`+cjkRe+`]`, `[，。！？：；）」》】”’]`, false, true)
	registerStrategery(`[‘“【「《（]`, `[\w`+cjkRe+`]`, false, true)
}

// removeFullDateSpacing
// 发布 2013 年 3 月 10 号公布 -> 发布2013年3月10号公布
func removeFullDateSpacing(in string) (out string) {
	spaceRe := regexp.MustCompile(spaceRe + "+")
	// Fix fulldate
	return fullDateRe.ReplaceAllStringFunc(in, func(part string) string {
		return spaceRe.ReplaceAllString(part, "")
	})
}

func spaceDashWithHans(in string) (out string) {
	// 自由-开放
	out = dashHansRe.ReplaceAllString(in, "$1 $2 $3")
	out = leftQuoteRe.ReplaceAllString(out, "$1")
	out = rightQuoteRe.ReplaceAllString(out, "$1")
	return out
}

// Format auto format string to add spaces between Chinese and English words.
func Format(in string) (out string) {
	out = in

	out = halfWidth(out)

	for _, s := range strategies {
		out = s.format(out)
	}

	out = removeFullDateSpacing(out)
	out = spaceDashWithHans(out)

	return
}
