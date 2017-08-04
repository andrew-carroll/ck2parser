package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func closeFile(file *os.File) {
	checkError(file.Close())
}

func readFileLine(reader *bufio.Reader) string {
	line, err := reader.ReadString('\n')
	checkError(err)
	return line
}

func main() {
	file, err := os.Open(os.Args[1])
	checkError(err)
	defer closeFile(file)
	reader := bufio.NewReader(file)
	fmt.Print(readFileLine(reader))
	fmt.Print(readFileLine(reader))
	fmt.Print(readFileLine(reader))
}
