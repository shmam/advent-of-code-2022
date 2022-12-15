package main

import (
	"fmt"
	"strconv"
	"strings"
)

func makeForest(fileContent *[]string) [][]int {
	n := len(*fileContent)

	forest := make([][]int, n)
	for x, eachLine := range *fileContent {
		row := strings.Split(eachLine, "")
		forest[x] = make([]int, n)

		for y, eachTree := range row {
			i, _ := strconv.Atoi(eachTree)
			forest[x][y] = i
		}
	}
	return forest
}

func calcTreeScore(x, y, n int, forest *[][]int) int {
	left, right, up, down := 0, 0, 0, 0
	for i := x - 1; i >= 0; i-- {
		left += 1
		if (*forest)[i][y] >= (*forest)[x][y] {
			break
		}
	}

	for i := x + 1; i < n; i++ {
		right += 1
		if (*forest)[i][y] >= (*forest)[x][y] {
			break
		}
	}

	for i := y - 1; i >= 0; i-- {
		up += 1
		if (*forest)[x][i] >= (*forest)[x][y] {
			break
		}
	}

	for i := y + 1; i < n; i++ {
		down += 1
		if (*forest)[x][i] >= (*forest)[x][y] {
			break
		}
	}
	return left * right * up * down
}

func main() {

	fileContent := readFileReturnArray("./inputs/8.input")

	n := len(fileContent)
	forest := makeForest(&fileContent)

	bestScore := 0
	for x := 1; x < n-1; x++ {
		for y := 1; y < n-1; y++ {
			currentScore := calcTreeScore(x, y, n, &forest)
			if currentScore > bestScore {
				bestScore = currentScore
			}
		}
	}
	fmt.Println("Best tree score in the forest: ", bestScore)
}
