package main

import (
	"log"
	"math"
	"regexp"
	"strconv"
)

type Range struct {
	start, end int
}

func mergeRange(r1, r2 Range) Range {
	if r1.end >= r2.start && r1.start < r2.end {
		return Range{r1.start, r2.end}
	} else if r2.end >= r1.start && r2.start < r1.end {
		return Range{r2.start, r1.end}
	}
	return Range{}
}

func rangesOverlap(r1, r2 Range) bool {
	return (r1.end >= r2.start && r1.start < r2.end) || (r2.end >= r1.start && r2.start < r1.end)
}

func manhattanDistance(x1, y1, x2, y2 int) int {
	return int(math.Abs(float64(x1-x2)) + math.Abs(float64(y1-y2)))
}

func markArea(senX, senY, beacX, beacY int, cave *map[int][]Range) {
	r := manhattanDistance(senX, senY, beacX, beacY)

	// figure out what Xs are on the circumfrence at any given y in range
	for y := senY- r; y <= senY+r; y++ {

		if y == 2000000 {
			aX := senX + (r - int(math.Abs(float64(senY-y))))
			bX := senX - ((-1 * int(math.Abs(float64(senY-y)))) - r)

			newRange := Range{aX, bX}
			if (*cave)[y] != nil {
				(*cave)[y] = append((*cave)[y], newRange)
			} else {
				(*cave)[y] = []Range{Range{aX, bX}}
			}
		}
	}

}

func extractAllWholeNumbersFromLine(line string) []int {
	numbers := make([]int, 0)
	re := regexp.MustCompile(`[-]?\d[\d]*[\.]?[\d{2}]*`)
	extractedStrings := re.FindAllString(line, -1)

	for _, eachString := range extractedStrings {
		val, _ := strconv.Atoi(eachString)
		numbers = append(numbers, val)
	}
	return numbers
}

func main() {
	fileContent := readFileReturnArray("./inputs/15.input")

	cave := make(map[int][]Range)

	for _, eachLine := range fileContent {
		n := extractAllWholeNumbersFromLine(eachLine)

		if len(n) != 4 {
			log.Fatal("error reading input")
		}

		markArea(n[0], n[1], n[2], n[3], &cave)
	}

	k := map[int]int{}
	for _, each := range cave[2000000] {
		for i := each.start; i <= each.end; i++ {
			k[i] = 1
		}
	}

	print(len(k))
}
