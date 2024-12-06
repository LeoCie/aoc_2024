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

type Position struct {
	row int
	col int
	dir []int
}

func star1() int {
	mapArray, currentPosition := getMapArrayAndStartingPosition(guardMap)
	travelledLocations := getTravelledLocations(mapArray, currentPosition)
	return len(travelledLocations)
}

func star2() int {
	mapArray, startingPosition := getMapArrayAndStartingPosition(guardMap)
	travelledLocations := getTravelledLocations(mapArray, startingPosition)

	numOfRows := len(mapArray)
	numOfCols := len(mapArray[0])
	loopFound:=0
	for _,location := range(travelledLocations) {
		if (mapArray[location.row][location.col] == '.') {
			mapArray[location.row][location.col] = '#'
		} else {
			continue
		}
		currentPosition := startingPosition
		visitedPositions := []Position{}
		guardMoveLoop:
		for !atEdge(currentPosition, numOfRows, numOfCols) {
			if slices.ContainsFunc(visitedPositions, func(pos Position) bool { return positionsAreEqual(pos, currentPosition, true)}) {
				loopFound++
				break guardMoveLoop
			}
			visitedPositions = append(visitedPositions, currentPosition)
			newPosition := getNewPosition(currentPosition)
			if mapArray[newPosition.row][newPosition.col] == '#' {
				currentPosition.dir = turnLeft(currentPosition.dir)
				continue guardMoveLoop
			}
			currentPosition = newPosition
		}
		mapArray[location.row][location.col] = '.'
	
	}
	return loopFound
}

func turnLeft(dir []int) []int {
	if dir[0] == -1 {
		return []int{0, 1}
	} else if dir[0] == 1 {
		return []int{0, -1}
	} else if dir[1] == -1 {
		return []int{-1, 0}
	} else if dir[1] == 1 {
		return []int{1, 0}
	}
	panic("wrong dir")
}

func atEdge(currentPosition Position, maxRowNum int, maxColNum int) bool {
	return currentPosition.row == 0 || currentPosition.row == maxRowNum - 1 || currentPosition.col == 0 || currentPosition.col == maxColNum - 1
}

func getTravelledLocations(mapArray [][]rune, currentPosition Position) []Position {
	visitedPositions := []Position{}
	numOfRows := len(mapArray)
	numOfCols := len(mapArray[0])
	for !atEdge(currentPosition, numOfRows, numOfCols) {
		newPosition := getNewPosition(currentPosition)
		if mapArray[newPosition.row][newPosition.col] == '#' {
			currentPosition.dir = turnLeft(currentPosition.dir)
			continue
		}
		if !slices.ContainsFunc(visitedPositions,func (pos Position) bool { return positionsAreEqual(pos, currentPosition, false) }) {
			visitedPositions = append(visitedPositions, currentPosition)
		}
		currentPosition = newPosition
	}
	visitedPositions = append(visitedPositions, currentPosition)
	return visitedPositions
}

func positionsAreEqual(pos1 Position, pos2 Position, careAboutDir bool) bool {
	return pos1.col == pos2.col && pos1.row == pos2.row && (!careAboutDir || (pos1.dir[0] == pos2.dir[0] && pos1.dir[1] == pos2.dir[1]))
}

func getMapArrayAndStartingPosition(guardMap string) ([][]rune, Position) {
	rows := strings.Split(guardMap, "\n")
	var mapArray [][]rune
	var currentPosition Position
	for i,row := range rows {
		mapArray = append(mapArray, []rune(row))
		if startIndex := slices.Index([]rune(row), '^'); startIndex > 0 {
			currentPosition.row = i
			currentPosition.col = startIndex
		}
	}
	currentPosition.dir = []int{-1, 0}
	return mapArray, currentPosition
}

func getNewPosition(currentPosition Position) Position {
	var newPosition Position
	newPosition.row = currentPosition.row + currentPosition.dir[0]
	newPosition.col = currentPosition.col + currentPosition.dir[1]
	newPosition.dir = currentPosition.dir
	return newPosition;
}