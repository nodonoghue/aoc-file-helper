package aocfilehelper

import "github.com/nodonoghue/aoc-file-helper/internal/filehelper"

// Pass in a file path to get a slice of file lines to pass through the parsing func of your choice
func ReadFileLines(filename string) []string {
	file := filehelper.OpenFile(filename)
	return filehelper.ReadFile(file)
}
