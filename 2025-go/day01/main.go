package main

import (
	"log"
	"strconv"

	"github.com/maximiliantech/advent-of-code/2022-go/common"
)

func main() {
	lines := common.ReadAllLines("./input.txt")
	log.Println("Day 01 Part 01")
	partOne(lines)
}

func partOne(lines []string) {

	dial := 50
	countZero := 0

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
			countZero++
		}
	}
	log.Println(countZero)
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
