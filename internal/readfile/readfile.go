package readfile

import (
	"bufio"
	"os"
)

func ReadFileLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	rlines := make([]string, 0)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rlines = append(rlines, scanner.Text())
	}

	return rlines, nil
}
