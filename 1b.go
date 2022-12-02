package main

import (
	"fmt"
    "bufio"
    "log"
	"os"
	"strconv"
	"sort"
)

const inputFile = "./inputs/1a.input"

func main() {
	fileContent, err := os.Open(inputFile)

	if err != nil {
		log.Fatal(err)
	}

	defer fileContent.Close() // will close at the end of the file 

	fileScanner := bufio.NewScanner(fileContent)


	runningCalorieSum := 0
	var largestCalorieArr []int

	for fileScanner.Scan() {
		lineText := fileScanner.Text()

		
		if lineText == "" {
			largestCalorieArr = append(largestCalorieArr, runningCalorieSum)
			runningCalorieSum = 0 
		} else {
			currentCalorie, err := strconv.Atoi(lineText)

			// if there was an issue converting stirng to int
			if err != nil {
				log.Fatal(err)
			}

			runningCalorieSum += currentCalorie
		}
	}

	// if there was an error in the scanner, log error
	if err := fileScanner.Err(); err != nil {
        log.Fatal(err)
    }

	// this was annoying to figure out (reverse sort)
	sort.Sort(sort.Reverse(sort.IntSlice(largestCalorieArr)))

	topThreeCalorieSum := largestCalorieArr[0] + largestCalorieArr[1] + largestCalorieArr[2] 
	fmt.Println("sum of the top three:", topThreeCalorieSum)
}