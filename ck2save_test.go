package ck2save

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var shortsave string = "./shortsave.ck2"

func TestParsesCK2Save(t *testing.T) {
	save := NewCK2Save(shortsave)
	t.Run("parses version", func(t *testing.T) {
		assert.Equal(t, `"2.7.1.0"`, save.propMapList[0].propMapList[0].property["version"])
	})
	t.Run("parses date", func(t *testing.T) {
		d := save.propMapList[0].propMapList[0].property["date"]
		assert.Equal(t, 2856, d.(ck2Date).year)
		assert.Equal(t, 5, d.(ck2Date).month)
		assert.Equal(t, 2, d.(ck2Date).day)
	})
}

func newTestCK2Save(name string) *CK2Save {
	s := CK2Save{}
	s.property = make(map[string]property)
	s.curPropMap = newPropMap("test", undefinedPattern, 0, -1)
	s.propMapList = append(s.propMapList, s.curPropMap)
	return &s
}

func TestMapStorage(t *testing.T) {
	t.Run("newNamedMapSameLinePattern", func(t *testing.T) {
		p := "\t\t\tunborn={\n"
		s := newTestCK2Save("test")
		assert.Equal(t, "test", s.curPropMap.name)
		assert.Equal(t, undefinedPattern, s.curPropMap.pattern)
		s.parseLine(p)
		par := s.propMapList[s.curPropMap.parentIndex]
		assert.Equal(t, newNamedMapSameLinePattern, s.curPropMap.pattern)
		assert.Equal(t, "test", par.name)
		assert.Equal(t, "unborn", s.curPropMap.name)
	})
	t.Run("newNamedMapPattern", func(t *testing.T) {
		p := "\tplayer=\n"
		s := newTestCK2Save("test")
		assert.Equal(t, s.curPropMap.name, "test")
		s.parseLine(p)
		par := s.propMapList[s.curPropMap.parentIndex]
		assert.Equal(t, newNamedMapPattern, s.curPropMap.pattern)
		assert.Equal(t, "test", par.name)
		assert.Equal(t, "player", s.curPropMap.name)
	})
	t.Run("newUnnamedMapPattern", func(t *testing.T) {
		p := "\t{\n"
		s := newTestCK2Save("test")
		s.curPropMap.pattern = newNamedMapSameLinePattern
		s.parseLine(p)
		assert.Equal(t, "", s.curPropMap.name)
		assert.Equal(t, newUnnamedMapPattern, s.curPropMap.pattern)
	})
	t.Run("endMapPattern", func(t *testing.T) {
		p := "\t}\n"
		s := newTestCK2Save("test")
		s.curPropMap.pattern = newUnnamedMapPattern
		s.newPropMap("player", newNamedMapPattern)
		assert.Equal(t, "player", s.curPropMap.name)
		s.parseLine(p)
		assert.Equal(t, "test", s.curPropMap.name)
	})
}

func TestPrintsSaveFile(t *testing.T) {
	s := newTestCK2Save("test")
	s.parseLine("CK2txt\n")
	s.parseLine("version=\"2.7.1.0\"\n")
	s.parseLine("date=\"2856.5.2\"\n")
	s.parseLine("player=\n")
	s.parseLine("{\n")
	s.parseLine("id=100\n")
	s.parseLine("type=66\n")
	s.parseLine("}\n")
	s.parseLine("}\n")
	expected := "CK2txt\n" +
		"version=\"2.7.1.0\"\n" +
		"date=\"2856.5.2\"\n" +
		"player=\n" +
		"{\n" +
		"id=100\n" +
		"type=66\n" +
		"}\n" +
		"}\n"
	builder := newCK2SaveFileBuilder(s)
	actual := builder.buildSaveFile()
	assert.Equal(t, expected, actual)
}
