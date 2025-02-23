package filehelper

import (
	"os"

	"github.com/nodonoghue/filehelper/internal/readfile"
	"github.com/nodonoghue/filehelper/internal/utils"
)

// Reads the given file and returns each line as []string
func ReadFileLines(filename string) ([]string, error) {
	file, err := utils.OpenFile(filename)
	if err != nil {
		return nil, err
	}
	return readfile.ReadFileLines(file), nil
}

// Opens a file and returns as *os.File
func OpenFile(filename string) (*os.File, error) {
	return utils.OpenFile(filename)
}
