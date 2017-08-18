package ck2save

type PropertyHolder interface {
	Property(string) *Property
	AddProperty(*Property)
}

type UnhandledLinePatternError struct {
	pattern pattern
}

func (e UnhandledLinePatternError) Error() string {
	return string(e.pattern)
}

func ParseProperty(l string) (*Property, error) {
	var name, value string
	pat := parsePattern(l)
	matches := reg[pat].FindAllStringSubmatch(l, -1)[0]
	switch pat {
	case newNamedMapPattern, newNamedMapSameLinePattern:
		name = matches[1]
	case newPropPattern:
		name, value = matches[1], matches[2]
	case headerPattern:
		name = "CK2txt"
	case newUnnamedMapPattern, endMapPattern, emptyLinePattern, undefinedPattern:
	default:
		return nil, UnhandledLinePatternError{pat}
	}
	p := newProperty(name, value, pat)
	p.pattern = pat
	return p, nil
}

func parsePattern(l string) pattern {
	p := undefinedPattern
	for pn, pr := range reg {
		if pr.MatchString(l) {
			p = pn
			break
		}
	}
	switch { // Bypass false positives
	case reg[newNamedMapSameLinePattern].MatchString(l):
		return newNamedMapSameLinePattern
	case reg[newUnnamedMapPattern].MatchString(l):
		return newUnnamedMapPattern
	default:
		return p
	}
}
