package main

import (
	"errors"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/bAngerman/adventofcode/utils"
	"github.com/kr/pretty"
)

type inputLine struct {
	sigPattern []string
	outputVal  []string
}

func parseItem(str string) (int, error) {
	var digit int

	if len(str) == 7 {
		return 8, nil
	}

	if len(str) == 4 {
		return 4, nil
	}

	if len(str) == 3 {
		return 7, nil
	}

	if len(str) == 2 {
		return 1, nil
	}

	return digit, errors.New(str + " was an not found")
}

func part1(lines []inputLine) {
	var digitCount int

	for _, item := range lines {
		for _, str := range item.outputVal {
			digit, err := parseItem(str)

			if err == nil {
				pretty.Println("String", str, "cooresponds to digit", digit)

				digitCount = digitCount + 1
			}

			if err != nil {
				log.Println(err)
			}
		}
	}

	log.Println("Count of digits is", digitCount)
}

// difference returns the elements in `a` that aren't in `b`.
func difference(a, b []string) []string {
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}

func sortCompare(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	sort.Strings(a)
	sort.Strings(b)

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func part2(lines []inputLine) {
	var outputTotal int

	for _, item := range lines {

		digitMap := make(map[int][]string)

		for _, str := range item.sigPattern {
			digit, err := parseItem(str)

			if err == nil {
				digitMap[digit] = strings.Split(str, "")
			}

			if err != nil {
				// log.Println(err)
			}
		}

		// Discover "3" pattern
		// These are 3, 2, or 5. "3" will contain all segments from "1"
		for _, str := range item.sigPattern {
			if len(str) == 5 {
				curString := strings.Split(str, "")
				if len(difference(curString, digitMap[1])) == 3 {
					digitMap[3] = curString
				}
			}
		}

		// Can now discover "9"
		// 6 length strings, with a near exact intersection with 3 (1 remaining char)
		for _, str := range item.sigPattern {
			if len(str) == 6 {
				curString := strings.Split(str, "")
				if len(difference(curString, digitMap[3])) == 1 {
					digitMap[9] = curString
				}
			}
		}

		// Now that we know "9", we can discover "6"
		// Only remaining 6 length item that is not "9" and intersects 8 w/ 1 remainder
		for _, str := range item.sigPattern {
			if len(str) == 6 {
				curString := strings.Split(str, "")

				// log.Println("Difference between", curString, ",", digitMap[8], "is", difference(digitMap[8], curString))
				// log.Println("Difference between", curString, ",", digitMap[8], "is", difference(digitMap[8], curString))
				// log.Println("Difference between", curString, ",", digitMap[9], "is", difference(digitMap[9], curString))

				if len(difference(digitMap[1], curString)) == 1 {
					digitMap[6] = curString
				}

				// if ( )
			}
		}

		// Discover "0"
		for _, str := range item.sigPattern {
			if len(str) == 6 {
				curString := strings.Split(str, "")

				// log.Println("Difference between", curString, ",", digitMap[8], "is", difference(digitMap[8], curString))
				// log.Println("Difference between", curString, ",", digitMap[7], "is", difference(digitMap[7], curString))
				// log.Println("Difference between", curString, ",", digitMap[6], "is", difference(digitMap[6], curString))
				// log.Println("Difference between", curString, ",", digitMap[4], "is", difference(digitMap[4], curString))
				// log.Println("Difference between", curString, ",", digitMap[3], "is", difference(digitMap[3], curString))

				if len(difference(digitMap[8], curString)) == 1 && len(difference(digitMap[7], curString)) == 0 && len(difference(digitMap[6], curString)) == 1 && len(difference(digitMap[4], curString)) == 1 {
					digitMap[0] = curString
				}
			}
		}

		// Now discover "5", "2" from our discovered "9"
		for _, str := range item.sigPattern {
			if len(str) == 5 {
				curString := strings.Split(str, "")

				// log.Println("Difference between", curString, ",", digitMap[6], "is", difference(digitMap[6], curString))

				if len(difference(digitMap[9], curString)) == 2 {
					digitMap[2] = curString
				}

				if len(difference(digitMap[6], curString)) == 1 {
					digitMap[5] = curString
				}
			}
		}

		// Now compute the signal value for this segment
		output := ""

		for _, str := range item.outputVal {
			strSplit := strings.Split(str, "")

			for idx, digitMapItem := range digitMap {
				if sortCompare(strSplit, digitMapItem) {
					output = output + strconv.Itoa(idx)
				}
			}
		}
		// log.Println("Output is", output)

		outputInt, _ := strconv.Atoi(output)
		outputTotal = outputTotal + outputInt
	}

	log.Println("Output Total is", outputTotal)
}

func main() {
	inputArr, _ := utils.ReadTextFile("day8/input")
	lines := []inputLine{}

	for _, input := range inputArr {
		parts := strings.Split(input, "|")
		l := inputLine{sigPattern: strings.Fields(parts[0]), outputVal: strings.Fields(parts[1])}
		lines = append(lines, l)
	}

	// pretty.Println(lines)

	part1(lines)
	part2(lines)
}
