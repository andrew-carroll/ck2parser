package ck2save

type propMap struct {
	name        string
	property    map[string]property
	propMapList []propMap
}

func newPropMap(name string) *propMap {
	p := propMap{}
	p.name = name
	return &p
}
