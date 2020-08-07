package autocorrect

import "regexp"

// Strategery for define rule
type strategery struct {
	addSpaceRe           *regexp.Regexp
	addSpaceReverseRe    *regexp.Regexp
	removeSpaceRe        *regexp.Regexp
	removeSpaceReverseRe *regexp.Regexp

	space   bool
	reverse bool
}

func newStrategery(one, other string, space, reverse bool) *strategery {
	addSpaceStr := "(" + one + `)(` + other + ")"
	addSpaceReverseStr := "(" + other + `)(` + one + ")"

	removeSpaceStr := `(` + one + `)` + spaceRe.String() + `(` + other + `)`
	removeSpaceReverseStr := "(" + other + `)` + spaceRe.String() + `(` + one + ")"

	return &strategery{
		addSpaceRe:           regexp.MustCompile(addSpaceStr),
		addSpaceReverseRe:    regexp.MustCompile(addSpaceReverseStr),
		removeSpaceRe:        regexp.MustCompile(removeSpaceStr),
		removeSpaceReverseRe: regexp.MustCompile(removeSpaceReverseStr),
		space:                space,
		reverse:              reverse,
	}
}

func (s *strategery) format(in string) (out string) {
	out = in
	if s.space {
		out = s.addSpace(out)
	} else {
		out = s.removeSpace(out)
	}

	return
}

func (s *strategery) addSpace(in string) (out string) {
	out = in
	out = s.addSpaceRe.ReplaceAllString(out, "$1 $2")
	if s.reverse {
		out = s.addSpaceReverseRe.ReplaceAllString(out, "$1 $2")
	}

	return
}

func (s *strategery) removeSpace(in string) (out string) {
	out = in
	out = s.removeSpaceRe.ReplaceAllString(out, "$1 $2")
	if s.reverse {
		out = s.removeSpaceReverseRe.ReplaceAllString(out, "$1 $2")
	}

	return
}
