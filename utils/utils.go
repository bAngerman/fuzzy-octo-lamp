package utils

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func ReadNumberFile(relativePath string) ([]int, error) {

	absPath, _ := filepath.Abs("../" + relativePath)
	f, err := os.Open(absPath)

	if err != nil {
		log.Panic("Error reading input!: ", err)
	}

	defer f.Close()

	var lines []int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, num)
	}
	return lines, scanner.Err()
}
