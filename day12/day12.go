package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println("star1", star1())
	fmt.Println("star2", star2())
}

type Stone struct {
	val   int
	split bool
}

var inPlot = []Location{}

func star1() int {
	farm := getInputAsRuneMatrix(input)
	fence := 0
	for i, farmRow := range farm {
		for j, _ := range farmRow {
			if slices.ContainsFunc(inPlot, func(loc2 Location) bool { return locationsEqual(Location{i, j}, loc2) }) {
				continue
			}
			area, perim := countPerimAndArea(Location{i, j}, farm)
			fence += area * perim
		}
	}
	return fence
}

func countPerimAndArea(loc Location, farm [][]rune) (int, int) {
	area := 1
	perim := 0
	inPlot = append(inPlot, loc)
	for _, dir := range orthDirections {
		if loc.i+dir.i < 0 || loc.i+dir.i >= len(farm) {
			perim++
			continue
		}
		if (loc.j+dir.j < 0 || loc.j+dir.j >= len(farm[0])) {
			perim++
			continue
		}
		newLocation := Location{loc.i + dir.i, loc.j + dir.j}
		if farm[loc.i][loc.j] != farm[newLocation.i][newLocation.j] {
			perim++
			continue
		}
		if slices.ContainsFunc(inPlot, func(loc2 Location) bool { return locationsEqual(newLocation, loc2) }) {
			continue
		}
		newArea, perimNewLoc := countPerimAndArea(newLocation, farm)
		area+=newArea
		perim += perimNewLoc
	}
	return area, perim
}

var plotId = 1
var plotSideMapI = map[int][]int{}
var plotSideMapJ = map[int][]int{}

func star2() int {
	inPlot = []Location{}
	farm := getInputAsRuneMatrix(input)
	fence := 0
	for i, farmRow := range farm {
		for j, _ := range farmRow {
			if slices.ContainsFunc(inPlot, func(loc2 Location) bool { return locationsEqual(Location{i, j}, loc2) }) {
				continue
			}
			area, perim := countPerimAndArea(Location{i, j}, farm)
			fence += area * perim
			plotId++
		}
	}
	return fence
}
