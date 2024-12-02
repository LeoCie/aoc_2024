package main

import (
	"fmt"
	"math"
	"slices"
	//"sort"
	"strings"
	//"slices"
	//"strings"
	//"regexp"
)

func main() {
	fmt.Println("star1", star1())
	fmt.Println("star2", star2())
}

func star1() int {
	reports := strings.Split(reports, "\n")
	safeCount := 0
	for _,report := range reports {
		valsInt := getIntReport(report)
		if (isSafe(valsInt)) {
			safeCount += 1;
		}
	}
	return safeCount
}

func star2() int {
	reports := strings.Split(reports, "\n")
	safeCount := 0
	reportLoop:
	for _,report := range reports {
		valsInt := getIntReport(report)
		if (isSafe(valsInt)) {
			safeCount++
			continue;
		}
		for i := 0; i < len(valsInt); i++ {
			newValues := make([]int, len(valsInt))
			copy(newValues, valsInt)
			newValues = slices.Delete(newValues, i, i+1)
			if (isSafe(newValues)) {
				safeCount++;
				continue reportLoop;
			}
		}
	}

	return safeCount
}

func getIntReport(report string) []int {
	vals := strings.Split(report, " ")
	var valsInt []int
	for _,v := range vals {
		valsInt = append(valsInt, getIntFromString(v))
	}
	return valsInt
}

func isSafe(input []int) bool {
	if !slices.IsSorted(input) {
		slices.Reverse(input)
		if !slices.IsSorted(input) {
			return false
		}
	}
	safe := true
	for i := 0; i < len(input) - 1; i++ {
		diff := int(math.Abs(float64(input[i] - input[i+1])))
		if diff > 3 || diff == 0 {
			safe = false
			break
		}
	}
	return safe
}