package main

import (
	"fmt"
	// "math"
	"regexp"
	"strings"
)

func main() {
	fmt.Println("star1", star1())
	fmt.Println("star2", star2())
}

func star1() int {
	maxX := 101
	maxY := 103
	robots := parseInputToRobots(input)
	quads := []int{0,0,0,0}
	for i := 0; i < 100; i++ {
		for j,robot := range robots {
			movedRobot := makeMove(robot, maxX, maxY)
			robots[j] = movedRobot
			if (i == 99) {
				if (movedRobot.positionX < (maxX/2)) {
					if (movedRobot.positionY < (maxY/2)) {
						quads[0]++
						continue
					} else if movedRobot.positionY > (maxY/2) {
						quads[1]++
					}
				} else if (movedRobot.positionX > (maxX/2)) {
					if (movedRobot.positionY < (maxY/2)) {
						quads[2]++
						continue
					} else if movedRobot.positionY > (maxY/2) {
						quads[3]++
					}
				}
			}
		}
	} 
	fmt.Println(quads)
	
	return quads[0]*quads[1]*quads[2]*quads[3]
}

func star2() int {
	maxX := 101
	maxY := 103
	robots := parseInputToRobots(input)
	minSafety := 0;
	minFrame := -1;
	for i := 0; i < 10000; i++ {
		quads := []int{0,0,0,0}
		for j,robot := range robots {
			movedRobot := makeMove(robot, maxX, maxY)
			robots[j] = movedRobot
		}
		for _, robot := range robots {
			if (robot.positionX < (maxX/2)) {
				if (robot.positionY < (maxY/2)) {
					quads[0]++
					continue
				} else if robot.positionY > (maxY/2) {
					quads[1]++
				}
			} else if (robot.positionX > (maxX/2)) {
				if (robot.positionY < (maxY/2)) {
					quads[2]++
					continue
				} else if robot.positionY > (maxY/2) {
					quads[3]++
				}
			}
		}
		safety := quads[0]*quads[1]*quads[2]*quads[3];
		if (minFrame == -1 || safety < minSafety) {
			minFrame = i;
			minSafety = safety
		}
	} 
	return minFrame + 1
}

func makeMove(robot Robot, xMax, yMax int) Robot {
	robot.positionX = getNewPosition(robot.positionX + robot.moveX, xMax);
	robot.positionY = getNewPosition(robot.positionY + robot.moveY, yMax);
	return robot;
}

func getNewPosition(value, max int) int {
	if value < 0 {
		return getNewPosition(max + value, max);
	}
	if value < max {
		return value;
	}
	return value % max;
}

type Robot struct {
	positionX, positionY, moveX, moveY int
}

func parseInputToRobots(input string) []Robot {
	robots := []Robot{}
	for _, robotString := range strings.Split(input, "\n") {
		var robot Robot
		g := regexp.MustCompile(`p=(\d+),(\d+) v=([-\d]+),([-\d]+)`)
		things := g.FindStringSubmatch(robotString);
		robot.positionX = getIntFromString(things[1]);
		robot.positionY = getIntFromString(things[2]);
		robot.moveX = getIntFromString(things[3]);
		robot.moveY = getIntFromString(things[4]);
		robots = append(robots, robot)
	}
	return robots
}
