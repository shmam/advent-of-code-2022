package main

import (
	"fmt"
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

	sum := 0
	cyclesToObserve := []int{20, 60, 100, 140, 180, 220}

	for _, i := range cyclesToObserve {
		fmt.Println(i, cycles[i-1])
		sum += i * cycles[i-1]
	}

	fmt.Println(sum)
}
