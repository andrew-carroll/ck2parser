package ck2save

type property interface{}

type CK2Save struct {
	Property    map[string]property
	PropMapList []*propMap
	curPropMap  *propMap
}

func NewCK2Save(filepath string) CK2Save {
	s := CK2Save{}
	s.Property = make(map[string]property)
	s.curPropMap = newPropMap("root", undefinedPattern, 0, -1)
	s.PropMapList = append(s.PropMapList, s.curPropMap)

	r, fClose := openFileReader(filepath)
	defer fClose()
	s.readLines(r)

	return s
}

func (s *CK2Save) parseLine(rawLine string) {
	l := newCK2Line(rawLine)
	switch l.pattern {
	case headerPattern:
		s.newPropMap("CK2txt", l.pattern)
	case newUnnamedMapPattern, newNamedMapSameLinePattern:
		s.newPropMap(l.name, l.pattern)
	case newNamedMapPattern:
		s.newPropMap(l.name, l.pattern)
	case endMapPattern:
		s.closePropMap()
		if s.curPropMap.pattern == newNamedMapPattern {
			s.closePropMap()
		}
	case newPropPattern:
		switch l.propertyType {
		case propQuotedDate:
			s.curPropMap.setProperty(l.name, newCK2Date(l.value, true))
		case propUnquotedDate:
			s.curPropMap.setProperty(l.name, newCK2Date(l.value, false))
		default:
			s.curPropMap.setProperty(l.name, l.value)
		}
	}
}

func (s *CK2Save) parent(p *propMap) *propMap {
	return s.PropMapList[p.parentIndex]
}

func (s *CK2Save) closePropMap() {
	s.curPropMap = s.parent(s.curPropMap)
}

func (s *CK2Save) newPropMap(name string, pattern pattern) {
	index := len(s.PropMapList)
	s.curPropMap = s.curPropMap.newPropMap(name, pattern, index)
	s.PropMapList = append(s.PropMapList, s.curPropMap)
}
