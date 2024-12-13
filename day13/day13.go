package main

import (
	"fmt"
	"math"
	"regexp"
	"strings"
	"gonum.org/v1/gonum/mat"
)


func main() {
	fmt.Println("star1", star1())
	fmt.Println("star2", star2())
}

func star1() int {
	machines := parseInputToMachine(input)
	output := 0
	for _,machine := range machines {
		lowestCost := 0
		for i := 0; i <= 100 ; i++ {
			aPresses := i
			xLeft := machine.xPrize - aPresses * machine.aButton.xMove
			yLeft := machine.yPrize - aPresses * machine.aButton.yMove
			bPresses := yLeft / machine.bButton.yMove
			if xLeft % machine.bButton.xMove != 0 || yLeft % machine.bButton.yMove != 0 || yLeft / machine.bButton.yMove != xLeft / machine.bButton.xMove {
				continue
			}
			cost := getCost(aPresses, bPresses)
			if lowestCost == 0 || cost < lowestCost {
				lowestCost = cost
			}
		}
		output += lowestCost
	}
	return output
}

func star2() int {
	machines := parseInputToMachine(input)
	output := 0
	for _,machine := range machines {
		cost := solveEquation(float64(machine.aButton.xMove), float64(machine.bButton.xMove), float64(machine.aButton.yMove), float64(machine.bButton.yMove), float64(machine.xPrize + 10000000000000), float64(machine.yPrize + 10000000000000))
		if (cost == -1) {
			continue
		}
		output += cost
	}
	fmt.Println("!!! Not correct")
	return output
}

func solveEquation(aButtonXMove, bButtonXMove, aButtonYMove, bButtonYMove, xPrize, yPrize float64) int {
	A := mat.NewDense(2, 2, []float64{aButtonXMove, bButtonXMove, aButtonYMove, bButtonYMove})
    b := mat.NewVecDense(2, []float64{xPrize, yPrize})

    var x mat.VecDense
    if err := x.SolveVec(A, b); err != nil {
        fmt.Println("failed, non solvable")
		return -1
    }

	if x.At(0,0) != math.Trunc(x.At(0,0)) || x.At(1,0) != math.Trunc(x.At(1,0)) {
		fmt.Println("failed, non int")
		return -1
	}
	fmt.Println("passed", x)
    return int(x.At(0,0) * 3 + x.At(1,0))
}

func getCost(aPress int, bPress int) int {
	return aPress * 3 + bPress
}

type Button struct {
	xMove, yMove int
}

type Machine struct {
	aButton, bButton Button
	xPrize, yPrize int
}

func parseInputToMachine(input string) []Machine {
	machines := []Machine{}
	for _,machineString := range strings.Split(input, "\n\n") {
		var machine Machine
		for i,row := range strings.Split(machineString, "\n") {
			if i == 0 || i == 1 { // Button A or B
				var button Button
				split := strings.Split(row, "+")
				button.xMove = getIntFromString(split[1][:2])
				button.yMove = getIntFromString(split[2][:2])
				if i == 0 {
					machine.aButton = button
				} else {
					machine.bButton = button
				}
			} else { // Prizes
				g := regexp.MustCompile(`\d+`)
				split := strings.Split(row, "=")
				machine.xPrize = getIntFromString(g.FindString(split[1]))
				machine.yPrize = getIntFromString(g.FindString(split[2]))
			}
		}
		machines = append(machines, machine)
	}
	return machines
}