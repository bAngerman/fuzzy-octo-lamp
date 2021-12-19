package main

import (
	"log"
	"math"
	"sort"
	"strings"

	"github.com/bAngerman/adventofcode/utils"
)

type pairInsert struct {
	pattern string
	val     string
}

type insert struct {
	val string
	idx int
}

func applyPolymerTemplate(tmpl []string, pairInserts []pairInsert) []string {
	var inserts []insert

	for _, item := range pairInserts {
		itemPat := strings.Split(item.pattern, "")

		// The commands to run after which tell us where to insert characters.

		// Reset the index of the inner loop.
		idx := 0
		for {
			if idx > len(tmpl) {
				break
			}

			if idx == len(tmpl)-1 {
				// No possibility to match 2 chars.
				break
			}

			if itemPat[0] == tmpl[idx] && itemPat[1] == tmpl[idx+1] {
				inserts = append(inserts, insert{val: item.val, idx: idx + 1})
				// log.Println("Insert found", item.val, "at index", idx+1)
			}

			idx++
		}
	}

	// We need insertIndex because every time we add an element, we need to offset our initial calculation for offset.
	insertCount := 0

	// Make sure we sort the inserts in ascending order so we insert correctly.
	sort.Slice(inserts, func(i, j int) bool {
		return inserts[i].idx < inserts[j].idx
	})

	for _, insert := range inserts {
		// log.Println(insert)
		idx := insert.idx + insertCount

		// log.Println("Inserting at", idx, tmpl)

		tmpl = append(tmpl[:idx+1], tmpl[idx:]...)

		tmpl[idx] = insert.val

		// log.Println("After insert", tmpl)

		insertCount++
	}

	return tmpl
}

// Brute force solution.
func part1(tmpl []string, pairInserts []pairInsert) {

	steps := 10
	for i := 0; i < steps; i++ {
		tmpl = applyPolymerTemplate(tmpl, pairInserts)

		// log.Println("Step", i+1)
		// log.Println(tmpl)
		// log.Println()
	}

	var counts = make(map[string]int)
	for _, tmplItem := range tmpl {

		_, ok := counts[tmplItem]

		if ok {
			counts[tmplItem] = counts[tmplItem] + 1
		} else {
			counts[tmplItem] = 1
		}
	}

	mostCommon := insert{}
	leastCommon := insert{idx: math.MaxInt}

	for letter, count := range counts {
		if mostCommon.idx < count {
			mostCommon = insert{idx: count, val: letter}
		}

		if leastCommon.idx > count {
			leastCommon = insert{idx: count, val: letter}
		}
	}

	log.Println("Counts", counts)
	log.Println("Diff", mostCommon.idx-leastCommon.idx)
}

// Keep track of pairs in a map containing counts of pairs.
// map[CB:1 NC:1 NN:1]
// TODO: There is a bug with the left and right neighbours incrememnting pair values as a result of inserts..
func part2(tmpl []string, pairInserts []pairInsert) {
	pairs := make(map[string]int)

	idx := 0
	for {
		if idx > len(tmpl) {
			break
		}

		if idx == len(tmpl)-1 {
			// No possibility to match 2 chars.
			break
		}

		pairString := tmpl[idx] + tmpl[idx+1]

		pairs[pairString] = 1

		idx++
	}

	log.Println("Pairs before", pairs)

	steps := 10
	for i := 0; i < steps; i++ {
		toInsert := make(map[string]int)

		for _, insert := range pairInserts {

			pairVal, pairExists := pairs[insert.pattern]

			if pairExists {

				if pairVal == 0 {
					// Skip, there are no occurances.
					continue
				}

				// log.Println("This pattern exists", insert.pattern, "should insert", insert.val)

				// Add new patterns with each end of the pattern to the map.
				// Or increment their count if they exist in the map already.
				// Decrement this pairs[insert.pattern] count
				parts := strings.Split(insert.pattern, "")

				firstNewPair := parts[0] + insert.val
				secondNewPair := insert.val + parts[1]

				// log.Println("First", firstNewPair)
				// log.Println("Second", secondNewPair)

				// log.Println("new pairs are", firstNewPair, "second pair is", secondNewPair)

				toInsert[firstNewPair] = pairs[insert.pattern]
				toInsert[secondNewPair] = pairs[insert.pattern]

				_, patternExists := toInsert[insert.pattern]

				if patternExists {
					toInsert[insert.pattern] = toInsert[insert.pattern] - pairs[insert.pattern]
				} else {
					toInsert[insert.pattern] = -1 * pairs[insert.pattern]
				}

				// log.Println("To Insert", toInsert)
			}

		}

		// log.Println("To Insert", toInsert)
		// log.Println("Before insert pairs", pairs)

		for insertVal, count := range toInsert {
			_, exists := pairs[insertVal]
			if exists {
				pairs[insertVal] = pairs[insertVal] + count
				if pairs[insertVal] < 0 {
					pairs[insertVal] = 0
				}
			} else {
				pairs[insertVal] = count

				if pairs[insertVal] < 0 {
					pairs[insertVal] = 0
				}
			}
		}

		log.Println("After step", i+1)
		log.Println("Pairs", pairs)

		log.Println()
		letterCounts := make(map[string]int)

		for val, count := range pairs {
			// log.Println("Value", val, "has count", count)
			letters := strings.Split(val, "")

			for _, letter := range letters {
				_, exists := letterCounts[letter]

				if exists {
					letterCounts[letter] = letterCounts[letter] + count
				} else {
					letterCounts[letter] = count
				}
			}
		}

		// We doubled counted values because pairs shared values.
		for letter, count := range letterCounts {
			letterCounts[letter] = int(math.Ceil(float64(count) / float64(2)))
		}

		log.Println("Letter counts", letterCounts)
	}
}

func main() {
	inputArr, _ := utils.ReadTextFile("day14/input2")

	var tmpl []string
	var inserts []pairInsert

	for idx, val := range inputArr {
		if val == "" {
			continue
		}

		if idx == 0 {
			// Polymer template
			tmpl = strings.Split(val, "")
		} else {
			parts := strings.Split(strings.ReplaceAll(val, " ", ""), "->")
			inserts = append(inserts, pairInsert{pattern: parts[0], val: parts[1]})
		}
	}

	// part1(tmpl, inserts)
	part2(tmpl, inserts)
}
