package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	lines := readAllLines("./2022/day01/input.txt")

	log.Println("Day 01 Part 01")
	partOne(lines)
}

func partOne(lines []string) {
	elfCaloriesMap := getAllocatedCaloriesPerElf(lines)
	elf, calories := getElfMostCalories(elfCaloriesMap)
	log.Println(elf, calories)
}

func getElfMostCalories(elfMap map[string]int) (string, int) {
	var elf string = "1"
	var calories int = 0
	for key, value := range elfMap {
		if value > calories {
			calories = value
			elf = key
		}
	}
	return elf, calories
}

func getAllocatedCaloriesPerElf(lines []string) map[string]int {
	var elfCalories = make(map[string]int)
	var elfNumber int = 1
	for _, stringCalories := range lines {
		if stringCalories == "" {
			elfNumber++
		} else {
			calorie, _ := strconv.Atoi(stringCalories)
			elfCalories[strconv.Itoa(elfNumber)] += calorie
		}
	}
	return elfCalories
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
