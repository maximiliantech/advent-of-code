package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	lines := readAllLines("./2022/day01/input_test.txt")

	log.Println("Day 01 Part 01")
	partOne(lines)

	log.Println()
}

func partOne(lines []string) {
	log.Println(lines)
}

func readAllLines(filePath string) []string {
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines
}
