package main

import (
	"fmt"
    "bufio"
    "log"
	"os"
	"strconv"
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
	largestCalorieSum := 0 

	for fileScanner.Scan() {
		lineText := fileScanner.Text()

		
		if lineText == "" {
			if runningCalorieSum > largestCalorieSum {
				largestCalorieSum = runningCalorieSum
			}

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

	fmt.Println("largest calorie sum", largestCalorieSum)
}