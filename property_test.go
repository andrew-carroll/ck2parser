package ck2save

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPropertyAddStringProperty(t *testing.T) {
	var p *Property = newProperty("parent", "", undefinedPattern)
	p.AddProperty(newProperty("foo", "bar", undefinedPattern))
	assert.Equal(t, "bar", p.Property("foo").Value)
}

func TestPropertyAddPropertyListProperty(t *testing.T) {
	var p *Property = newProperty("parent", "", undefinedPattern)
	var c *Property = newProperty("child", "", undefinedPattern)
	var g *Property = newProperty("grandchild", "foo", undefinedPattern)
	p.AddProperty(c)
	c.AddProperty(g)
	assert.Equal(t, c, p.Property("child"))
	assert.Equal(t, g, c.Property("grandchild"))
	gg := p.Property("child").Property("grandchild").Value
	assert.Equal(t, "foo", gg)
}

func TestPropertyAddPropertyAddsParent(t *testing.T) {
	var p *Property = newProperty("parent", "", undefinedPattern)
	var c *Property = newProperty("child", "", undefinedPattern)
	p.AddProperty(c)
	assert.Equal(t, c.Parent, p)
}
