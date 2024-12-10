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

type Location struct {
	i int
	j int
}

func star1() int {
	mapArray := buildMapArray(topMap)
	found := 0
	for i,row := range mapArray {
		for j,val := range row {
			nines := []Location{}
			if val != 0 {
				continue
			}
			currentNumber := 0
			currentLocations := []Location{{i,j}}
			for currentNumber < 9 {
				nextLocations := []Location{}
				for _,loc := range currentLocations {
					nextLocations = append(nextLocations, findNextNumbers(mapArray, currentNumber, loc)...)
				}
				if (len(nextLocations) == 0) {
					break
				}
				currentLocations = nextLocations
				currentNumber++
			}
			
			if (currentNumber == 9) {
				for _,loc := range currentLocations {
					if !slices.ContainsFunc(nines, func(loc2 Location) bool {return loc.i == loc2.i && loc.j == loc2.j}) {
						nines = append(nines, loc)
						found++
					}
				}
			}
		}
	}
	return found
}

func star2() int {
	mapArray := buildMapArray(topMap)
	distinctPath := 0
	for i,row := range mapArray {
		for j,val := range row {
			if val != 0 {
				continue
			}
			currentNumber := 0
			currentLocations := []Location{{i,j}}
			for currentNumber < 9 {
				nextLocations := []Location{}
				for _,loc := range currentLocations {
					nextLocations = append(nextLocations, findNextNumbers(mapArray, currentNumber, loc)...)
				}
				if (len(nextLocations) == 0) {
					break
				}
				currentLocations = nextLocations
				currentNumber++
			}
			
			if (currentNumber == 9) {
				for i := 0; i < len(currentLocations); i++ {
					distinctPath++
				}
			}
		}
	}
	return distinctPath
}

func findNextNumbers(mapArray [][]int, currentNumber int, location Location) []Location {
	nextLocations := []Location{}
	if (location.i - 1 >= 0 && mapArray[location.i - 1][location.j] == currentNumber + 1) {
		nextLocations = append(nextLocations, Location{location.i - 1, location.j})
	}
	if (location.j - 1 >= 0 && mapArray[location.i][location.j - 1] == currentNumber + 1) {
		nextLocations = append(nextLocations, Location{location.i, location.j - 1})
	}
	if (location.i + 1 < len(mapArray) && mapArray[location.i + 1][location.j] == currentNumber + 1) {
		nextLocations = append(nextLocations, Location{location.i + 1, location.j})
	}
	if (location.j + 1 < len(mapArray[0]) && mapArray[location.i][location.j + 1] == currentNumber + 1) {
		nextLocations = append(nextLocations, Location{location.i, location.j + 1})
	}
	return nextLocations
}


func buildMapArray(topMap string) [][]int {
	output := [][]int{}
	for _,row := range strings.Split(topMap, "\n") {
		rowi := []int{}
		for _,val := range row {
			rowi = append(rowi, getIntFromString(string(val)))
		}
		output = append(output, rowi)
	}
	return output
}