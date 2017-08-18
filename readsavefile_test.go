package ck2save

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"io"
	"strings"
	"testing"
)

var shortsave = `CK2txt
	version="2.7.1.0"
	date="2856.5.2"
	player=
	{
		id=3022622
		type=66
	}
	player_realm="e_gothamite"
	is_zeus_save=yes
	generated_societies=yes
	generated_artifacts=yes
	vc_data="CK2/branches/2_7_1\\n37283\\n"
}
`

func TestReadsFile(t *testing.T) {
	s := NewSaveData()
	r := bufio.NewReader(strings.NewReader(shortsave))
	var lines int
	e := ReadSaveFile(s, r, func(s string) *Property {
		lines++
		return &Property{}
	})
	assert.Equal(t, io.EOF, e)
	assert.Equal(t, 14, lines)
}
