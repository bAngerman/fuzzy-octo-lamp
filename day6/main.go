package main

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func getAges(items []int) map[int]int {
	m := make(map[int]int)
	allAges := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	for age := range allAges {
		m[age] = 0
	}

	for _, itemAge := range items {
		m[itemAge] = m[itemAge] + 1
	}

	return m
}

func countLanternfish(items []int, maxDay int) int {
	ages := getAges(items)
	allAges := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	var day, count int

	// pretty.Println("First Age", ages)

	// While loop.
	for {

		if day >= maxDay {
			break
		}

		newFishCount := ages[0]
		ages[0] = 0

		// Go through allAges 8->0 and shift them.
		for _, age := range allAges {

			if age == 0 {
				continue
			}

			newCount := ages[age-1] + ages[age]
			ages[age] = 0
			ages[age-1] = newCount
		}

		// Add new fish count as 8, and add items to ages[6]
		ages[8] = newFishCount
		ages[6] = ages[6] + newFishCount

		// pretty.Println("Ages", ages)

		// Check for any 0 counts in our all ages group, add that many as ages[8], then set those to 6
		day++
	}

	for _, age := range allAges {
		count = count + ages[age]
	}

	return count
}

func part1(items []int) {
	log.Println("Part one final count:", countLanternfish(items, 80))
}

func part2(items []int) {
	log.Println("Part two final count:", countLanternfish(items, 256))
}

func main() {
	fp := "day6/input"
	absPath, _ := filepath.Abs("../" + fp)
	f, err := os.Open(absPath)

	if err != nil {
		log.Panic("Error reading input!: ", err)
	}

	defer f.Close()

	var items []int
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	input := strings.Split(scanner.Text(), ",")

	for _, item := range input {
		val, _ := strconv.Atoi(item)
		items = append(items, val)
	}

	// part1(items)
	part2(items)
}
