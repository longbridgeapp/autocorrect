package autospace

import (
	regexp "github.com/dlclark/regexp2"
)

type strategory struct {
	one   string
	other string

	opt option
}

type option struct {
	noSpace         bool
	reverseValidate bool
}

var strategies []*strategory

func addStrategory(one, other string, opt option) {
	strategies = append(strategies, &strategory{
		one:   one,
		other: other,
		opt:   opt,
	})
}

func (s *strategory) format(in string) (out string) {
	out = in
	if s.opt.noSpace {
		out = s.removeSpace(out)
		return
	}

	out = s.addSpace(out)
	return
}

func (s *strategory) addSpace(in string) (out string) {
	out = in

	re := regexp.MustCompile("("+s.one+`)(`+s.other+")", 0)

	out, _ = re.Replace(out, "$1 $2", -1, -1)

	if !s.opt.reverseValidate {
		return
	}

	re = regexp.MustCompile("("+s.other+`)(`+s.one+")", 0)
	out, _ = re.Replace(out, "$1 $2", -1, -1)

	return
}

func (s *strategory) removeSpace(in string) (out string) {
	out = in

	re := regexp.MustCompile("("+s.one+`)\s+(`+s.other+")", 0)

	out, _ = re.Replace(out, "$1 $2", -1, -1)

	if !s.opt.reverseValidate {
		return
	}

	re = regexp.MustCompile("("+s.other+`)\s+(`+s.one+")", 0)
	out, _ = re.Replace(out, "$1 $2", -1, -1)

	return
}

func init() {
	// EnglishLetter
	addStrategory(`(?![年月日号])\p{Han}`, `[a-zA-Z]`, option{reverseValidate: true})
	// Number
	addStrategory(`(?![年月日号])\p{Han}`, `[0-9]`, option{reverseValidate: true})
	// SpecialSymbol
	addStrategory(`(?![年月日号])\p{Han}`, `[+$@#\/]`, option{reverseValidate: true})
	addStrategory(`(?![年月日号])\p{Han}`, `[\[\(‘“]`, option{})
	addStrategory(`[’”\]\)!]`, `(?![年月日号])\p{Han}`, option{})
	addStrategory(`[”\]\)!]`, `[a-zA-Z0-9]+`, option{})
	// Date
	addStrategory(`[\d[年月日]]{2,}`, `(?![年月日号])\p{Han}`, option{reverseValidate: true})
	// FullwidthPunctuation
	addStrategory(`(?![年月日号])[\w\p{Han}]`, `[，。！？：；”’]`, option{noSpace: true})
	addStrategory(`[‘“]`, `(?![年月日号])[\w\p{Han}]`, option{noSpace: true})
}

// Format auto format string to add spaces between Chinese and English words.
func Format(in string) (out string) {
	out = in

	for _, s := range strategies {
		out = s.format(out)
	}

	return
}
