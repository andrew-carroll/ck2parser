package ck2save

import (
	"io"
)

type StringReader interface {
	ReadString(byte) (string, error)
}

func ReadSaveFile(s *SaveData, r StringReader, parseLine func(string) *Property) error {
	for {
		l, e := r.ReadString('\n')
		if e != nil {
			switch e {
			case io.EOF:
				return e
			default:
				panic(e)
			}
		}
		parseLine(l)
	}
}
