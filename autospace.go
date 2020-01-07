package autospace

import (
	"fmt"
	"regexp"
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

	re := regexp.MustCompile("(" + s.one + `)(` + s.other + ")")
	fmt.Println(re)
	out = re.ReplaceAllString(out, "$1AAA$2")

	if !s.opt.reverseValidate {
		return
	}

	re = regexp.MustCompile("(" + s.other + `)(` + s.one + ")")
	out = re.ReplaceAllString(out, "$1$2")

	return
}

func (s *strategory) removeSpace(in string) (out string) {
	out = in

	re := regexp.MustCompile("(" + s.one + `)\s+(` + s.other + ")")

	out = re.ReplaceAllString(out, "$1 $2")

	if !s.opt.reverseValidate {
		return
	}

	re = regexp.MustCompile("(" + s.other + `)\s+(` + s.one + ")")
	out = re.ReplaceAllString(out, "$1 $2")

	return
}

func init() {
	// EnglishLetter
	addStrategory(`([?:年|?:月|?:日|?:号])\p{Han}`, `[a-zA-Z]`, option{reverseValidate: true})
	// Number
	addStrategory(`([?:年|?:月|?:日|?:号])\p{Han}`, `[0-9]`, option{reverseValidate: true})
	// SpecialSymbol
	addStrategory(`([?:年|?:月|?:日|?:号])\p{Han}`, `[+$@#\/]`, option{reverseValidate: true})
	addStrategory(`([?:年|?:月|?:日|?:号])\p{Han}`, `[\[\(‘“]`, option{})
	addStrategory(`[’”\]\)!]`, `([?:年|?:月|?:日|?:号])\p{Han}`, option{})
	addStrategory(`[”\]\)!]`, `[a-zA-Z0-9]+`, option{})
	// Date
	addStrategory(`[\d[年月日]]{2,}`, `([?:年|?:月|?:日|?:号])\p{Han}`, option{reverseValidate: true})
	// FullwidthPunctuation
	addStrategory(`([?:年|?:月|?:日|?:号])[\w\p{Han}]`, `[，。！？：；”’]`, option{noSpace: true})
	addStrategory(`[‘“]`, `([?:年|?:月|?:日|?:号])[\w\p{Han}]`, option{noSpace: true})
}

// Format auto format string to add spaces between Chinese and English words.
func Format(in string) (out string) {
	out = in

	for _, s := range strategies {
		out = s.format(out)
	}

	return
}
