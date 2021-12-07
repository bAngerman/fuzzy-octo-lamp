import (
	"log"
	"strings"
)

func getCounts(inputArr []string) [][2]int {
	inputArrCount := len(inputArr)
	binaryLength := 0

	if len(inputArr) > 0 {
		firstBin := inputArr[0]
		binaryLength = len(firstBin)
	}

	counts := make([][2]int, binaryLength)

	posCountArr := make([]int, binaryLength) // Holds counts of "1" for each index in our input.

	for _, str := range inputArr {
		s := strings.Split(str, "")

		for idx, char := range s {
			if char == "1" {
				posCountArr[idx] = posCountArr[idx] + 1
			}
		}
	}

	for _, posCount := range posCountArr {
		counts = append(counts, [2]int{inputArrCount - posCount, posCount})
	}

	return counts
}

func calculateGamma(inputArr []string) string {
	var gamma = ""

	posCountArr := getCounts(inputArr)

	log.Println(inputArr)

	for idx, posCount := range posCountArr {
		// Gamma prioritizes most "1" bits.

		log.Println(idx, posCount)

		// log.Println(posCount)
		// if posCount[1] > (inputArrCount / 2) {
		// 	gamma = gamma + "1"
		// } else {
		// 	gamma = gamma + "0"
		// }
	}

	return gamma
}

func calculateEpsilon(inputArr []string) string {
	var epsilon = ""

	gamma := calculateGamma(inputArr)
	gammaSplit := strings.Split(gamma, "")

	// Flip values and return
	for _, val := range gammaSplit {
		if val == "1" {
			epsilon = epsilon + "0"
		} else {
			epsilon = epsilon + "1"
		}
	}

	return epsilon
}

func filterItems(inputArr []string, bitIndex int, bitValue string) []string {
	filtered := *new([]string)

	for _, str := range inputArr {
		chars := strings.Split(str, "")
		if chars[bitIndex] == bitValue {
			filtered = append(filtered, str)
		}
	}

	return filtered
}

func getRating(inputArr []string, calcFn func([]string) string) string {
	var res = ""

	binaryLength := 0

	if len(inputArr) > 0 {
		firstBin := inputArr[0]
		binaryLength = len(firstBin)
	}

	idx := 0
	for idx < binaryLength {
		bitPattern := calcFn(inputArr)
		bitPatternSplit := strings.Split(bitPattern, "")

		log.Println(bitPatternSplit)

		val := bitPatternSplit[idx]

		inputArr = filterItems(inputArr, idx, val)
		log.Println("Input after filter", inputArr)

		if len(inputArr) == 1 {
			res = inputArr[0]
			break
		}

		idx++
	}

	return res
}

func calculateLifeSupportRating(inputArr []string) int64 {

	oxygenGeneratorRating := getRating(inputArr, calculateGamma)

	log.Println(oxygenGeneratorRating)
	CO2ScrubberRating := getRating(inputArr, calculateEpsilon)

	oxygenRating := convertToBin(oxygenGeneratorRating)
	CO2Rating := convertToBin(CO2ScrubberRating)

	log.Println("Oxygen Rating", oxygenRating)
	log.Println("CO2 Rating", CO2Rating)

	lifeSupportRating := oxygenRating * CO2Rating

	return lifeSupportRating
}