package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {

	fileContent := readFileReturnArray("./inputs/1a.input")

	runningCalorieSum := 0
	largestCalorieSum := 0

	for _, line := range fileContent {

		if line == "" {
			if runningCalorieSum > largestCalorieSum {
				largestCalorieSum = runningCalorieSum
			}

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

	fmt.Println("largest calorie sum", largestCalorieSum)
}
