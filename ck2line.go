package ck2save

import (
	"regexp"
)

type CK2Line struct {
	rawLine string
	pattern Pattern
	name    string
	value   string
}

type Pattern string

const (
	headerPattern              Pattern = "headerPattern"
	emptyLinePattern           Pattern = "emptyLinePattern"
	newNamedMapPattern         Pattern = "newNamedMapPattern"
	newUnnamedMapPattern       Pattern = "newUnnamedMapPattern"
	newNamedMapSameLinePattern Pattern = "newNamedMapSameLinePattern"
	endMapPattern              Pattern = "endMapPattern"
	newPropPattern             Pattern = "newPropPattern"
	undefinedPattern           Pattern = "undefinedPattern"
)

var w string = `([\w\d\_\-\.]+)`
var reg = map[Pattern]*regexp.Regexp{
	headerPattern:              regexp.MustCompile(`\ACK2txt\n`),
	emptyLinePattern:           regexp.MustCompile(`^\t*\n`),
	newNamedMapPattern:         regexp.MustCompile(`^\t*` + w + `=\n`),
	newNamedMapSameLinePattern: regexp.MustCompile(`^\t*` + w + `=\{\n`),
	newUnnamedMapPattern:       regexp.MustCompile(`^\t*\{\n`),
	endMapPattern:              regexp.MustCompile(`^\t*}\n`),
	newPropPattern:             regexp.MustCompile(`^\t*` + w + `=([^{].*)\n`),
}

func (ck2line *CK2Line) DeterminePattern() {
	ck2line.pattern = undefinedPattern
	for n, r := range reg {
		if r.MatchString(ck2line.rawLine) {
			ck2line.pattern = n
			break
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
	case newUnnamedMapPattern, headerPattern, endMapPattern, emptyLinePattern, undefinedPattern:
	default:
		panic("Expected ck2line.pattern to be defined")
	}

	return ck2line
}
