package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("star1", star1())
	fmt.Println("star2", star2())
}

var mulExp = regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
func star1() int {
	allMuls := mulExp.FindAllString(memory, -1)
	return getResult(allMuls)
}

func star2() int {
	enabledExp := regexp.MustCompile(`(do\(\)|^)[\s|\S]*?(don't\(\)|$)`)
	enabledStretches := enabledExp.FindAllString(memory, -1)
	var allMuls []string
	for _,stretch := range enabledStretches {
		mulsInStretch := mulExp.FindAllString(stretch, -1)
		allMuls = append(allMuls, mulsInStretch...)
	}
	return getResult(allMuls)
}

func getResult(allMuls []string) int {
	result := 0
	for _,mul := range allMuls {
		num := regexp.MustCompile(`\d+`)
		nums := num.FindAllString(mul, -1)
		result += getIntFromString(nums[0])*getIntFromString(nums[1])
	}
	return result
}