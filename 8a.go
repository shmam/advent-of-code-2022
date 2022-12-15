package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func max(a, b, c, d int) int {
	temp := []int{a, b, c, d}
	sort.Slice(temp, func(i, j int) bool {
		return temp[i] > temp[j]
	})
	return temp[0]
}

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

// checks if a certain tree is visible in the forest, using the
// forest as a passed by refrence parameter
func checkIfVisible(x, y, n int, forest *[][]int) int {
	left, right, up, down := 1, 1, 1, 1
	for i := x - 1; i >= 0; i-- {
		if (*forest)[i][y] >= (*forest)[x][y] {
			left = 0
		}
	}

	for i := x + 1; i < n; i++ {
		if (*forest)[i][y] >= (*forest)[x][y] {
			right = 0
		}
	}

	for i := y - 1; i >= 0; i-- {
		if (*forest)[x][i] >= (*forest)[x][y] {
			up = 0
		}
	}

	for i := y + 1; i < n; i++ {
		if (*forest)[x][i] >= (*forest)[x][y] {
			down = 0
		}
	}
	return max(left, right, up, down)
}

func main() {

	fileContent := readFileReturnArray("./inputs/8.input")

	n := len(fileContent)
	forest := makeForest(&fileContent)

	alreadyVisible := 4 * (n - 1)
	sum := alreadyVisible
	for x := 1; x < n-1; x++ {
		for y := 1; y < n-1; y++ {
			sum += checkIfVisible(x, y, n, &forest)
		}
	}
	fmt.Println("Number of trees visible from outside the forest: ", sum)
}
