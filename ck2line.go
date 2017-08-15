package ck2save

type ck2Line struct {
	rawLine      string
	pattern      pattern
	name         string
	value        string
	propertyType propertyType
}

type pattern string
type propertyType string

func (ck2line *ck2Line) determinePattern() {
	ck2line.pattern = undefinedPattern
	for n, r := range reg {
		if r.MatchString(ck2line.rawLine) {
			ck2line.pattern = n
			break
		}
	}
	switch { // override false-positive newPropPattern matches
	case reg[newNamedMapSameLinePattern].MatchString(ck2line.rawLine):
		ck2line.pattern = newNamedMapSameLinePattern
	case reg[newUnnamedMapPattern].MatchString(ck2line.rawLine):
		ck2line.pattern = newUnnamedMapPattern
	}
}

func (ck2line *ck2Line) determinePropertyType() {
	ck2line.propertyType = propString
	for p, r := range propReg {
		if r.MatchString(ck2line.value) {
			ck2line.propertyType = p
		}
	}
	if ck2line.propertyType == propString {
		switch {
		case propReg[propQuotedDate].MatchString(ck2line.value):
			ck2line.propertyType = propQuotedDate
		case propReg[propUnquotedDate].MatchString(ck2line.value):
			ck2line.propertyType = propUnquotedDate
		}
	}
}

func newCK2Line(s string) ck2Line {
	l := ck2Line{}
	l.rawLine = s
	l.determinePattern()
	switch l.pattern {
	case newNamedMapPattern:
		matches := reg[newNamedMapPattern].FindAllStringSubmatch(s, -1)[0]
		l.name = matches[1]
	case newNamedMapSameLinePattern:
		matches := reg[newNamedMapSameLinePattern].FindAllStringSubmatch(s, -1)[0]
		l.name = matches[1]
	case newPropPattern:
		matches := reg[newPropPattern].FindAllStringSubmatch(s, -1)[0]
		l.name, l.value = matches[1], matches[2]
		l.determinePropertyType()
	case newUnnamedMapPattern, headerPattern, endMapPattern, emptyLinePattern, undefinedPattern:
	default:
		panic("Expected ck2line.pattern to be defined")
	}

	return l
}
