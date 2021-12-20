package main

import (
	"log"

	"github.com/kr/pretty"
)

type point struct {
	x  int
	y  int
	xV int
	yV int
}

func abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}

// target area: x=56..76, y=-162..-134
// test target area: x=20..30, y=-10..-5
func part1() {

	// minX := 20
	// maxX := 30
	// minY := -5
	// maxY := -10

	minX := 56
	maxX := 76
	minY := -134
	maxY := -162

	log.Println("Target area is", "x", minX, maxX, "y", minY, maxY)

	var xVelocities, yVelocities []int

	for i := 1; i <= maxX; i++ {
		xVelocities = append(xVelocities, i)
	}

	for i := maxY * 10; i < abs(minY)*10; i++ {
		yVelocities = append(yVelocities, i)
	}

	// pretty.Println(xVelocities, yVelocities)

	var hits []point

	// xVelocities = []int{6}
	// yVelocities = []int{3}

	for _, x := range xVelocities {
		for _, y := range yVelocities {

			var xPos, yPos int

			xV := x
			yV := y

			for {

				// If we are to the right of the target, iterate.
				// If we are below the target, iterate.
				// Or if we are not moving, and not directly over the zone.
				if xPos > maxX || yPos < maxY || xPos < minX && xV == 0 {
					// log.Println("First", xPos > maxX)
					// log.Println("Second", yPos < maxY)
					// log.Println("Third", xPos < minX && xV == 0)

					// log.Println("x", xPos, "y", yPos, "xV", xV, "yV", yV)
					// log.Println(minX, maxX, minY, maxY)
					// log.Println()

					break
				}

				/**
				Calculate the new position as a result of the velocity.
				The probe's x position increases by its x velocity.
				The probe's y position increases by its y velocity.

				THEN check if it is within the target zone.
				*/

				// log.Println(xPos >= minX)
				// log.Println(xPos <= maxX)

				// log.Println(yPos >= minY)
				// log.Println(yPos <= maxY)

				// log.Println()

				// If x pos within minX, maxX and y pos within minY, maxY
				if xPos >= minX && xPos <= maxX && yPos <= minY && yPos >= maxY {
					// log.Println("HIT. Position is", xPos, yPos)
					// log.Println("velocities are", xV, yV)
					// log.Println()
					hits = append(hits, point{x: xPos, y: yPos, xV: x, yV: y})
				}

				xPos = xPos + xV
				yPos = yPos + yV

				/**
				Apply the acceleration to the velocity for the next iteration of time.
				Due to drag, the probe's x velocity changes by 1 toward the value 0; that is,
				it decreases by 1 if it is greater than 0, increases by 1 if it is less than
				0,or does not change if it is already 0.
				Due to gravity, the probe's y velocity decreases by 1.
				*/

				if xV > 0 {
					// Decrement by 1 due to drag.
					xV = xV - 1
				}

				yV = yV - 1 // Apply gravity
			}
		}
	}

	// These are all the acceleration sets that hit the area.
	// Now run the kinematic model on these guys to evaluate max height
	maxHeight := 0
	maxHit := point{}

	// pretty.Println(hits)

	for _, hit := range hits {
		var xPos, yPos int

		xV := hit.xV
		yV := hit.yV

		for {
			if xPos > maxX || yPos < maxY || xPos < minX && xV == 0 {
				break
			}

			xPos = xPos + xV
			yPos = yPos + yV

			// log.Println(maxHeight, yPos)

			if maxHeight < yPos {
				maxHeight = yPos
				maxHit = hit
			}

			if xV > 0 {
				xV = xV - 1
			}

			yV = yV - 1
		}
	}

	pretty.Println("Best hit was", maxHit, "with a height of", maxHeight)
}

// target area: x=56..76, y=-162..-134
// test target area: x=20..30, y=-10..-5
func part2() {
	// minX := 20
	// maxX := 30
	// minY := -5
	// maxY := -10

	minX := 56
	maxX := 76
	minY := -134
	maxY := -162

	// log.Println("Target area is", "x", minX, maxX, "y", minY, maxY)

	var xVelocities, yVelocities []int

	for i := 1; i <= maxX; i++ {
		xVelocities = append(xVelocities, i)
	}

	for i := maxY * 10; i < abs(minY)*10; i++ {
		yVelocities = append(yVelocities, i)
	}

	// pretty.Println(xVelocities, yVelocities)

	var hits []point

	// xVelocities = []int{6}
	// yVelocities = []int{3}

	for _, x := range xVelocities {
		for _, y := range yVelocities {

			var xPos, yPos int

			xV := x
			yV := y

			for {

				// If we are to the right of the target, iterate.
				// If we are below the target, iterate.
				// Or if we are not moving, and not directly over the zone.
				if xPos > maxX || yPos < maxY || xPos < minX && xV == 0 {
					// log.Println("First", xPos > maxX)
					// log.Println("Second", yPos < maxY)
					// log.Println("Third", xPos < minX && xV == 0)

					// log.Println("x", xPos, "y", yPos, "xV", xV, "yV", yV)
					// log.Println(minX, maxX, minY, maxY)
					// log.Println()

					break
				}

				/**
				Calculate the new position as a result of the velocity.
				The probe's x position increases by its x velocity.
				The probe's y position increases by its y velocity.

				THEN check if it is within the target zone.
				*/

				// log.Println(xPos >= minX)
				// log.Println(xPos <= maxX)

				// log.Println(yPos >= minY)
				// log.Println(yPos <= maxY)

				// log.Println()

				// If x pos within minX, maxX and y pos within minY, maxY
				if xPos >= minX && xPos <= maxX && yPos <= minY && yPos >= maxY {
					// log.Println("HIT. Position is", xPos, yPos)
					// log.Println("velocities are", xV, yV)
					// log.Println()
					hits = append(hits, point{x: xPos, y: yPos, xV: x, yV: y})
				}

				xPos = xPos + xV
				yPos = yPos + yV

				/**
				Apply the acceleration to the velocity for the next iteration of time.
				Due to drag, the probe's x velocity changes by 1 toward the value 0; that is,
				it decreases by 1 if it is greater than 0, increases by 1 if it is less than
				0,or does not change if it is already 0.
				Due to gravity, the probe's y velocity decreases by 1.
				*/

				if xV > 0 {
					// Decrement by 1 due to drag.
					xV = xV - 1
				}

				yV = yV - 1 // Apply gravity
			}
		}
	}

	// pretty.Println(hits)

	distinct := []point{}

	for _, hit := range hits {

		found := false

		for _, item := range distinct {
			if item.xV == hit.xV && item.yV == hit.yV {
				found = true
			}
		}

		if !found {
			distinct = append(distinct, point{xV: hit.xV, yV: hit.yV})
		}
	}

	// pretty.Println(distinct)
	pretty.Println("Distinct velocity count", len(distinct))
}

func main() {
	// inputArr, _ := utils.ReadNumberFile("day17/input2")

	// part1()
	part2()
}
