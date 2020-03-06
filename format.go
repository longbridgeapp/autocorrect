package autocorrect

import (
	"regexp"
)

var (
	// Strategies all rules
	strategies   []*strategery
	fullDateRe   = regexp.MustCompile(`[\s]{0,}\d+[\s]{0,}年[\s]{0,}\d+[\s]{0,}月[\s]{0,}\d+[\s]{0,}[日号][\s]{0,}`)
	spaceRe      = regexp.MustCompile(`\s+`)
	dashHansRe   = regexp.MustCompile(`([\p{Han}）】」》”’])([\-]+)([\p{Han}（【「《“‘])`)
	leftQuoteRe  = regexp.MustCompile(`\s([（【「《])`)
	rightQuoteRe = regexp.MustCompile(`([）】」》])\s`)
)

// RegisterStrategery a new strategery
func registerStrategery(one, other string, opt option) {
	strategies = append(strategies, &strategery{
		one:   one,
		other: other,
		opt:   opt,
	})
}

func init() {
	// EnglishLetter
	registerStrategery(`\p{Han}`, `[a-zA-Z]`, option{space: true, reverse: true})

	// Number
	registerStrategery(`\p{Han}`, `[0-9]`, option{space: true, reverse: true})

	// SpecialSymbol
	registerStrategery(`\p{Han}`, `[\|+$@#*]`, option{space: true, reverse: true})
	registerStrategery(`\p{Han}`, `[\[\(‘“]`, option{space: true})
	registerStrategery(`[’”\]\)!%]`, `\p{Han}`, option{space: true})
	registerStrategery(`[”\]\)!]`, `[a-zA-Z0-9]+`, option{space: true})

	// FullwidthPunctuation
	registerStrategery(`[\w\p{Han}]`, `[，。！？：；）」》】”’]`, option{reverse: true})
	registerStrategery(`[‘“【「《（]`, `[\w\p{Han}]`, option{reverse: true})
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
