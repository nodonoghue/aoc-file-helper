package filehelper

import "github.com/nodonoghue/aoc-file-helper/internal/utils"

// Pass in a file path to get a slice of file lines to pass through the parsing func of your choice
func ReadFileLines(filename string) []string {
	file := utils.OpenFile(filename)
	return utils.ReadFile(file)
}
