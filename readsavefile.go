package ck2save

import (
	"bufio"
	"io"
)

func ReadSaveFile(r io.Reader, parseLine func(string)) error {
	b := bufio.NewReader(r)
	for {
		l, e := b.ReadString('\n')
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
