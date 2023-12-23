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

// readCalibrationValue2 reads the calibration value from a line.
func readCalibrationValue2(line string) int {
	var firstDigit int
	var lastDigit int
	digits := map[string]int{
		"one":   49, // ASCII code for 1 etc.
		"two":   50,
		"three": 51,
		"four":  52,
		"five":  53,
		"six":   54,
		"seven": 55,
		"eight": 56,
		"nine":  57,
	}

	// search for first occurrence of spelled out digit or actual digit
	firstDigit = searchForCalibrationValueFromLeftToRight(line, digits)

	lastDigit = searchForCalibrationValueFromRightToLeft(line, digits)

	n, _ := strconv.Atoi(string(rune(firstDigit)) + "" + string(rune(lastDigit)))
	return n
}

// isActualDigit checks if a character is a digit between 1 and 9.
func isActualDigit(char int) bool {
	return char >= '1' && char <= '9'
}

// searchForCalibrationValueFromLeftToRight searches for the first occurrence of a spelled out digit or actual digit from left to right
func searchForCalibrationValueFromLeftToRight(line string, digits map[string]int) int {
	for i := 0; i < len(line); i++ {
		for j, k := range digits {
			if len(j)+i <= len(line) {
				word := line[i : i+len(j)]
				if word == j {
					return k
				}
			}
		}
		char := int(line[i])
		if isActualDigit(char) {
			return char
		}
	}
	return 0
}

// searchForCalibrationValueFromRightToLeft searches for the first occurrence of a spelled out digit or actual digit from right to left
func searchForCalibrationValueFromRightToLeft(line string, digits map[string]int) int {
	for i := len(line) - 1; i >= 0; i-- {
		for j, k := range digits {
			if i-len(j) >= 0 {
				word := line[i-len(j)+1 : i+1]
				if word == j {
					return k
				}
			}
		}
		char := int(line[i])
		if isActualDigit(char) {
			return char
		}
	}
	return 0
}
