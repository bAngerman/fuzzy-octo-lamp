package main

import (
	"strconv"
	"strings"

	"github.com/bAngerman/adventofcode/utils"
)

type route struct {
	points  []point
	current point
	risk    int
}

type point struct {
	x int
	y int
}

func traverse(routes []route, cave [][]int) []route {
	// maxX := len(cave[0])
	// maxY := len(cave)

}

func part1(cave [][]int) {
	/*
		We start at 0,0. There is only two positions to travel to x+1 (1,0), y+1 (0,1)
		Each position then has two options to traverse to. (left, right)
	*/

	routes := []route{
		route{
			points: []point{
				point{
					x: 0, y: 0,
				},
			},
			current: point{x: 0, y: 0},
			risk:    0,
		},
	}

	routes = traverse(routes, cave)
}

func part2(cave [][]int) {

}

// func renderPath(cave [][]int, routes []route) {
// 	render := ""
// 	for colIdx, col := range cave {
// 		curRowString := ""
// 		for rowIdx, item := range col {
// 			found := false
// 			for _, point := range routes.points {
// 				if point.y == colIdx && point.x == rowIdx {
// 					found = true
// 				}
// 			}
// 			if found {
// 				curRowString = curRowString + "+"
// 			} else {
// 				curRowString = curRowString + fmt.Sprint(item)
// 			}
// 			curRowString = curRowString + "\t"
// 		}
// 		render = render + "\n" + curRowString
// 	}
// 	pretty.Println(render)
// }

func main() {
	var cave [][]int
	inputArr, _ := utils.ReadTextFile("day15/input2")

	for _, rowStr := range inputArr {
		rowSplit := strings.Split(rowStr, "")
		var row []int
		for _, val := range rowSplit {
			intVal, _ := strconv.Atoi(val)

			row = append(row, intVal)
		}
		cave = append(cave, row)
	}

	// path := []path{
	// 	path{x: 0, y: 0},
	// 	path{x: 0, y: 1},
	// 	path{x: 0, y: 2},
	// }

	// renderPath(cave, path)

	part1(cave)
	// part2(inputArr)
}
