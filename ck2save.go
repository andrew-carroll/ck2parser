package ck2save

import (
	"bufio"
	"io"
)

type property interface{}

type propMap struct {
	name        string
	property    map[string]property
	propMapList []propMap
}

type CK2Save struct {
	property    map[string]property
	propMapList []propMap
}

func NewCK2Save(filepath string) CK2Save {
	s := CK2Save{}
	s.property = make(map[string]property)

	r, fClose := openFileReader(filepath)
	defer fClose()
	s.readLines(r)

	return s
}

func (s *CK2Save) readLines(r *bufio.Reader) {
	for {
		l, e := r.ReadString('\n')
		if e == io.EOF {
			break
		}
		checkError(e)
		s.parseLine(l)
	}
}

func (s *CK2Save) parseLine(line string) {
	l := NewCK2Line(line)
	switch l.pattern {
	case newPropPattern:
		switch l.propertyType {
		case propQuotedDate:
			s.property[l.name] = NewCK2Date(l.value, true)
		case propUnquotedDate:
			s.property[l.name] = NewCK2Date(l.value, false)
		default:
			s.property[l.name] = l.value
		}
	}
}
