package ck2save

import (
	"regexp"
)

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
	newPropPattern:             regexp.MustCompile(`^\t*` + w + `=("?.+"?)\n$`),
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
