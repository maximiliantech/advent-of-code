package main

import (
	"log"

	"github.com/maximiliantech/advent-of-code/2022-go/common"
)

func main() {
	lines := common.ReadAllLines("./2022-go/day06/input.txt")
	log.Println("Day 06 Part 01")
	partOne(lines)
	log.Println("Day 06 Part 02")
	partTwo(lines)
}

func partOne(lines []string) {
	var datastream string = lines[0]
	var marker int = 0
	for i := range datastream {
		if (i >= 3) && !duplicate(datastream[i-3:i+1]) {
			marker = i + 1
			break
		}
	}
	log.Println(marker)
}

func partTwo(lines []string) {
	var datastream string = lines[0]
	var marker int = 0
	for i := range datastream {
		if (i >= 13) && !duplicate(datastream[i-13:i+1]) {
			marker = i + 1
			break
		}
	}
	log.Println(marker)
}

func duplicate(slice string) bool {
	counts := map[rune]int{}
	for _, char := range slice {
		if counts[char] == 1 {
			return true
		}
		counts[char]++
	}
	return false
}
