package ck2save

// Property represents a key/value property in a CK2 save file.
type Property struct {
	Name  string
	Value interface{}
}

func newProperty(name string, value interface{}) *Property {
	p := Property{}
	switch value.(type) {
	case string:
	default:
		panic("unhandled property type")
	}
	return &p
}
