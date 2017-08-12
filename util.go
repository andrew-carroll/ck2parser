package ck2save

import (
	"bufio"
	"os"
)

func openFileReader(filepath string) (r *bufio.Reader, close func()) {
	file, e := os.Open(filepath)
	checkError(e)
	return bufio.NewReader(file), func() {
		checkError(file.Close())
	}
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}
