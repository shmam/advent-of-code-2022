package main

import (
	"fmt"
	"log"
)

func priority(char rune) int {
	ascii := int(char)
	if ascii >= 97 && ascii <= 122 {
		return ascii - 96
	} else if ascii >= 65 && ascii <= 90 {
		return ascii - 38
	}
	return -1
}

func unionFirstItem(a, b []rune) rune {
	hash := make(map[rune]rune)

	for _, each := range a {
		hash[each] = each
	}

	for _, each := range b {
		if val, ok := hash[each]; ok {
			return val
		}
	}

	log.Fatal("no items found")
	return '-'
}

func findCommonItemRuckSack(rucksack string) rune {
	length := len(rucksack)
	splitIdx := (length / 2)
	comp1 := rucksack[:splitIdx]
	comp2 := rucksack[splitIdx:]

	// ******************************************
	// this is where I wish there was a native SET type
	// and I could just preform a UNION on two SETS :[
	// ******************************************

	// had to implement something custom instead
	firstCommonItem := unionFirstItem([]rune(comp1), []rune(comp2))
	return firstCommonItem
}

func main() {
	fileContent := readFileReturnArray("./inputs/3.input")

	prioritySum := 0
	for _, eachLine := range fileContent {
		commonRune := findCommonItemRuckSack(eachLine)
		prioritySum += priority(commonRune)
	}

	fmt.Println("total sum of priorities: ", prioritySum)
}
