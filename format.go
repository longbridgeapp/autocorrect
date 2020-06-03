package autocorrect

import (
	"regexp"
)

const (
	cjkRe = `\p{Han}|\p{Hangul}|\p{Hanunoo}|\p{Katakana}|\p{Hiragana}|\p{Bopomofo}`
)

var (
	// Strategies all rules
	strategies   []*strategery
	fullDateRe   = regexp.MustCompile(`[\s]{0,}\d+[\s]{0,}年[\s]{0,}\d+[\s]{0,}月[\s]{0,}\d+[\s]{0,}[日号][\s]{0,}`)
	spaceRe      = regexp.MustCompile(`\s+`)
	dashHansRe   = regexp.MustCompile(`([` + cjkRe + `）】」》”’])([\-]+)([` + cjkRe + `（【「《“‘])`)
	leftQuoteRe  = regexp.MustCompile(`\s([（【「《])`)
	rightQuoteRe = regexp.MustCompile(`([）】」》])\s`)
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
	registerStrategery(cjkRe, `[\|+$@#*]`, true, true)
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

	for _, s := range strategies {
		out = s.format(out)
	}

	out = removeFullDateSpacing(out)
	out = spaceDashWithHans(out)

	return
}
