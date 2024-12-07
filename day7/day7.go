package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("star1", star1())
	fmt.Println("star2", star2())
}

type Equation struct {
	vals []int
	target int
}

func star1() int {
	equations := getEquationsFromInput(possibleEquations)
	summedPossibleEqs := 0
	eqLoop:
	for _,eq := range equations {
		possibleOperatorCombinations := getPossibleOperators(len(eq.vals) - 1, []string{"+", "*"})
		operatorCombinationLoop:
		for _,opCombinations := range possibleOperatorCombinations {
			currentResult := eq.vals[0]
			for i,op := range opCombinations {
				if (op == "+") {
					currentResult += eq.vals[i + 1]
				} else if (op == "*") {
					currentResult *= eq.vals[i + 1]
				}
				if (currentResult > eq.target) {
					continue operatorCombinationLoop
				}
			}
			if (currentResult == eq.target) {
				summedPossibleEqs += eq.target
				continue eqLoop
			}
		}
	}
	return summedPossibleEqs
}

func star2() int {
	equations := getEquationsFromInput(possibleEquations)
	summedPossibleEqs := 0
	eqLoop:
	for _,eq := range equations {
		possibleOperatorCombinations := getPossibleOperators(len(eq.vals) - 1, []string{"+", "*", "|"})
		operatorCombinationLoop:
		for _,opCombinations := range possibleOperatorCombinations {
			currentResult := eq.vals[0]
			for i,op := range opCombinations {
				if (op == "+") {
					currentResult += eq.vals[i + 1]
				} else if (op == "*") {
					currentResult *= eq.vals[i + 1]
				} else if (op == "|") {
					currentResult = getIntFromString(strconv.Itoa(currentResult) + strconv.Itoa(eq.vals[i + 1]))
				}
				if (currentResult > eq.target) {
					continue operatorCombinationLoop
				}
			}
			if (currentResult == eq.target) {
				summedPossibleEqs += eq.target
				continue eqLoop
			}
		}
	}
	return summedPossibleEqs
}

func getPossibleOperators(numberOfOperators int, operators []string) [][]string {
	allCombinations := [][]string{}
	for i := 0; i < int(math.Pow(float64(len(operators)), float64(numberOfOperators))); i++ {
		newCombination := []string{}
		binaryRep := strconv.FormatInt(int64(i), len(operators))
		for (len(binaryRep) < numberOfOperators) {
			binaryRep = "0" + binaryRep
		}
		for _,dig := range binaryRep {
			newCombination = append(newCombination, operators[getIntFromString(string(dig))])
		}
		allCombinations = append(allCombinations, newCombination)
	}
	return allCombinations
}

func getEquationsFromInput(equationString string) ([]Equation) {
	rows := strings.Split(equationString, "\n")
	var equations []Equation
	for _,row := range rows {
		targetAndVals := strings.Split(row, ":")
		var eq Equation
		eq.target = getIntFromString(targetAndVals[0])
		for _, val := range strings.Split(targetAndVals[1], " ") {
			eq.vals = append(eq.vals, getIntFromString(val))
		}
		equations = append(equations, eq)
	}
	return equations
}