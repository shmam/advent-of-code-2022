package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Range struct {
	start int
	end   int
}

func isPartialOverlap(r1, r2 Range) bool {
	return (r1.end >= r2.start && r1.start <= r2.end) || (r2.end >= r1.start && r2.start <= r1.end)
}

func parseLineToRanges(line string) (Range, Range) {
	strRanges := strings.Split(line, ",")

	strRange1 := strings.Split(strRanges[0], "-")
	strRange2 := strings.Split(strRanges[1], "-")

	r1Start, _ := strconv.Atoi(strRange1[0])
	r1End, _ := strconv.Atoi(strRange1[1])

	r2Start, _ := strconv.Atoi(strRange2[0])
	r2End, _ := strconv.Atoi(strRange2[1])

	return Range{r1Start, r1End}, Range{r2Start, r2End}
}

func main() {
	fileContent := readFileReturnArray("./inputs/4.input")

	runningSum := 0
	for _, eachLine := range fileContent {
		r1, r2 := parseLineToRanges(eachLine)

		if isPartialOverlap(r1, r2) {
			fmt.Println(r1, r2)
			runningSum += 1
		}
	}

	fmt.Println("Total number of partial overlaps: ", runningSum)
}
