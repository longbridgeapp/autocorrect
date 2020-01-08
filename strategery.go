package autocorrect

import "regexp"

// Strategery for define rule
type strategery struct {
	one   string
	other string

	opt option
}

type option struct {
	space   bool
	reverse bool
}

func (s *strategery) format(in string) (out string) {
	out = in
	if s.opt.space {
		out = s.addSpace(out)
	} else {
		out = s.removeSpace(out)
	}

	return
}

func (s *strategery) addSpace(in string) (out string) {
	out = in

	re := regexp.MustCompile("(" + s.one + `)(` + s.other + ")")

	out = re.ReplaceAllString(out, "$1 $2")

	if s.opt.reverse {
		re = regexp.MustCompile("(" + s.other + `)(` + s.one + ")")
		out = re.ReplaceAllString(out, "$1 $2")
	}

	return
}

func (s *strategery) removeSpace(in string) (out string) {
	out = in

	re := regexp.MustCompile("(" + s.one + `)\s+(` + s.other + ")")

	out = re.ReplaceAllString(out, "$1 $2")

	if s.opt.reverse {
		re = regexp.MustCompile("(" + s.other + `)\s+(` + s.one + ")")
		out = re.ReplaceAllString(out, "$1 $2")
	}

	return
}
