package ck2save

import (
	"github.com/stretchr/testify/assert"
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

func TestSaveDataAddProperty(t *testing.T) {
	var s *SaveData = NewSaveData()
	s.AddProperty("version", `"2.7.1.0"`)
	assert.Equal(t, property(`"2.7.1.0"`), s.Property("version"))
}
