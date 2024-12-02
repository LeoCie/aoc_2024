package main

import (
	"fmt"
	"math"
	"sort"
	"strings"

	//"slices"
	//"strings"
	"regexp"
)

func main() {
	fmt.Println("star1", star1())
	fmt.Println("star2", star2())
}

func star1() int {
	a := regexp.MustCompile(` +`)
	listRows := strings.Split(lists, "\n")
	var list1 []int
	var list2 []int
	for i := 0; i < len(listRows); i++ {
		values := a.Split(listRows[i], -1)
		list1 = append(list1, getIntFromString(values[0]))
		list2 = append(list2, getIntFromString(values[1]))
	}
	sort.Ints(list1)
	sort.Ints(list2)
	result := 0
	for i := 0; i < len(listRows); i++ {
		result += int(math.Abs(float64(list1[i] - list2[i])))
	}
	return result
}

func star2() int {
	a := regexp.MustCompile(` +`)
	listRows := strings.Split(lists, "\n")
	var list1 []int
	var list2 []int
	for i := 0; i < len(listRows); i++ {
		values := a.Split(listRows[i], -1)
		list1 = append(list1, getIntFromString(values[0]))
		list2 = append(list2, getIntFromString(values[1]))
	}
	sort.Ints(list1)
	sort.Ints(list2)
	result := 0
	scores := make(map[int]int)
	for i := 0; i < len(listRows); i++ {
		if scores[list1[i]] != 0 {
			result += scores[list1[i]]
		} else {
			timesIdd := 0;
			for _,k := range list2 {
				if k == list1[i] {
					timesIdd++
				}
			}
			score := list1[i] * timesIdd
			scores[list1[i]] = score
			result += score
		}
	}
	return result
}