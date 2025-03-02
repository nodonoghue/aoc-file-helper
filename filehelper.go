package filehelper

import (
	"github.com/nodonoghue/filehelper/internal/readfile"
)

// Reads the given file and returns each line as []string
func ReadFileLines(filename string) ([]string, error) {
	return readfile.ReadFileLines(filename)
}

//TODO: Add additional funcs here
