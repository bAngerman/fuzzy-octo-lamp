package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/bAngerman/adventofcode/utils"
)

type line struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

// func removeLine(s []line, i int) []line {
// 	s[i] = s[len(s)-1]
// 	return s[:len(s)-1]
// }

func filterOnlyHorizontalVerticalLines(lines []line) []line {
	var idx int
	var filteredLines []line

	for {
		if lines[idx].x1 == lines[idx].x2 || lines[idx].y1 == lines[idx].y2 {
			filteredLines = append(filteredLines, lines[idx])
		}

		idx++

		if idx > len(lines)-1 {
			break
		}
	}

	return filteredLines
}

// func checkOverlaps(lineOne line, lineTwo line) int {
// 	var overlaps int
// 	var segment1, segment2 []int

// 	// These are vertically lines overlapping
// 	if lineOne.x1 == lineTwo.x1 && lineOne.x2 == lineTwo.x2 {
// 		// log.Println("Check line\t\t", lineOne, "\ncrosses other line\t\t\t", lineTwo, "\vertically")
// 		segment1 = makeRange(lineOne.y1, lineOne.y2)
// 		segment2 = makeRange(lineTwo.y1, lineTwo.y2)

// 		overlaps = len(intersect.Simple(segment1, segment2))
// 	}

// 	// These are horizontal lines overlapping
// 	if lineOne.y1 == lineTwo.y1 && lineOne.y2 == lineTwo.y2 {
// 		// log.Println("Check line\t\t", lineOne, "\ncrosses other line\t\t\t", lineTwo, "\nhorizontally")
// 		segment1 = makeRange(lineOne.x1, lineOne.x2)
// 		segment2 = makeRange(lineTwo.x1, lineTwo.x2)

// 		overlaps = len(intersect.Simple(segment1, segment2))
// 	}

// 	// Need to check for intersections of the lines now.
// 	// This checks for the specific scenario when a lineOne is a vertical line,
// 	// and lineTwo is a horizontal line.
// 	// The first statement checks for lineOne having x valuese within the bounds of lineTwo
// 	//
// 	// Line One: 7,0 -> 7,4
// 	// Line Two: 3,4 -> 9,4

// 	if lineOne.x1 == lineTwo.x1 || line {
// 		// log.Println(lineOne, lineTwo)
// 		// This condition checks for the vertical intersect.
// 		if lineOne.y1 <= lineTwo.y1 && lineOne.y2 <= lineTwo.y2 {
// 			if overlaps == 0 {
// 				overlaps = 1
// 			}
// 		}
// 	}

// 	return overlaps
// }

// This is woefully inefficient.
func part1(lines []line) {
	var overlaps, currentOverlaps, maxX, maxY int
	var xRange, yRange []int

	lines = filterOnlyHorizontalVerticalLines(lines)

	for _, item := range lines {
		if item.x1 > maxX {
			maxX = item.x1
		}
		if item.x2 > maxX {
			maxX = item.x2
		}
		if item.y1 > maxY {
			maxY = item.y1
		}
		if item.y2 > maxY {
			maxY = item.y2
		}
	}

	// pretty.Println("Max X:", maxX, "Max Y:", maxY)

	// This is woefully inefficient.
	for _, xCoord := range makeRange(0, maxX) {
		// log.Println("X coordinate:", xCoord)
		for _, yCoord := range makeRange(0, maxY) {

			currentOverlaps = 0
			for _, item := range lines {

				xRange = makeRange(item.x1, item.x2)
				yRange = makeRange(item.y1, item.y2)

				for _, itemX := range xRange {
					for _, itemY := range yRange {
						if xCoord == itemX && yCoord == itemY {
							currentOverlaps = currentOverlaps + 1
						}
					}
				}
			}

			if currentOverlaps > 1 {
				// log.Println("Lines intersect")
				overlaps = overlaps + currentOverlaps - 1
			}

		}
	}

	log.Println("Number of overlaps:", overlaps)

	return
}

func part2(lines []line) {

}

func processInput(inputArr []string) []line {
	var lines []line

	for _, input := range inputArr {
		halves := strings.Split(input, "-> ")
		firstCoords := strings.Split(halves[0], ",")
		for i := range firstCoords {
			firstCoords[i] = strings.TrimSpace(firstCoords[i])
		}

		secondCoords := strings.Split(halves[1], ",")
		for i := range secondCoords {
			secondCoords[i] = strings.TrimSpace(secondCoords[i])
		}

		x1, _ := strconv.Atoi(firstCoords[0])
		y1, _ := strconv.Atoi(firstCoords[1])
		x2, _ := strconv.Atoi(secondCoords[0])
		y2, _ := strconv.Atoi(secondCoords[1])

		if x1 > x2 {
			tmp := x2
			x2 = x1
			x1 = tmp
		}

		if y1 > y2 {
			tmp := y2
			y2 = y1
			y1 = tmp
		}

		lines = append(lines, line{x1: x1, x2: x2, y1: y1, y2: y2})
	}

	return lines
}

func main() {
	inputArr, _ := utils.ReadTextFile("day5/input2")
	lines := processInput(inputArr)

	part1(lines)

	// log.Println(inputArr)

	// part1()
	// part2()
}
