package ck2save

import (
	"bufio"
	"io"
	"os"
)

func openFileReader(filepath string) (r *bufio.Reader, close func()) {
	file, e := os.Open(filepath)
	checkError(e)
	return bufio.NewReader(file), func() {
		checkError(file.Close())
	}
}

func openFileWriter(filepath string) (w *bufio.Writer, close func()) {
	file, e := os.Create(filepath)
	checkError(e)
	return bufio.NewWriter(file), func() {
		checkError(file.Close())
	}
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

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}
