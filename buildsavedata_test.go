package ck2save

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestBuildSaveData(t *testing.T) {
	r := strings.NewReader(shortsave)
	s := BuildSaveData(r)
	h := s.Property("CK2txt")
	u := h.Property("unborn")
	assert.Equal(t, headerPattern, h.pattern)
	// Here we compare using property[] instead of Property() because
	// end brackets will override propertyMap[""] while property[0]
	// is guaranteed to be the open bracket to a newNamedMap.
	assert.Equal(t, newNamedMapPattern, u.pattern)
	assert.Equal(t, newUnnamedMapPattern, u.property[0].pattern)
	// Similarly, Property("") is guratanteed to be the end bracket to a newNamedMap.
	assert.Equal(t, endMapPattern, u.Property(""))
	// However, newNamedMapSameLine holds no such guarantee.
	p := h.Property("player")
	assert.Equal(t, newNamedMapSameLinePattern, p.pattern)
	assert.Equal(t, newPropPattern, p.property[0].pattern)

	assert.Equal(t, endMapPattern, h.Property(""))
}
