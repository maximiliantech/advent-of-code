package main

import (
	"log"
	"strconv"

	"github.com/maximiliantech/advent-of-code/2022-go/common"
)

var countZeroPartOne = 0
var countZeroPartTwo = 0

func main() {
	lines := common.ReadAllLines("./input_test.txt")

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
	rest := number + dial
	if rest >= 0 && rest <= 99 {
		return rest
	} else if rest < 0 && rest > 99 {
		return 100 - rest
	} else {
		return turnRight(rest-100, 0)
	}
}

func turnLeft(number, dial int) int {
	rest := dial - number
	if rest < 0 && rest > -100 {
		return 100 + rest
	} else if rest >= 0 && rest <= 99 {
		return rest
	} else {
		return turnLeft(-rest-100, 0)
	}
}
