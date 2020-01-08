package autospace

import (
	"regexp"
)

type strategory struct {
	one   string
	other string

	opt option
}

type option struct {
	space   bool
	reverse bool
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
	if s.opt.space {
		out = s.addSpace(out)
	} else {
		out = s.removeSpace(out)
	}

	return
}

func (s *strategory) addSpace(in string) (out string) {
	out = in

	re := regexp.MustCompile("(" + s.one + `)(` + s.other + ")")

	out = re.ReplaceAllString(out, "$1 $2")

	if s.opt.reverse {
		re = regexp.MustCompile("(" + s.other + `)(` + s.one + ")")
		out = re.ReplaceAllString(out, "$1 $2")
	}

	return
}

func (s *strategory) removeSpace(in string) (out string) {
	out = in

	re := regexp.MustCompile("(" + s.one + `)\s+(` + s.other + ")")

	out = re.ReplaceAllString(out, "$1 $2")

	if s.opt.reverse {
		re = regexp.MustCompile("(" + s.other + `)\s+(` + s.one + ")")
		out = re.ReplaceAllString(out, "$1 $2")
	}

	return
}

func init() {
	// EnglishLetter
	addStrategory(`\p{Han}`, `[a-zA-Z]`, option{space: true, reverse: true})

	// Number
	addStrategory(`\p{Han}`, `[0-9]`, option{space: true, reverse: true})

	// SpecialSymbol
	addStrategory(`\p{Han}`, `[\|+$@#]`, option{space: true, reverse: true})
	addStrategory(`\p{Han}`, `[\[\(‘“]`, option{space: true})
	addStrategory(`[’”\]\)!]`, `\p{Han}`, option{space: true})
	addStrategory(`[”\]\)!]`, `[a-zA-Z0-9]+`, option{space: true})

	// FullwidthPunctuation
	addStrategory(`[\w\p{Han}]`, `[，。！？：；」》】”’]`, option{reverse: true})
	addStrategory(`[‘“【「《]`, `[\w\p{Han}]`, option{reverse: true})

	// Fix full date
}

// Format auto format string to add spaces between Chinese and English words.
func Format(in string) (out string) {
	out = in

	for _, s := range strategies {
		out = s.format(out)
	}

	return
}
