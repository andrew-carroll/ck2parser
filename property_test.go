package ck2save

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPropertyAddStringProperty(t *testing.T) {
	var p *Property = newProperty("parent", "")
	p.AddProperty(newProperty("foo", "bar"))
	assert.Equal(t, "bar", p.Property("foo").Value)
}

func TestPropertyAddPropertyListProperty(t *testing.T) {
	var p *Property = newProperty("parent", "")
	var c *Property = newProperty("child", "")
	var g *Property = newProperty("grandchild", "foo")
	p.AddProperty(c)
	c.AddProperty(g)
	assert.Equal(t, c, p.Property("child"))
	assert.Equal(t, g, c.Property("grandchild"))
	gg := p.Property("child").Property("grandchild").Value
	assert.Equal(t, "foo", gg)
}
