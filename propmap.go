package ck2save

type propMap struct {
	name        string
	property    map[string]property
	propMapList []propMap
	pattern     pattern
	parentIndex int
	index       int
}

func newPropMap(name string, pattern pattern, index int, parentIndex int) *propMap {
	p := propMap{}
	p.name = name
	p.index = index
	p.parentIndex = parentIndex
	p.pattern = pattern
	p.property = make(map[string]property)
	return &p
}

func (p *propMap) newPropMap(name string, pattern pattern, index int) *propMap {
	pm := *newPropMap(name, pattern, index, p.index)
	p.propMapList = append(p.propMapList, pm)
	return &pm
}
