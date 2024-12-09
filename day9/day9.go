package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println("star1", star1())
	fmt.Println("star2", star2())
}

func star1() int {
	outputDisk := buildOutputDisk(disk)
	for i:= len(outputDisk) - 1; i >= 0; i-- {
		val := outputDisk[i]
		if val != -1 {
			firstEmptySpace := slices.Index(outputDisk, -1)
			if (firstEmptySpace == -1 || firstEmptySpace > i) {
				break
			}
			outputDisk[firstEmptySpace] = outputDisk[i]
			outputDisk[i] = -1
		}
	}
	return sumOutputDisk(outputDisk)
}

func star2() int {
	outputDisk := buildOutputDisk(disk)
	for i := len(outputDisk) - 1; i >= 0; {
		val := outputDisk[i]
		if val != -1 {
			size := 1
			sizeLoop:
			for i-size >=0 {
				d := outputDisk[i-size]
				if d != val {
					break sizeLoop
				}
				size++
			}
			gapIndex := findFirstGapOfSize(outputDisk, size)
			if gapIndex > i - size || gapIndex == -1 {
				i -= size
				continue
			}
			for gi := gapIndex; gi < gapIndex + size; gi++ {
				outputDisk[gi] = val
			}
			startingi := i
			for mi := startingi; mi > startingi-size; mi-- {
				outputDisk[mi] = -1
			}
			i -= size
		} else {
			i--
		}
	}
	return sumOutputDisk(outputDisk)
}

func findFirstGapOfSize(disk []int, size int) int {
	found := 0
	firstIndex := -1
	for i,v := range disk {
		if v == -1 {
			found++
			if (firstIndex == -1) {
				firstIndex = i
			}
			if (found == size) {
				break
			}
		} else {
			firstIndex = -1
			found = 0
		}
	}
	if found == size {
		return firstIndex
	}
	return -1
}

func buildOutputDisk(disk string) []int {
	fileBlockNum := 0
	isFile := true
	outputDisk := []int{}
	for _,block := range disk {
		var toAppend int
		if isFile {
			toAppend = fileBlockNum
			fileBlockNum++
		} else {
			toAppend = -1
		}
		for i:= 0; i < getIntFromString(string(block)); i++ {
			outputDisk = append(outputDisk, toAppend)
		}
		isFile = !isFile
	}
	return outputDisk
}

func sumOutputDisk(disk []int) int {
	sum := 0
	for i:=0;i<len(disk);i++ {
		if (disk[i] == -1) {
			continue
		}
		sum += i*disk[i]
	}
	return sum
}