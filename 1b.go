package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
)

func main() {

	fileContent := readFileReturnArray("./inputs/1a.input")

	runningCalorieSum := 0
	var largestCalorieArr []int

	for _, line := range fileContent {

		if line == "" {
			largestCalorieArr = append(largestCalorieArr, runningCalorieSum)
			runningCalorieSum = 0
		} else {
			currentCalorie, err := strconv.Atoi(line)

			// if there was an issue converting stirng to int
			if err != nil {
				log.Fatal(err)
			}

			runningCalorieSum += currentCalorie
		}
	}

	// this was annoying to figure out (reverse sort)
	sort.Sort(sort.Reverse(sort.IntSlice(largestCalorieArr)))

	topThreeCalorieSum := largestCalorieArr[0] + largestCalorieArr[1] + largestCalorieArr[2]
	fmt.Println("sum of the top three:", topThreeCalorieSum)
}
