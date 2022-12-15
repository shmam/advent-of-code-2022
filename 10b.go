package main

import (
	"strconv"
	"strings"
)

func main() {
	fileContent := readFileReturnArray("./inputs/10.input")

	X := 1
	cycles := []int{}
	for _, eachLine := range fileContent {
		args := strings.Split(eachLine, " ")

		if args[0] == "addx" {
			val, _ := strconv.Atoi(args[1])

			cycles = append(cycles, X, X)
			X += val
		} else if args[0] == "noop" {
			cycles = append(cycles, X)
		}
	}

	width, height := 40, 6
	crt := make([]rune, width*height)

	for crtPosition, currentCycleValue := range cycles {
		adjustedCRTValue := (crtPosition) % 40
		if currentCycleValue == adjustedCRTValue || currentCycleValue+1 == adjustedCRTValue || currentCycleValue-1 == adjustedCRTValue {
			crt[crtPosition] = '#'
		} else {
			crt[crtPosition] = '.'
		}
	}

	for idx, eachRune := range crt {
		if idx%40 == 0 {
			print("\n")
		}
		print(string(eachRune))
	}
	print("\n")
}
