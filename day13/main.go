package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/bAngerman/adventofcode/utils"
)

type command struct {
	foldAlong string
	offset    int
}

func remove(s [][]int, i int) [][]int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func dotExists(dots [][]int, x int, y int) bool {

	for _, dot := range dots {
		if dot[0] == x && dot[1] == y {
			// log.Println("There is already a dot at ", x, y)
			return true
		}
	}

	return false
}

/*
 * Apply a command (fold) to the dots. This will transpose dots to a value, and possibly modify the x / y coordinates of the dot.
 * If the dot is overlaying with another dot, then it should NOT be added.
 * Only the TOP and LEFT items are maintained upon folding. Dots should not appear on a fold line.
 */
func fold(dots [][]int, c command) [][]int {
	var idx int

	// pretty.Println(dots, len(dots))

	for {

		// pretty.Println("Iteration", idx)

		if idx >= len(dots) {
			break
		}

		dot := dots[idx]
		dotVal := 0
		otherVal := 0

		// Fold along y=7 means y values will change.
		if c.foldAlong == "x" {
			dotVal = dot[0]   // x val
			otherVal = dot[1] // y val
		} else {
			otherVal = dot[0] // x val
			dotVal = dot[1]   // y val
		}

		// log.Println("Comparing", c.offset, "to", dotVal, "on this dot", dot)

		if c.offset < dotVal {
			// log.Println(dot, "should flip over offset", c.offset)
			// log.Println(dot, "has values", "dot value:", dotVal, "other value:", otherVal)

			newVal := c.offset - (dotVal - c.offset)

			if newVal < 0 {
				// pretty.Println("dot", dot, "folded into inaccessible area. command is", c)
				break
			}

			check := false

			if c.foldAlong == "x" {
				check = dotExists(dots, newVal, otherVal)
			} else {
				check = dotExists(dots, otherVal, newVal)
			}

			// if c.foldAlong == "x" {
			// log.Println("Proposed new position: (x,y)", newVal, otherVal, "Check is", check)
			// } else {
			// log.Println("Proposed new position (x,y):", otherVal, newVal, "Check is", check)
			// }

			if !check {
				if c.foldAlong == "x" {
					// log.Println("Flipping over x foldAlong", dots[idx], newVal)
					// log.Println("dot", dots[idx], "x value set to", newVal)
					// Flip over x foldAlong, new x val
					dots[idx][0] = newVal
				} else {
					// log.Println("Flipping over y foldAlong", dots[idx], newVal)
					// log.Println("dot", dots[idx], "y value set to", newVal)
					// Flip over y foldAlong, new y val
					dots[idx][1] = newVal
				}
			} else {
				// log.Println("Size before removing...", len(dots))
				dots = remove(dots, idx)
				idx-- // We removed this dot, reduce our index to compensate
				// log.Println("Size after removing...", len(dots))
			}

			// log.Println()
		}

		idx++
	}

	return dots
}

func renderDots(dots [][]int) {
	maxX := 0
	maxY := 0

	for _, dot := range dots {
		if dot[0] > maxX {
			maxX = dot[0]
		}

		if dot[1] > maxY {
			maxY = dot[1]
		}
	}

	idxX := 0
	idxY := 0
	for {
		if idxY > maxY {
			break
		}

		idxX = 0
		str := ""
		for {
			if idxX > maxX {
				break
			}

			if dotExists(dots, idxX, idxY) {
				str = str + "#"
			} else {
				str = str + "."
			}

			idxX++
		}

		log.Println(str)

		idxY++
	}
}

func part1(commands []command, dots [][]int) {
	fc := commands[0]

	dots = fold(dots, fc)
	log.Println("Count of dots after 1 command:", len(dots))
}

func part2(commands []command, dots [][]int) {
	for _, c := range commands {
		dots = fold(dots, c)
	}

	renderDots(dots)
}

func main() {
	var dots [][]int
	var commands []command
	inputArr, _ := utils.ReadTextFile("day13/input")

	for _, val := range inputArr {

		if val == "" {
			continue
		}

		if strings.Contains(val, ",") {
			// Dot
			parts := strings.Split(val, ",")
			x, _ := strconv.Atoi(parts[0])
			y, _ := strconv.Atoi(parts[1])
			dots = append(dots, []int{x, y})
		} else {
			// Command
			parts := strings.Split(val, " ")
			parts = strings.Split(parts[2], "=")
			offset, _ := strconv.Atoi(parts[1])

			commands = append(commands, command{foldAlong: parts[0], offset: offset})
		}
	}

	part1(commands, dots)
	part2(commands, dots)
}
