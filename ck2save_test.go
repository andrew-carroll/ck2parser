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
	t.Run("parses date", func(t *testing.T) {
		d := NewCK2Date("2856.5.2")
		expectValue(t, 2856, d.year)
		expectValue(t, 5, int(d.month))
		expectValue(t, 2, d.day)
	})
}

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

func TestPropertyTypes(t *testing.T) {
	var propertyTypeTests = []struct {
		line         string
		propertyType propertyType
	}{
		{"\tversion=\"2.7.1.0\"\n", propString},
		{"\tdate=\"2856.5.2\"\n", propQuotedDate},
		{"\t\tid=3022622\n", propInteger},
		{"\t\tdata={0 8 0 30 0 0 0}\n", propIntArray},
		{"\tis_zeus_save=yes\n", propBool},
		{"\t\tscenario_initialized=2666.7.4\n", propUnquotedDate},
		{"\t\t\ttraits={266 }\n", propIntArray},
		{"\t\t\thealth=5.800\n", propDecimal},
		{"\t\t\tc_d=death_trait\n", propString},
		{"\t\t\t\tgelleys_f={27.292 33.000}\n", propDecArray},
	}
	for _, tt := range propertyTypeTests {
		l := NewCK2Line(tt.line)
		if l.propertyType != tt.propertyType {
			t.Errorf("NewCK2Line(%q).propertyType => %q,\t\twant %q (name: %q | value: %q)", tt.line, l.propertyType, tt.propertyType, l.name, l.value)
		}
	}
}

func TestOpensMaps(t *testing.T) {

}

func TestMapStorage(t *testing.T) {
	t.Run("newNamedMapSameLinePattern", func(t *testing.T) {
	})
	t.Run("newUnnamedMapPattern", func(t *testing.T) {
	})
	t.Run("endMapPattern", func(t *testing.T) {
	})
}
