package main

import (
	"fmt"
)

// this should not change
func priority(char rune) int {
	ascii := int(char)
	if ascii >= 97 && ascii <= 122 {
		return ascii - 96
	} else if ascii >= 65 && ascii <= 90 {
		return ascii - 38
	}
	return -1
}

// this was made more flexible to return a rune array
// so it can be called multiple times
func union(a, b []rune) []rune {
	hash := make(map[rune]rune)
	ret := []rune{}

	for _, each := range a {
		hash[each] = each
	}

	for _, each := range b {
		if val, ok := hash[each]; ok {
			ret = append(ret, val)
		}
	}

	return ret
}

// no more splitting, we just want to find common items
// among a group of rucksacks
func findCommonItemsRuckSack(groupRucksacks []string) []rune {

	retVal := []rune(groupRucksacks[0])

	for _, eachRucksack := range groupRucksacks {
		retVal = union([]rune(eachRucksack), retVal)
	}

	return retVal
}

func main() {
	fileContent := readFileReturnArray("./inputs/3.input")

	prioritySum := 0
	group := []string{}
	for _, eachLine := range fileContent {
		group = append(group, eachLine)

		if len(group) == 3 {
			commonRunes := findCommonItemsRuckSack(group)

			prioritySum += priority(commonRunes[0])
			group = []string{}
		}
	}

	fmt.Println("total sum of priorities: ", prioritySum)
}
