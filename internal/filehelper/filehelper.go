package filehelper

import (
	"bufio"
	"os"
)

func OpenFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	return file
}

func ReadFile(file *os.File) []string {
	rlines := make([]string, 0)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rlines = append(rlines, scanner.Text())
	}

	return rlines
}
