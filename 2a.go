package main

import (
	"fmt"
	"strings"
)

var score = [3][3]int{
	{4, 8, 3},
	{1, 5, 9},
	{7, 2, 6},
}

var translation = map[string]int{
	"A": 0,
	"B": 1,
	"C": 2,
	"X": 0,
	"Y": 1,
	"Z": 2,
}

func calculateScore(opponentInput, playerInput string) int {
	return score[translation[opponentInput]][translation[playerInput]]
}

func main() {
	fileContent := readFileReturnArray("./inputs/2.input")

	gameTotal := 0
	for _, eachLine := range fileContent {
		splitInput := strings.Fields(eachLine)

		gameTotal += calculateScore(splitInput[0], splitInput[1])
	}

	fmt.Println("Game Total: ", gameTotal)
}
