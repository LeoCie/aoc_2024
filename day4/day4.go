package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println("star1", star1())
	fmt.Println("star2", star2())
}

func star1() int {
	search := "XMAS"
	reverseWordSearch := Reverse(wordSearch)
	forward := forwardSearch(wordSearch, search)
	backward := forwardSearch(reverseWordSearch, search)
	up := searchUpwards(wordSearch, search)
	down := searchUpwards(reverseWordSearch, search)
	return forward + backward + up + down;
}

func star2() int {
	numFound := 0
	rows := strings.Split(wordSearch, "\n")
	var letterArray [][]rune
	for _,row := range rows {
		letterArray = append(letterArray, []rune(row))
	}
	for i,row := range letterArray {
		if (i == 0 || i == len(letterArray) - 1) {
			continue
		}
		for j,letter := range row {
			if letter != 'A' || j == 0 || j == len(row) - 1 {
				continue
			}
			foundMas := 0
			for _,dir := range [][]int{{1,1}, {-1,1}} {
				if letterArray[i + dir[0]][j + dir[1]] == 'M' && letterArray[i + dir[0]*-1][j + dir[1]*-1] == 'S' {
					foundMas++
				}
				if letterArray[i + dir[0]][j + dir[1]] == 'S' && letterArray[i + dir[0]*-1][j + dir[1]*-1] == 'M'  {
					foundMas++
				}
			}
			if (foundMas == 2) {
				numFound++
			}
		}
	}
	return numFound
}

func searchUpwards(wordSearch string, search string) int {
	numFound := 0
	rows := strings.Split(wordSearch, "\n")
	var letterArray [][]rune
	for _,row := range rows {
		letterArray = append(letterArray, []rune(row))
	}
	searchList := []rune(search)
	for i,row := range letterArray {
		if (i < len(search) - 1) {
			continue
		}
		for j,letter := range row {
			if letter != rune(searchList[0]) {
				continue
			}
			if (foundInDir(letterArray, searchList, i, j, -1, 0)) {
				numFound++
			}
		}
		for j,letter := range row {
			if letter != rune(searchList[0]) || j < len(search) - 1 {
				continue
			}
			if (foundInDir(letterArray, searchList, i, j, -1, -1)) {
				numFound++
			}
		}
		for j,letter := range row {
			if letter != rune(searchList[0]) || j > len(row) - len(search) {
				continue
			}
			if (foundInDir(letterArray, searchList, i, j, -1, +1)) {
				numFound++
			}
		}
	}
	
	return numFound
}

func foundInDir(letterArray [][]rune, searchList []rune, rowIndex int, colIndex int, rowDirection int, colDirection int) (bool) {
	for _,s := range searchList[1:] {
		colIndex += colDirection
		rowIndex += rowDirection
		if letterArray[rowIndex][colIndex] != s {
			return false
		}
	}
	return true
}

func forwardSearch(wordSearch string, search string) int {
	a := regexp.MustCompile(search)
	return len(a.FindAllStringIndex(wordSearch, -1))
}

func Reverse(s string) (result string) {
	for _,v := range s {
	  result = string(v) + result
	}
	return 
}