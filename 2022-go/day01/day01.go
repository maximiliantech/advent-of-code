package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	lines := readAllLines("./2022-go/day01/input.txt")

	log.Println("Day 01 Part 01")
	partOne(lines)
	//log.Println("Day 01 Part 02")
	//partTwo(lines)
}

func partOne(lines string) {
	var elves = getElves(lines)
	elvesMostCalories := getElvesMostCalories(elves, 1)
	solution := 0
	for _, elf := range elvesMostCalories {
		solution += elf.CaloriesSum
	}
	log.Println(solution)
}

func partTwo(lines []string) {

}

func getElvesMostCalories(elves []Elf, count int) []Elf {
	elvesMostCalories := make([]Elf, count)
	sort.Slice(elves, func(i, j int) bool {
		return elves[i].CaloriesSum > elves[j].CaloriesSum
	})
	for i, _ := range elvesMostCalories {
		elvesMostCalories[i] = elves[i]
	}
	return elvesMostCalories
}

type Elf struct {
	Number      string
	CaloriesSum int
}

func getElves(lines string) []Elf {
	var elves []Elf
	elvesOriginal := strings.Split(lines, "\n\n")
	for index, calorieBundle := range elvesOriginal {
		var elf Elf
		caloriesPerElf := strings.Split(calorieBundle, "\n")
		elf.Number = strconv.Itoa(index + 1)
		elf.CaloriesSum = sumOfStringArray(caloriesPerElf)
		elves = append(elves, elf)
	}
	return elves
}

func sumOfStringArray(arr []string) int {
	var result = 0
	for _, v := range arr {
		var number, _ = strconv.Atoi(v)
		result += number
	}
	return result
}

func readAllLines(filePath string) string {
	b, err := os.ReadFile(filePath) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)
	return str
}
