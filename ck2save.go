package ck2save

type property interface{}

type CK2Save struct {
	property    map[string]property
	propMapList []*propMap
	curPropMap  *propMap
}

func NewCK2Save(filepath string) CK2Save {
	s := CK2Save{}
	s.property = make(map[string]property)
	s.curPropMap = newPropMap("root", undefinedPattern, 0, -1)
	s.propMapList = append(s.propMapList, s.curPropMap)

	r, fClose := openFileReader(filepath)
	defer fClose()
	s.readLines(r)

	return s
}

func (s *CK2Save) parseLine(rawLine string) {
	l := newCK2Line(rawLine)
	switch l.pattern {
	case newUnnamedMapPattern, newNamedMapSameLinePattern, newNamedMapPattern:
		s.newPropMap(l.name, l.pattern)
	case endMapPattern:
		s.closePropMap()
	case newPropPattern:
		switch l.propertyType {
		case propQuotedDate:
			s.property[l.name] = newCK2Date(l.value, true)
		case propUnquotedDate:
			s.property[l.name] = newCK2Date(l.value, false)
		default:
			s.property[l.name] = l.value
		}
	}
}

func (s *CK2Save) closePropMap() {
	parent := s.propMapList[s.curPropMap.parentIndex]
	s.curPropMap = parent
}

func (s *CK2Save) newPropMap(name string, pattern pattern) {
	index := len(s.propMapList)
	s.curPropMap = s.curPropMap.newPropMap(name, pattern, index)
	s.propMapList = append(s.propMapList, s.curPropMap)
}
