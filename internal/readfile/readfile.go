package readfile

import (
	"bufio"
	"os"
)

func ReadFileLines(file *os.File) []string {
	rlines := make([]string, 0)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rlines = append(rlines, scanner.Text())
	}

	return rlines
}
