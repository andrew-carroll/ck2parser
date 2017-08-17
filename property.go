package ck2save

import (
	"fmt"
)

// Property represents a key/value property in a CK2 save file.
type Property struct {
	Name        string
	Value       interface{}
	property    []*Property
	propertyMap map[string]*Property
}

func newProperty(name string, value interface{}) *Property {
	p := Property{Name: name, propertyMap: make(map[string]*Property)}
	switch v := value.(type) {
	case string:
		p.Value = value
	case *[]*Property:
		p.Value = value
	default:
		panic(fmt.Sprintf("unhandled property type %T", v))
	}
	return &p
}

func (p *Property) Property(name string) interface{} {
	prop := p.propertyMap[name]
	return prop.Value
}

func (p *Property) AddProperty(prop *Property) {
	p.property = append(p.property, prop)
	p.propertyMap[prop.Name] = prop
}
