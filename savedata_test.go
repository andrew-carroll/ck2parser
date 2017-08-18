package ck2save

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSaveDataAddProperty(t *testing.T) {
	var sd *SaveData = NewSaveData()
	var p *Property = newProperty("version", `"2.7.1.0"`, undefinedPattern)
	sd.AddProperty(p)
	assert.Equal(t, p, sd.Property("version"))
}

func TestSaveDataAddPropertyListProperty(t *testing.T) {
	var sd *SaveData = NewSaveData()
	var p *Property = newProperty("test", "", undefinedPattern)
	sd.AddProperty(p)
	p.AddProperty(newProperty("test", "one", undefinedPattern))
	p.AddProperty(newProperty("test2", "two", undefinedPattern))
	assert.Equal(t, "one", sd.Property("test").Property("test").Value)
	assert.Equal(t, "two", sd.Property("test").Property("test2").Value)
}
