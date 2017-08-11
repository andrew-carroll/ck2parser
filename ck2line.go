package ck2save

import (
	"regexp"
)

type CK2Line struct {
	rawLine      string
	pattern      pattern
	name         string
	value        string
	propertyType propertyType
}

type pattern string
type propertyType string

const (
	headerPattern              pattern      = "headerPattern"
	emptyLinePattern           pattern      = "emptyLinePattern"
	newNamedMapPattern         pattern      = "newNamedMapPattern"
	newUnnamedMapPattern       pattern      = "newUnnamedMapPattern"
	newNamedMapSameLinePattern pattern      = "newNamedMapSameLinePattern"
	endMapPattern              pattern      = "endMapPattern"
	newPropPattern             pattern      = "newPropPattern"
	undefinedPattern           pattern      = "undefinedPattern"
	propQuotedDate             propertyType = "propQuotedDate"
	propUnquotedDate           propertyType = "propUnquotedDate"
	propInteger                propertyType = "propInteger"
	propIntArray               propertyType = "propIntArray"
	propDecimal                propertyType = "propDecimal"
	propDecArray               propertyType = "propDecArray"
	propBool                   propertyType = "propBool"
	propString                 propertyType = "propString"
)

var w string = `([\w\d\_\-\.]+)`
var reg = map[pattern]*regexp.Regexp{
	headerPattern:              regexp.MustCompile(`\ACK2txt\n$`),
	emptyLinePattern:           regexp.MustCompile(`^\t*\n$`),
	newNamedMapPattern:         regexp.MustCompile(`^\t*` + w + `=\n$`),
	newNamedMapSameLinePattern: regexp.MustCompile(`^\t*` + w + `=\{\n$`),
	newUnnamedMapPattern:       regexp.MustCompile(`^\t*\{\n$`),
	endMapPattern:              regexp.MustCompile(`^\t*}\n$`),
	newPropPattern:             regexp.MustCompile(`^\t*` + w + `=(.+)\n$`),
}

func (ck2line *CK2Line) DeterminePattern() {
	ck2line.pattern = undefinedPattern
	for n, r := range reg {
		if r.MatchString(ck2line.rawLine) {
			ck2line.pattern = n
			break
		}
	}
	if reg[newNamedMapSameLinePattern].MatchString(ck2line.rawLine) {
		ck2line.pattern = newNamedMapSameLinePattern
	}
	if reg[newUnnamedMapPattern].MatchString(ck2line.rawLine) {
		ck2line.pattern = newUnnamedMapPattern
	}
}

var propDate string = `(\d{3,4})\.(\d{1,2})\.(\d{1,2})`
var propReg = map[propertyType]*regexp.Regexp{
	propQuotedDate:   regexp.MustCompile(`^"` + propDate + `"$`),
	propUnquotedDate: regexp.MustCompile(`^` + propDate + `$`),
	propInteger:      regexp.MustCompile(`^([1-9]\d*)$`),
	propIntArray:     regexp.MustCompile(`^{((?:\d+\s?)+)}$`),
	propDecimal:      regexp.MustCompile(`^\d+\.\d+$`),
	propDecArray:     regexp.MustCompile(`^{((?:\d+\.\d+\s?)+)}$`),
	propBool:         regexp.MustCompile(`^yes`),
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
