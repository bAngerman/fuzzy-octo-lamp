package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/bAngerman/adventofcode/utils"
	"github.com/kr/pretty"
)

func incrementItem(items [][]int, x int, y int, flashes int) ([][]int, int) {
	var maxX, maxY int

	if len(items) != 0 {
		maxY = len(items)
		maxX = len(items[0])
	}

	// Skip, this already flashed and is limited to one flash
	if items[y][x] <= 9 {
		items[y][x] = items[y][x] + 1

		// pretty.Println("Item to increment is", items[y][x])

		if items[y][x] > 9 {
			// Apply flash to surroundings if they exist.

			flashes = flashes + 1

			// Left
			if x != 0 {
				items, flashes = incrementItem(items, x-1, y, flashes)
				// log.Println("Left Flashes", flashes)

				// Top left
				if y != 0 {
					items, flashes = incrementItem(items, x-1, y-1, flashes)
				}

				// Bottom left
				if y != maxY-1 {
					items, flashes = incrementItem(items, x-1, y+1, flashes)
				}
			}

			// Right
			if x != maxX-1 {
				items, flashes = incrementItem(items, x+1, y, flashes)
				// log.Println("Right Flashes", flashes)

				// Top right
				if y != 0 {
					items, flashes = incrementItem(items, x+1, y-1, flashes)
				}

				// Top right
				if y != maxY-1 {
					items, flashes = incrementItem(items, x+1, y+1, flashes)
				}
			}

			// Top
			if y != 0 {
				items, flashes = incrementItem(items, x, y-1, flashes)
				// log.Println("Top Flashes", flashes)
			}

			// Below
			if y != maxY-1 {
				items, flashes = incrementItem(items, x, y+1, flashes)
			}
		}
	}

	return items, flashes
}

func part1(items [][]int) {
	var x, y, step, maxStep, maxX, maxY, flashes int

	if len(items) != 0 {
		maxY = len(items)
		maxX = len(items[0])
	} else {
		panic("Invalid items")
	}

	maxStep = 100

	pretty.Println("Initial state", items)

	for {
		if step >= maxStep {
			break
		}

		y = 0
		for {
			if y >= maxY {
				break
			}

			x = 0
			for {
				if x >= maxX {
					break
				}

				newItems, itemFlashes := incrementItem(items, x, y, 0)

				items = newItems
				flashes = flashes + itemFlashes

				x++
			}
			y++
		}

		for yIdx, row := range items {
			for xIdx, item := range row {
				// Set all the flashers to 0
				if item > 9 {
					items[yIdx][xIdx] = 0
				}
			}
		}

		pretty.Println("After Step", step+1, "Flashes", flashes)
		pretty.Println(items)
		step++
	}

	pretty.Println("Number of flashes", flashes)
}

func flashesAreAligned(items [][]int) bool {
	matchVal := items[0][0]

	for _, row := range items {
		for _, val := range row {
			if matchVal != val {
				return false
			}
		}
	}

	return true
}

func part2(items [][]int) {
	var x, y, step, maxX, maxY, flashes int

	if len(items) != 0 {
		maxY = len(items)
		maxX = len(items[0])
	} else {
		panic("Invalid items")
	}

	for {
		if flashesAreAligned(items) {
			log.Println("Flashes aligned after", step, "steps")
			pretty.Println(items)
			break
		}

		y = 0
		for {
			if y >= maxY {
				break
			}

			x = 0
			for {
				if x >= maxX {
					break
				}

				newItems, itemFlashes := incrementItem(items, x, y, 0)

				items = newItems
				flashes = flashes + itemFlashes

				x++
			}
			y++
		}

		for yIdx, row := range items {
			for xIdx, item := range row {
				// Set all the flashers to 0
				if item > 9 {
					items[yIdx][xIdx] = 0
				}
			}
		}

		step++
	}
}

func main() {
	inputArr, _ := utils.ReadTextFile("day11/input")

	var items [][]int

	for _, row := range inputArr {
		rowString := strings.Split(row, "")
		var numberRow []int

		for _, str := range rowString {
			num, _ := strconv.Atoi(str)
			numberRow = append(numberRow, num)
		}

		items = append(items, numberRow)
	}

	// part1(items)
	part2(items)
}
