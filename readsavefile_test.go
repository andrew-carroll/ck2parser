package ck2save

import (
	"github.com/stretchr/testify/assert"
	"io"
	"strings"
	"testing"
)

func TestReadsFile(t *testing.T) {
	r := strings.NewReader(shortsave)
	var lines int
	e := ReadSaveFile(r, func(s string) *Property {
		lines++
		return &Property{}
	})
	assert.Equal(t, io.EOF, e)
	assert.Equal(t, 14, lines)
}
