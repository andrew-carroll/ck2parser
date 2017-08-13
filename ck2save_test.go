package ck2save

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var shortsave string = "./shortsave.ck2"

func TestParsesCK2Save(t *testing.T) {
	save := NewCK2Save(shortsave)
	t.Run("parses version", func(t *testing.T) {
		assert.Equal(t, `"2.7.1.0"`, save.property["version"])
	})
	t.Run("parses date", func(t *testing.T) {
		d := save.property["date"]
		assert.Equal(t, 2856, d.(ck2Date).year)
		assert.Equal(t, 5, d.(ck2Date).month)
		assert.Equal(t, 2, d.(ck2Date).day)
	})
}

func TestMapStorage(t *testing.T) {
	t.Run("newNamedMapSameLinePattern", func(t *testing.T) {
		p := "\t\tunborn={\n"
		s := CK2Save{}
		s.property = make(map[string]property)
		s.curPropMap = *newPropMap("test")
		s.parseLine(p)
	})
	t.Run("newUnnamedMapPattern", func(t *testing.T) {
	})
	t.Run("endMapPattern", func(t *testing.T) {
	})
}
