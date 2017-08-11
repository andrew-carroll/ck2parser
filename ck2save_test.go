package ck2save

import (
	"testing"
)

func expectValue(t *testing.T, want interface{}, got interface{}) {
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

var shortsave string = "./shortsave.ck2"

func TestParsesCK2Save(t *testing.T) {
	save := NewCK2Save(shortsave)
	t.Run("parses version", func(t *testing.T) {
		expectValue(t, `"2.7.1.0"`, save.property["version"])
	})
}

var patternMatchTests = []struct {
	line    string
	pattern Pattern
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

func TestMatchesLinePatterns(t *testing.T) {
	for _, tt := range patternMatchTests {
		l := NewCK2Line(tt.line)
		if l.pattern != tt.pattern {
			t.Errorf("NewCK2Line(%q).pattern => %q, want %q", tt.line, l.pattern, tt.pattern)
		}
	}
}

func TestParsesPropKeysAndValues(t *testing.T) {
	line := "\t\tid=3022622\n"
	l := NewCK2Line(line)
	k, v := l.name, l.value
	wk, wv := "id", "3022622"
	if v != wv {
		t.Errorf("NewCK2Line(%q).value => %q, want %q", line, v, wv)
	}
	if k != wk {
		t.Errorf("NewCK2Line(%q).name => %q, want %q", line, k, wk)
	}
}
