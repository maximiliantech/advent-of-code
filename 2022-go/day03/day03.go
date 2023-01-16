package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	lines := readAllLines("./2022-go/day03/input_test.txt")
	log.Println("Day 02 Part 01")
	partOne(lines)
}

func partOne(inputLines []string) {

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
