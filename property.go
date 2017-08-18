package ck2save

// Property represents a key/value property in a CK2 save file.
type Property struct {
	Name        string
	Value       string
	property    []*Property
	propertyMap map[string]*Property
	pattern     pattern
}

func newProperty(name string, value string) *Property {
	p := Property{
		Name:        name,
		Value:       value,
		propertyMap: make(map[string]*Property),
	}
	return &p
}

// Property gets the Value of the specified member Property.
func (p *Property) Property(name string) *Property {
	return p.propertyMap[name]
}

// AddProperty adds the specified Property as a member.
func (p *Property) AddProperty(prop *Property) {
	p.property = append(p.property, prop)
	p.propertyMap[prop.Name] = prop
}
