package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	lines := readAllLines("./2022-go/day01/input_test.txt")

	log.Println("Day 01 Part 01")
	partOne(lines)
	//log.Println("Day 01 Part 02")
	//partTwo(lines)
}

func partOne(lines []string) {
	elves := getElves(lines)
	elvesMostCalories := getElvesMostCalories(elves, 1)
	solution := 0
	for _, elf := range elvesMostCalories {
		solution += elf.CaloriesSum
	}
	log.Println(solution)
}

func partTwo(lines []string) {
	elves := getElves(lines)
	elvesMostCalories := getElvesMostCalories(elves, 3)
	solution := 0
	for _, elf := range elvesMostCalories {
		solution += elf.CaloriesSum
	}
	log.Println(solution)
}

func getElvesMostCalories(elfs []Elf, count int) []Elf {
	elvesMostCalories := make([]Elf, count)
	sort.Slice(elfs, func(i, j int) bool {
		return elfs[i].CaloriesSum > elfs[j].CaloriesSum
	})
	for i, _ := range elvesMostCalories {
		elvesMostCalories[i] = elfs[i]
	}
	return elvesMostCalories
}

type Elf struct {
	Number      string
	CaloriesSum int
}

func getElves(lines []string) []Elf {
	var elves []Elf
	var elfNumber = 0
	for _, line := range lines {
		if line == "" {
			elfNumber++
		} else {
			elves[elfNumber].Number = strconv.Itoa(elfNumber + 1)
			calorie, _ := strconv.Atoi(line)
			elves[elfNumber].CaloriesSum += calorie
		}
	}
	return elves
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
