package main

import (
	"log"
	"regexp"
	"strconv"
)

type stack []rune

func (s stack) push(val rune) stack {
	return append(s, val)
}

func (s stack) pop() (stack, rune) {
	l := len(s)

	if l == 0 {
		log.Fatal("stack empty")
	}
	return s[:l-1], s[l-1]
}

func extractInstructionsFromLine(line string) (num, src, dest int) {
	re := regexp.MustCompile("[0-9]+")
	matches := re.FindAllString(line, -1)

	if matches == nil || len(matches) != 3 {
		log.Fatal("no matches found")
	}

	n, _ := strconv.Atoi(matches[0])
	s, _ := strconv.Atoi(matches[1])
	d, _ := strconv.Atoi(matches[2])
	return n, s, d
}

func main() {
	fileContent := readFileReturnArray("./inputs/5.input")

	// [T]     [Q]             [S]
	// [R]     [M]             [L] [V] [G]
	// [D] [V] [V]             [Q] [N] [C]
	// [H] [T] [S] [C]         [V] [D] [Z]
	// [Q] [J] [D] [M]     [Z] [C] [M] [F]
	// [N] [B] [H] [N] [B] [W] [N] [J] [M]
	// [P] [G] [R] [Z] [Z] [C] [Z] [G] [P]
	// [B] [W] [N] [P] [D] [V] [G] [L] [T]
	//  1   2   3   4   5   6   7   8   9
	var ship = map[int]stack{
		1: stack{'B', 'P', 'N', 'Q', 'H', 'D', 'R', 'T'},
		2: stack{'W', 'G', 'B', 'J', 'T', 'V'},
		3: stack{'N', 'R', 'H', 'D', 'S', 'V', 'M', 'Q'},
		4: stack{'P', 'Z', 'N', 'M', 'C'},
		5: stack{'D', 'Z', 'B'},
		6: stack{'V', 'C', 'W', 'Z'},
		7: stack{'G', 'Z', 'N', 'C', 'V', 'Q', 'L', 'S'},
		8: stack{'L', 'G', 'J', 'M', 'D', 'N', 'V'},
		9: stack{'T', 'P', 'M', 'F', 'Z', 'C', 'G'},
	}

	for _, eachLine := range fileContent {
		num, src, dest := extractInstructionsFromLine(eachLine)

		for ; num > 0; num-- {
			temp := rune(0)
			ship[src], temp = ship[src].pop()
			ship[dest] = ship[dest].push(temp)
		}
	}

	bag := [10]string{}
	for k, v := range ship {
		_, topValue := v.pop()
		bag[k] = string(topValue)
	}

	for _, s := range bag {
		print(s)
	}

}
