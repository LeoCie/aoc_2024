package main

import (
	"fmt"
	"math"
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
	nodeArray := getNodesFromInput(nodeLocations)
	mapOfLocations := make(map[rune]([]Location))
	possibleNodeValues := []rune{}
	for i,row := range nodeArray {
		for j,loc := range row {
			if loc == '.' {
				continue
			}
			mapOfLocations[loc] = append(mapOfLocations[loc], Location{i, j})
			if !slices.Contains(possibleNodeValues, loc) {
				possibleNodeValues = append(possibleNodeValues, loc)
			}
		}
	}
	antiNodeLocations := []Location{}
	for _,nodeValue := range possibleNodeValues {
		allLocations := mapOfLocations[nodeValue]
		for _,loc := range allLocations {
			for _,loc2 := range allLocations {
				if (loc.i == loc2.i && loc.j == loc2.j) {
					continue
				}
				iDiff := int(math.Abs(float64(loc2.i - loc.i)))
				jDiff := int(math.Abs(float64(loc2.j - loc.j)))
				var antiNode1 Location
				var antiNode2 Location
				if (loc2.i > loc.i) {
					antiNode1.i = loc.i - iDiff
					antiNode2.i = loc2.i + iDiff
				} else {
					antiNode1.i = loc.i + iDiff
					antiNode2.i = loc2.i - iDiff
				}
				if (loc2.j > loc.j) {
					antiNode1.j = loc.j - jDiff
					antiNode2.j = loc2.j + jDiff
				} else {
					antiNode1.j = loc.j + jDiff
					antiNode2.j = loc2.j - jDiff
				}
				if locInBounds(len(nodeArray), len(nodeArray[0]), antiNode1) && !sliceContains(antiNodeLocations, antiNode1) {
					antiNodeLocations = append(antiNodeLocations, antiNode1)
				}
				if locInBounds(len(nodeArray), len(nodeArray[0]), antiNode2) && !sliceContains(antiNodeLocations, antiNode2) {
					antiNodeLocations = append(antiNodeLocations, antiNode2)
				}
			}
		}
	}
	return len(antiNodeLocations)
}

func star2() int {
	nodeArray := getNodesFromInput(nodeLocations)
	mapOfLocations := make(map[rune]([]Location))
	possibleNodeValues := []rune{}
	for i,row := range nodeArray {
		for j,loc := range row {
			if loc == '.' {
				continue
			}
			mapOfLocations[loc] = append(mapOfLocations[loc], Location{i, j})
			if !slices.Contains(possibleNodeValues, loc) {
				possibleNodeValues = append(possibleNodeValues, loc)
			}
		}
	}
	antiNodeLocations := []Location{}
	for _,nodeValue := range possibleNodeValues {
		allLocations := mapOfLocations[nodeValue]
		for _,loc := range allLocations {
			if len(allLocations)>2 && !sliceContains(antiNodeLocations, loc) {
				antiNodeLocations = append(antiNodeLocations, loc)
			}
			for _,loc2 := range allLocations {
				if (loc.i == loc2.i && loc.j == loc2.j) {
					continue
				}
				iDiff := int(math.Abs(float64(loc2.i - loc.i)))
				jDiff := int(math.Abs(float64(loc2.j - loc.j)))
				for multiplier := 1; multiplier < len(nodeArray); multiplier ++ {
					var antiNode1 Location
					var antiNode2 Location
					if (loc2.i > loc.i) {
						antiNode1.i = loc.i - multiplier*iDiff
						antiNode2.i = loc2.i + multiplier*iDiff
					} else {
						antiNode1.i = loc.i + multiplier*iDiff
						antiNode2.i = loc2.i - multiplier*iDiff
					}
					if (loc2.j > loc.j) {
						antiNode1.j = loc.j - multiplier*jDiff
						antiNode2.j = loc2.j + multiplier*jDiff
					} else {
						antiNode1.j = loc.j + multiplier*jDiff
						antiNode2.j = loc2.j - multiplier*jDiff
					}
					if locInBounds(len(nodeArray), len(nodeArray[0]), antiNode1) && !sliceContains(antiNodeLocations, antiNode1) {
						antiNodeLocations = append(antiNodeLocations, antiNode1)
					}
					if locInBounds(len(nodeArray), len(nodeArray[0]), antiNode2) && !sliceContains(antiNodeLocations, antiNode2) {
						antiNodeLocations = append(antiNodeLocations, antiNode2)
					}
				}
			}
		}
	}
	return len(antiNodeLocations)
}

func sliceContains(s []Location, loc1 Location) bool {
	return slices.ContainsFunc(s, func (loc2 Location) bool {return loc1.i == loc2.i && loc1.j == loc2.j})
}

func locInBounds(xMax int, yMax int, loc Location) bool {
	return loc.i >= 0 && loc.i < xMax && loc.j >= 0 && loc.j < yMax
}

func getNodesFromInput(stringInput string) ([][]rune) {
	rows := strings.Split(stringInput, "\n")
	output := [][]rune{}
	for _,row := range rows {
		output = append(output, []rune(row))
	}
	return output
}