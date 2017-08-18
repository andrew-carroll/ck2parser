package ck2save

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatchesLinePatterns(t *testing.T) {
	var patternMatchTests = []struct {
		line    string
		pattern pattern
	}{
		{"CK2txt\n", headerPattern},
		{"\tversion=\"2.7.1.0\"\n", newPropPattern},
		{"\tplayer=\n", newNamedMapPattern},
		{"\t{\n", newUnnamedMapPattern},
		{"\t\tid=3022622\n", newPropPattern},
		{"\t}\n", endMapPattern},
		{"\t\n", emptyLinePattern},
		{"\t\tunborn={\n", newNamedMapSameLinePattern},
	}
	for _, tt := range patternMatchTests {
		p, e := ParseProperty(tt.line)
		assert.Nil(t, e)
		assert.Equal(t, tt.pattern, p.pattern)
	}
}

func TestParsesPropertyNamesAndValues(t *testing.T) {
	var propertyNameValueTests = []struct {
		line  string
		name  string
		value string
	}{
		{"CK2txt\n", "CK2txt", ""},
		{"\tversion=\"2.7.1.0\"\n", "version", `"2.7.1.0"`},
		{"\tplayer=\n", "player", ""},
		{"\t{\n", "", ""},
		{"\t\tid=3022622\n", "id", "3022622"},
		{"\t}\n", "", ""},
		{"\t\n", "", ""},
		{"\t\tunborn={\n", "unborn", ""},
	}
	for _, tt := range propertyNameValueTests {
		p, e := ParseProperty(tt.line)
		assert.Nil(t, e)
		assert.Equal(t, tt.name, p.Name)
		assert.Equal(t, tt.value, p.Value)
	}
}
