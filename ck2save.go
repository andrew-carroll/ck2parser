package ck2save

import (
	"bufio"
	"io"
)

type propMap struct {
	name        string
	property    map[string]string
	propMapList []propMap
}

type CK2Save struct {
	property    map[string]string
	propMapList []propMap
}

func NewCK2Save(filepath string) CK2Save {
	s := CK2Save{}
	s.property = make(map[string]string)

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
		s.property[l.name] = l.value
	}
}
