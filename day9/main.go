package main

import (
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/bAngerman/adventofcode/utils"
)

type point struct {
	x   int
	y   int
	val int
}

func checkItem(x int, y int, items [][]int) bool {
	var maxX, maxY, val int
	isLowPoint := false

	maxY = len(items)
	maxX = len(items[0])

	val = items[y][x]

	// Check top (y-1)
	// Check below (y+1)
	// Check left (x-1)
	// Check right (x+1)
	if (y == 0 || val < items[y-1][x]) && (y == maxY-1 || val < items[y+1][x]) && (x == 0 || val < items[y][x-1]) && (x == maxX-1 || val < items[y][x+1]) {
		isLowPoint = true
	}

	return isLowPoint
}

func part1(items [][]int) []point {
	var lowPoints []point
	var rowIdx, colIdx, maxRow, maxCol, lowPointTotal int

	if len(items) != 0 {
		maxCol = len(items)
		maxRow = len(items[0])
	} else {
		panic("No items passed.")
	}

	// log.Println("max row", maxRow)
	// log.Println("max col", maxCol)

	for {
		rowIdx = 0

		if colIdx >= maxCol {
			break
		}

		for {
			if rowIdx >= maxRow {
				break
			}

			if checkItem(rowIdx, colIdx, items) {
				// log.Println("y", colIdx, "\tx", rowIdx, "\t\tis low point", "\t\tValue:", items[colIdx][rowIdx])
				point := point{x: rowIdx, y: colIdx, val: items[colIdx][rowIdx]}
				lowPoints = append(lowPoints, point)
			}

			rowIdx++
		}

		colIdx++
	}

	for _, lowPoint := range lowPoints {
		lowPointTotal = lowPointTotal + (lowPoint.val + 1)
	}

	// log.Println("Low Points:", lowPoints)
	// log.Println("Total of low points (+1 ea)", lowPointTotal)

	return lowPoints
}

func inBasin(basin []point, item point) bool {

	for _, basinPoint := range basin {
		if basinPoint.x == item.x && basinPoint.y == item.y {
			return true
		}
	}

	return false
}

// Check the point in the basin (start with just the low point)
// Then evaluate points to the left, right, top, and bottom if they are set.
// Call this function recursively to continually grow outward
func growBasin(basin []point, items [][]int) []point {
	var maxX, maxY int
	didGrow := false

	maxY = len(items)
	maxX = len(items[0])

	for _, basinPoint := range basin {
		// Check point for surrounding points which both exist,
		// and are not "9" (which arent already in the basin)

		// Check top for a value. If point.y is 0 there are no points possible above.
		// If value is not 9, we can add it to the basin.
		if basinPoint.y != 0 && items[basinPoint.y-1][basinPoint.x] != 9 {
			// Check if it is already in our basin.
			abovePoint := point{x: basinPoint.x, y: basinPoint.y - 1, val: items[basinPoint.y-1][basinPoint.x]}
			if !inBasin(basin, abovePoint) {
				basin = append(basin, abovePoint)
				didGrow = true
			}
		}

		// Check right
		if basinPoint.x != maxX-1 && items[basinPoint.y][basinPoint.x+1] != 9 {
			// Check if it is already in our basin.
			rightPoint := point{x: basinPoint.x + 1, y: basinPoint.y, val: items[basinPoint.y][basinPoint.x+1]}
			if !inBasin(basin, rightPoint) {
				basin = append(basin, rightPoint)
				didGrow = true
			}
		}

		// Check bottom
		if basinPoint.y != maxY-1 && items[basinPoint.y+1][basinPoint.x] != 9 {
			// Check if it is already in our basin.
			bottomPoint := point{x: basinPoint.x, y: basinPoint.y + 1, val: items[basinPoint.y+1][basinPoint.x]}
			if !inBasin(basin, bottomPoint) {
				basin = append(basin, bottomPoint)
				didGrow = true
			}
		}

		// Check left
		if basinPoint.x != 0 && items[basinPoint.y][basinPoint.x-1] != 9 {
			// Check if it is already in our basin.
			leftPoint := point{x: basinPoint.x - 1, y: basinPoint.y, val: items[basinPoint.y][basinPoint.x-1]}
			if !inBasin(basin, leftPoint) {
				basin = append(basin, leftPoint)
				didGrow = true
			}
		}

	}

	// Didnt, grow, we must be done!
	if !didGrow {
		return basin
	}

	// Grow it again
	return growBasin(basin, items)
}

func part2(items [][]int) {
	var basins [][]point

	lowPoints := part1(items)

	// Now that we have low points... we can find basins which grow outward from the basin.
	for _, lowPoint := range lowPoints {
		var curBasin []point

		// Add our lowpoint
		curBasin = append(curBasin, lowPoint)
		// Grow it outward to edges
		curBasin = growBasin(curBasin, items)

		basins = append(basins, curBasin)
	}

	// pretty.Println(basins)

	var basinSizes []int

	for _, basin := range basins {
		basinSizes = append(basinSizes, len(basin))
	}

	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))
	var idx, top3 int

	for {
		if idx > 2 {
			break
		}
		if top3 == 0 {
			top3 = basinSizes[idx]
		} else {
			top3 = top3 * basinSizes[idx]
		}

		idx++
	}

	log.Println("Top 3 basins multipled", top3)
}

func main() {
	inputArr, _ := utils.ReadTextFile("day9/input")

	var items [][]int

	for _, item := range inputArr {
		points := strings.Split(item, "")
		var pointInts []int
		for _, point := range points {
			pointInt, _ := strconv.Atoi(point)
			pointInts = append(pointInts, pointInt)
		}

		items = append(items, pointInts)
	}

	// part1(items)
	part2(items)
}
