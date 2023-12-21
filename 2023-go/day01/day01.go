package main

import (
	"github.com/maximiliantech/advent-of-code/2022-go/common"
	"log"
	"strconv"
)

func main() {
	lines := common.ReadAllLines("../2023-go/day01/input.txt")
	log.Println("Day 01 Part 01")
	partOne(lines)
	log.Println("Day 01 Part 02")
	partTwo(lines)
}

func partOne(lines []string) {
	var sum int
	for _, line := range lines {
		value := readCalibrationValue(line)
		sum += value
	}
	log.Println(sum)
}

// readCalibrationValue reads the calibration value from a line.
func readCalibrationValue(line string) int {
	var firstDigit int
	var lastDigit int
	for _, char := range line {
		char := int(char)
		// if char is a digit
		if isActualDigit(char) {
			firstDigit = char
			break
		}
	}
	for i := len(line) - 1; i >= 0; i-- {
		char := int(line[i])

		// if char is a digit
		if isActualDigit(char) {
			lastDigit = char
			break
		}
	}
	n, _ := strconv.Atoi(string(rune(firstDigit)) + "" + string(rune(lastDigit)))
	return n
}

func partTwo(lines []string) {
	var sum int
	for _, line := range lines {
		value := readCalibrationValue2(line)
		sum += value
	}
	log.Println(sum)
}

// please create a function that solves the problem
func readCalibrationValue2(line string) int {
	digits := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	// search for first occurrence of spelled out digit or actual digit
	for i := 0; i < len(line); i++ {

	}
	return 0
}

func isActualDigit(char int) bool {
	return char >= '1' && char <= '9'
}
