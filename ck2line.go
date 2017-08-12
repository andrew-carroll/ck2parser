package ck2save

type CK2Line struct {
	rawLine      string
	pattern      pattern
	name         string
	value        string
	propertyType propertyType
}

type pattern string
type propertyType string

func (ck2line *CK2Line) DeterminePattern() {
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

func (ck2line *CK2Line) DeterminePropertyType() {
	ck2line.propertyType = propString
	for p, r := range propReg {
		if r.MatchString(ck2line.value) {
			ck2line.propertyType = p
		}
	}
}

func NewCK2Line(l string) CK2Line {
	ck2line := CK2Line{}
	ck2line.rawLine = l
	ck2line.DeterminePattern()
	switch ck2line.pattern {
	case newNamedMapPattern:
		matches := reg[newNamedMapPattern].FindAllStringSubmatch(l, -1)[0]
		ck2line.name = matches[1]
	case newNamedMapSameLinePattern:
		matches := reg[newNamedMapSameLinePattern].FindAllStringSubmatch(l, -1)[0]
		ck2line.name = matches[1]
	case newPropPattern:
		matches := reg[newPropPattern].FindAllStringSubmatch(l, -1)[0]
		ck2line.name, ck2line.value = matches[1], matches[2]
		ck2line.DeterminePropertyType()
	case newUnnamedMapPattern, headerPattern, endMapPattern, emptyLinePattern, undefinedPattern:
	default:
		panic("Expected ck2line.pattern to be defined")
	}

	return ck2line
}
