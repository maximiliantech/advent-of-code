package main

import (
	"log"
	"strconv"

	"github.com/maximiliantech/advent-of-code/2022-go/common"
)

const maxDialPosition = 99
const minDialPosition = 0

var countZeroPartOne = 0
var countZeroPartTwo = 0

func main() {
	lines := common.ReadAllLines("./input.txt")

	dial := 50

	for _, line := range lines {
		direction := string([]rune(line)[0])
		number, err := strconv.Atoi(string([]rune(line)[1:]))
		if err != nil {
			panic(0)
		}

		// turn left
		if direction == "L" {
			dial = turnLeft(number, dial)
		} else { // turn right
			dial = turnRight(number, dial)
		}

		if dial == 0 {
			countZeroPartOne++
		}
	}
	log.Println("Day 01 Part 01")
	log.Println(countZeroPartOne)
	log.Println("Day 01 Part 02")
	log.Println(countZeroPartOne + countZeroPartTwo)
}

func turnRight(number, dial int) int {
	newDial := number + dial
	if newDial >= minDialPosition && newDial <= maxDialPosition {
		return newDial
	} else if newDial == 100 {
		return 0
	} else {
		if dial != 0 {
			countZeroPartTwo++
		}
		return turnRight(newDial-100, 0)
	}
}

func turnLeft(number, dial int) int {
	newDial := dial - number

	if newDial >= minDialPosition && newDial <= maxDialPosition {
		return newDial
	} else if newDial == 0 {
		return 0
	} else {
		if dial != 0 {
			countZeroPartTwo++
		}
		return turnLeft(-newDial-100, 0)
	}
}
