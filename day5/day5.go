package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	fmt.Println("star1", star1())
	fmt.Println("star2", star2())
}

func star1() int {
	orderMap := buildOrderingMap(orderingRules)
	midSum := 0
	for _, manual := range strings.Split(manuals, "\n") {
		manualSlice := strings.Split(manual, ",")
		if (manualIsInOrder(orderMap, manualSlice)) {
			midSum += getIntFromString(manualSlice[len(manualSlice) / 2])
		}
	}
	return midSum
}

func star2() int {
	orderMap := buildOrderingMap(orderingRules)
	midSum := 0
	for _, manual := range strings.Split(manuals, "\n") {
		manualSlice := strings.Split(manual, ",")
		wasFailing := false
		for !manualIsInOrder(orderMap, manualSlice) {
			wasFailing = true
			manualSlice  = fixLeftMostFailure(orderMap, manualSlice)
		}
		if (wasFailing) {
			midSum += getIntFromString(manualSlice[len(manualSlice) / 2])
		}
	}
	return midSum
}

func fixLeftMostFailure(orderMap map[int][]int, manual []string) []string {
	for i, pageNumber := range(manual) {
		val, exists := orderMap[getIntFromString(pageNumber)]
		if (!exists) {
			continue
		}
		for j, pageNumberToCheck := range(manual[:i]) {
			if slices.Contains(val, getIntFromString(pageNumberToCheck)) {
				manual[i] = pageNumberToCheck
				manual[j] = pageNumber
				return manual 
			}
		}
	}
	return manual
}

func manualIsInOrder(orderMap map[int][]int, manual []string) bool {
	for i, pageNumber := range(manual) {
		val, exists := orderMap[getIntFromString(pageNumber)]
		if (!exists) {
			continue
		}
		for _, pageNumberToCheck := range(manual[:i]) {
			if slices.Contains(val, getIntFromString(pageNumberToCheck)) {
				return false
			}
		}
	}
	return true

}

func buildOrderingMap(orders string) map[int][]int {
	orderMap := make(map[int][]int)
	for _,rule := range(strings.Split(orders, "\n")) {
		ruleNumbers := strings.Split(rule, "|")
		_, exists := orderMap[getIntFromString(ruleNumbers[0])]
		if (exists) {
			orderMap[getIntFromString(ruleNumbers[0])] = append(orderMap[getIntFromString(ruleNumbers[0])], getIntFromString(ruleNumbers[1]))
		} else {
			orderMap[getIntFromString(ruleNumbers[0])] = []int{getIntFromString(ruleNumbers[1])}
		}
	}
	return orderMap
}