package main

import (
	"fmt"
	"log"

	"github.com/maximiliantech/advent-of-code/2022-go/common"
)

func main() {
	lines := common.ReadAllLines("./2022-go/day04/input.txt")
	log.Println("Day 04 Part 01")
	partOne(lines)
	log.Println("Day 04 Part 02")
	partTwo(lines)
}

func partOne(lines []string) {
	pairs := getPairs(lines)

	count := countPairs(pairs, func(pair1, pair2 []int) bool {
		return ((pair1[0] <= pair2[0]) && (pair1[1] >= pair2[1])) || ((pair1[0] >= pair2[0]) && (pair1[1] <= pair2[1]))
	})

	log.Println(count)
}

func partTwo(lines []string) {
	pairs := getPairs(lines)

	count := countPairs(pairs, func(pair1, pair2 []int) bool {
		return (pair1[1] >= pair2[0]) && (pair1[0] <= pair2[1])
	})

	log.Println(count)
}

func countPairs(pairs [][]int, f func(pair1, pair2 []int) bool) int {
	count := 0

	for _, pair := range pairs {
		if f(pair[:2], pair[2:]) {
			count++
		}
	}

	return count
}

func getPairs(lines []string) [][]int {
	pairs := make([][]int, len(lines))
	for i, line := range lines {
		pair := make([]int, 4)
		fmt.Sscanf(line, "%d-%d,%d-%d", &pair[0], &pair[1], &pair[2], &pair[3])
		pairs[i] = pair
	}
	return pairs
}
