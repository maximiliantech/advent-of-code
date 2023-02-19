package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	lines := readAllLines("./input.txt")
	log.Println("Day 04 Part 01")
	partOne(lines)
	log.Println("Day 04 Part 02")
	partTwo(lines)
}

func partOne(lines []string) {
	pairs := getPairs(lines)

	count := countContainingPairs(pairs)

	log.Println(count)
}

func partTwo(lines []string) {
	pairs := getPairs(lines)

	count := countOverlappingPairs(pairs)

	log.Println(count)
}

func countContainingPairs(pairs [][]int) int {
	count := 0

	for _, pair := range pairs {
		pair1 := []int{pair[0], pair[1]}
		pair2 := []int{pair[2], pair[3]}
		if contains(pair1, pair2) {
			count++
		}
	}

	return count
}

func contains(pair1, pair2 []int) bool {
	return ((pair1[0] <= pair2[0]) && (pair1[1] >= pair2[1])) || ((pair1[0] >= pair2[0]) && (pair1[1] <= pair2[1]))
}

func overlaps(pair1, pair2 []int) bool {
	return (pair1[1] >= pair2[0]) && (pair1[0] <= pair2[1])
}

func countOverlappingPairs(pairs [][]int) int {
	count := 0

	for _, pair := range pairs {
		pair1 := []int{pair[0], pair[1]}
		pair2 := []int{pair[2], pair[3]}
		if overlaps(pair1, pair2) {
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

func readAllLines(filePath string) []string {
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines
}
