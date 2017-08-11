package ck2save

import (
	"bufio"
	"io"
	"os"
)

type player struct {
	id string
}

type CK2Save struct {
	property map[string]string
}

func NewCK2Save(filepath string) CK2Save {
	s := CK2Save{}
	s.property = make(map[string]string)

	f, e := os.Open(filepath)
	checkError(e)
	defer closeFile(f)
	s.readLines(bufio.NewReader(f))

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

func closeFile(f *os.File) {
	checkError(f.Close())
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}
