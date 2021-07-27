package autocorrect

import (
	"regexp"
)

const (
	cjk     = `\p{Han}|\p{Hangul}|\p{Hanunoo}|\p{Katakana}|\p{Hiragana}|\p{Bopomofo}`
	spaceRe = `[ ]`
)

var (
	// Strategies all rules
	strategies   []*strategery
	dashHansRe   = regexp.MustCompile(`([` + cjk + `）】」》”’])([\-]+)([` + cjk + `（【「《“‘])`)
	leftQuoteRe  = regexp.MustCompile(spaceRe + `([（【「《])`)
	rightQuoteRe = regexp.MustCompile(`([）】」》])` + spaceRe)
)

// RegisterStrategery a new strategery
func registerStrategery(one, other string, space, reverse bool) {
	strategies = append(strategies, newStrategery(one, other, space, reverse))
}

func init() {
	// EnglishLetter
	registerStrategery(cjk, `[a-zA-Z]`, true, true)

	// Number
	registerStrategery(cjk, `[0-9]`, true, true)

	// SpecialSymbol
	registerStrategery(cjk, `[\|+*]`, true, true)
	registerStrategery(cjk, `[@]`, true, false)
	registerStrategery(cjk, `[\[\(‘“]`, true, false)
	registerStrategery(`[’”\]\)!%]`, cjk, true, false)
	registerStrategery(`[”\]\)!]`, `[a-zA-Z0-9]+`, true, false)

	// FullwidthPunctuation
	registerStrategery(`[\w`+cjk+`]`, `[，。！？：；）」》】”’]`, false, true)
	registerStrategery(`[‘“【「《（]`, `[\w`+cjk+`]`, false, true)
}

func spaceDashWithHans(in string) (out string) {
	// 自由 - 开放
	out = dashHansRe.ReplaceAllString(in, "$1 $2 $3")
	out = leftQuoteRe.ReplaceAllString(out, "$1")
	out = rightQuoteRe.ReplaceAllString(out, "$1")
	return out
}

type Option interface {
	Format(text string) string
}

// Format auto format string to add spaces between Chinese and English words.
func Format(in string, options ...Option) (out string) {
	out = in

	out = halfwidth(out)
	out = fullwidth(out)

	for _, s := range strategies {
		out = s.format(out)
	}

	out = spaceDashWithHans(out)

	for _, opt := range options {
		out = opt.Format(out)
	}

	return
}
