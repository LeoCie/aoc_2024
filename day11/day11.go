package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("star1", star1())
	fmt.Println("star2", star2())
}

type Stone struct {
	val   int
	split bool
}

type CacheStone struct {
	val    int
	blinks int
}

var stonesMap = make(map[CacheStone]int)

func star1() int {
	stones := []Stone{}
	for _,stringDigit := range strings.Split(input, " ") {
		var stone Stone
		stone.val = getIntFromString(stringDigit)
		stones = append(stones, stone)
	}
	for i:= 0; i< 25; i++ {
		stones = blink(stones)
	}
	return len(stones)
}

func star2() int {
	stones := []Stone{}
	for _,stringDigit := range strings.Split(input, " ") {
		var stone Stone
		stone.val = getIntFromString(stringDigit)
		stones = append(stones, stone)
	}
	
	stoneCount := 0
	for _,stone := range stones {
		stoneCount += stonesAfterBlinks(stone.val, 75)
	}
	
	return stoneCount
}

func stonesAfterBlinks(stone int, blinks int) int {
	if (blinks == 0) {
		return 1;
	}
	cacheKey := CacheStone{stone, blinks}
	if stonesMap[cacheKey] > 0 {
		return stonesMap[cacheKey]
	}
	numberOfStones := 0
	if (stone == 0 ) {
		numberOfStones = stonesAfterBlinks(1, blinks - 1)
	} else if strStone := strconv.Itoa(stone); len(strStone) % 2 == 0 {
		numberOfStones = stonesAfterBlinks(getIntFromString(strStone[:len(strStone)/2]), blinks - 1) + stonesAfterBlinks(getIntFromString(strStone[len(strStone)/2:]), blinks - 1)
	} else {
		numberOfStones = stonesAfterBlinks(stone * 2024, blinks - 1)
	}

	stonesMap[cacheKey] = numberOfStones
	return numberOfStones
}

func blink(stones []Stone) []Stone {
	for i,stone := range stones {
		if stone.val == 0 {
			stones[i].val = 1
			continue
		}
		if strStone := strconv.Itoa(stone.val); len(strStone) % 2 == 0 {
			stones[i].split = true
			continue
		}
		stones[i].val = stone.val * 2024
	}
	newStones := []Stone{}
	for _,stone := range stones {
		if stone.split {
			strStone := strconv.Itoa(stone.val)
			newStones = append(newStones, Stone{getIntFromString(strStone[:len(strStone)/2]), false})
			newStones = append(newStones, Stone{getIntFromString(strStone[len(strStone)/2:]), false})
		} else {
			newStones = append(newStones, stone)
		}
	}
	return newStones
}