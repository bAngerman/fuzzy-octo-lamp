package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/kr/pretty"
)

func part1(items []float64) {
	var horizontalPos, max float64

	for _, item := range items {
		if item > max {
			max = item
		}
	}

	distanceMap := make(map[float64]float64)
	itemCount := float64(len(items))

	for {
		if horizontalPos >= max {
			break
		}

		totalDistance := float64(0)

		for _, item := range items {
			// log.Println("Adding to total distance ", math.Abs(item-horizontalPos))
			totalDistance = totalDistance + math.Abs(item-horizontalPos)
		}

		avgDistance := totalDistance / itemCount

		distanceMap[horizontalPos] = avgDistance

		horizontalPos++
	}

	pretty.Println(distanceMap)
	var bestDistance, worstDistance float64

	// (gross) set to high value so loop finds correct value.
	bestDistance = 1000

	for _, avgDistance := range distanceMap {
		if avgDistance < bestDistance {
			bestDistance = avgDistance
		}

		if avgDistance > worstDistance {
			worstDistance = avgDistance
		}
	}

	var bestPos, worstPos float64

	for idx, item := range distanceMap {
		log.Println(item, bestDistance)
		if item == bestDistance {
			bestPos = idx
		}
		if item == worstDistance {
			worstPos = idx
		}
	}

	log.Println("Best horizontal pos is", bestPos, "with an avg distance of", bestDistance)
	log.Println("Worst horizontal pos is", worstPos, "with an avg distance of", worstDistance)

	var travelledDistance float64
	// Using the best distance, count the total distance required to travel.
	for _, item := range items {
		// log.Println("Adding", math.Abs(item-bestPos))
		travelledDistance = travelledDistance + math.Abs(item-bestPos)
	}

	log.Println("Best horizontal pos,", bestPos, ", has a total distance of", travelledDistance)
}

// func factorial(n float64) float64 {
// 	var factVal, i float64

// 	if n < 0 {
// 		fmt.Print("Factorial of negative number doesn't exist.")
// 	} else {
// 		factVal = 1
// 		for i = 1; i <= n; i++ {
// 			factVal = factVal * i
// 		}
// 	}
// 	log.Println("fact val", factVal)
// 	return factVal
// }

func addAddlFuel(carry, distance float64) float64 {
	var distanceAddIdx float64

	distanceAddIdx = 1
	for {
		if distanceAddIdx > distance-1 {
			break
		}

		carry = carry + distanceAddIdx

		distanceAddIdx++
	}

	carry = carry + distance

	return carry
}

func part2(items []float64) {
	var horizontalPos, max float64

	for _, item := range items {
		if item > max {
			max = item
		}
	}

	distanceMap := make(map[float64]float64)
	itemCount := float64(len(items))

	for {
		if horizontalPos >= max {
			break
		}

		totalDistance := float64(0)

		for _, item := range items {
			thisDistance := math.Abs(item - horizontalPos)
			// totalDistance = totalDistance + thisDistance

			totalDistance = addAddlFuel(totalDistance, thisDistance)
			// if horizontalPos == 5 {
			// 	log.Println(thisDistance, horizontalPos)
			// 	log.Println("Total Distance", totalDistance)
			// }
		}

		avgDistance := totalDistance / itemCount

		distanceMap[horizontalPos] = avgDistance

		horizontalPos++
	}

	var bestDistance, worstDistance float64

	// (gross) set to high value so loop finds correct value.
	bestDistance = 100000

	for _, avgDistance := range distanceMap {
		if avgDistance < bestDistance {
			bestDistance = avgDistance
		}

		if avgDistance > worstDistance {
			worstDistance = avgDistance
		}
	}

	var bestPos, worstPos float64

	for idx, item := range distanceMap {
		// log.Println(item, bestDistance)
		if item == bestDistance {
			bestPos = idx
		}
		if item == worstDistance {
			worstPos = idx
		}
	}

	fmt.Printf("Best horizontal pos is %f with an avg distance of %f", bestPos, bestDistance)
	fmt.Printf("Worst horizontal pos is %f with an avg distance of %f", worstPos, worstDistance)

	var travelledDistance float64
	// Using the best distance, count the total distance required to travel.
	for _, item := range items {

		thisDistance := math.Abs(item - bestPos)
		// travelledDistance = travelledDistance + thisDistance
		// Add on the extra distance for each distance unit.
		travelledDistance = addAddlFuel(travelledDistance, thisDistance)
	}

	fmt.Printf("Best horizontal pos %f has a total distance of %f", bestPos, travelledDistance)
}

func main() {
	fp := "day7/input"
	absPath, _ := filepath.Abs("../" + fp)
	f, err := os.Open(absPath)

	if err != nil {
		log.Panic("Error reading input!: ", err)
	}

	defer f.Close()

	var items []float64
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	input := strings.Split(scanner.Text(), ",")

	for _, item := range input {
		val, _ := strconv.Atoi(item)
		items = append(items, float64(val))
	}

	// part1(items)
	part2(items)
}
