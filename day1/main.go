package main

import (
	"log"

	"github.com/bAngerman/adventofcode/utils"
)

func countIncreases(input []int) int {
	var count = 0

	for idx, num := range input {
		if idx == 0 {
			continue
		}

		if num > input[idx-1] {
			count = count + 1
		} else {
		}
	}

	return count
}

func partOne() {
	inputArr, _ := utils.ReadNumberFile("day1/input")

	increasedCount := countIncreases(inputArr)

	log.Println("Part One: Increased count is", increasedCount)
}

func partTwo() {
	inputArr, _ := utils.ReadNumberFile("day1/input")
	var windowSize = 3

	windows := make([]int, 0)

	for idx, _ := range inputArr {
		if (idx + windowSize) > len(inputArr) {
			continue
		}

		window := inputArr[idx:(idx + windowSize)]
		windowSum := 0

		for _, windowVal := range window {
			windowSum = windowSum + windowVal
		}

		windows = append(windows, windowSum)
	}

	increasedCount := countIncreases(windows)

	log.Println("Part Two: Increased count is", increasedCount)
}

func main() {
	// partOne()
	partTwo()
}
