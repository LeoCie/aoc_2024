package main
import (
	"strconv"
	"strings"
)

func GCD(a, b int) int {
	for b != 0 {
			t := b
			b = a % b
			a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
			result = LCM(result, integers[i])
	}

	return result
}

func All[T any](ts []T, pred func(T) bool) bool {
    for _, t := range ts {
        if !pred(t) {
            return false
        }
    }
    return true
}

func first(n int, _ error) int {
	return n
}
func getIntFromString(param string) int {
	return first(strconv.Atoi(param))
}

func ReverseString(s string) (result string) {
	for _,v := range s {
	  result = string(v) + result
	}
	return 
}

type Location struct {
	i,j int
}

func locationsEqual(loc1 Location, loc2 Location) bool {
	return loc1.i == loc2.i && loc1.j == loc2.j
}

func getInputAsIntMatrix(input string) [][]int {
	output := [][]int{}
	for _,row := range strings.Split(input, "\n") {
		rowi := []int{}
		for _,val := range row {
			rowi = append(rowi, getIntFromString(string(val)))
		}
		output = append(output, rowi)
	}
	return output
}

func getInputAsRuneMatrix(input string) [][]rune {
	output := [][]rune{}
	for _,row := range strings.Split(input, "\n") {
		output = append(output, []rune(row))
	}
	return output
}

var orthDirections = []Location{{1,0}, {0,1}, {-1,0}, {0,-1}}
var diagonalDirections = []Location{{1,1}, {1,-1}, {-1,1}, {-1,-1}}
var allDirections = append(orthDirections, diagonalDirections...)