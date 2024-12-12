package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println("star1", star1())
	fmt.Println("star2", star2())
}

var inPlot = []Location{}

func star1() int {
	farm := getInputAsRuneMatrix(input)
	fence := 0
	for i, farmRow := range farm {
		for j := range farmRow {
			if slices.ContainsFunc(inPlot, func(loc2 Location) bool { return locationsEqual(Location{i, j}, loc2) }) {
				continue
			}
			area, perims := getAreaAndPerim(Location{i, j}, farm)
			fence += area * len(perims)
		}
	}
	return fence
}

type Edge struct {
	loc Location
	dir Location
}
func getAreaAndPerim(loc Location, farm [][]rune) (int, []Edge) {
	area := 1
	perims := []Edge{}
	inPlot = append(inPlot, loc)
	for _, dir := range orthDirections {
		newLocation := Location{loc.i + dir.i, loc.j + dir.j}
		if newLocation.i < 0 || newLocation.i >= len(farm) || newLocation.j < 0 || newLocation.j >= len(farm[0]) || farm[loc.i][loc.j] != farm[newLocation.i][newLocation.j] {
			perims = append(perims, Edge{loc, dir})
			continue
		}
		if slices.ContainsFunc(inPlot, func(loc2 Location) bool { return locationsEqual(newLocation, loc2) }) {
			continue
		}
		newArea, perimNewLoc := getAreaAndPerim(newLocation, farm)
		area += newArea
		perims = append(perims, perimNewLoc...)
	}
	return area, perims
}

func star2() int {
	inPlot = []Location{}
	farm := getInputAsRuneMatrix(input)
	fence := 0
	for i, farmRow := range farm {
		for j := range farmRow {
			if slices.ContainsFunc(inPlot, func(loc2 Location) bool { return locationsEqual(Location{i, j}, loc2) }) {
				continue
			}
			area, perims := getAreaAndPerim(Location{i, j}, farm)
			fenceForPlot := 0
			for _, dir := range orthDirections {
				edgeInDir := []Edge{}
				for _,p := range perims {
					if p.dir.i == dir.i && p.dir.j == dir.j {
						edgeInDir = append(edgeInDir, p)
						continue
					}
				}
				if (len(edgeInDir) == 0) {
					continue
				}

				var sortFunc func (e1 Edge, e2 Edge) int
				isI := dir.i == 1 || dir.i == -1
				if isI {
					sortFunc = iSortFunc
				} else {
					sortFunc = jSortFunc
				}

				slices.SortFunc(edgeInDir, sortFunc)

				fenceForPlot++
				curEdge := edgeInDir[0]
				for r,p := range edgeInDir {
					if r == 0 {
						continue
					}
					if isI && (p.loc.i == curEdge.loc.i && p.loc.j == curEdge.loc.j + 1) {
						curEdge = p
						continue
					}
					if !isI && (p.loc.j == curEdge.loc.j && p.loc.i == curEdge.loc.i + 1) {
						curEdge = p
						continue
					}
					fenceForPlot++
					curEdge = p
				}
			}
			
			fence += fenceForPlot*area
		}
	}
	return fence
}

func iSortFunc(e1 Edge, e2 Edge) int {
	if (e1.loc.i > e2.loc.i) {return 1}
	if (e2.loc.i > e1.loc.i) {return -1}
	if (e1.loc.j > e2.loc.j) {return 1}
	if (e2.loc.j > e1.loc.j) {return -1}
	return 0
}

func jSortFunc(e1 Edge, e2 Edge) int {
	if (e1.loc.j > e2.loc.j) {return 1}
	if (e2.loc.j > e1.loc.j) {return -1}
	if (e1.loc.i > e2.loc.i) {return 1}
	if (e2.loc.i > e1.loc.i) {return -1}
	return 0
}