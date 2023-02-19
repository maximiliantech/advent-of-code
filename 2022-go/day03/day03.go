package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	lines := readAllLines("./input.txt")
	log.Println("Day 03 Part 01")
	partOne(lines)
	log.Println("Day 03 Part 02")
	partTwo(lines)
}

func partOne(rucksacks []string) {
	var sum int

	for _, rucksack := range rucksacks {
		n := len(rucksack) / 2
		firstCompartment := rucksack[:n]
		secondCompartment := rucksack[n:]

		item_type := findDuplicate(firstCompartment, secondCompartment)
		sum += getPriority(item_type)
	}

	log.Println(sum)
}

func partTwo(rucksacks []string) {
	var sum int
	var groups [][]string = make([][]string, (len(rucksacks) / 3))
	var gcount int = 0

	for rcount, rucksack := range rucksacks {
		groups[gcount] = append(groups[gcount], rucksack)
		rcount += 1

		if rcount%3 == 0 {
			gcount += 1
		}
	}

	for _, group := range groups {
		badge := findBadge(group)
		sum += getPriority(badge)
	}

	log.Println(sum)
}

func findBadge(group []string) byte {
	occurence_list_first_rucksack := make(map[rune]int)

	for _, character := range group[0] {
		occurence_list_first_rucksack[character] += 1
	}

	for _, c2 := range group[1] {
		if occurence_list_first_rucksack[c2] > 0 {
			for _, c3 := range group[2] {
				if (occurence_list_first_rucksack[c3] > 0) && (c2 == c3) {
					return byte(c3)
				}
			}
		}
	}
	return byte('a')
}

func getPriority(char byte) int {
	if char >= 'a' && char <= 'z' {
		return int(char - 'a' + 1)
	} else if char >= 'A' && char <= 'Z' {
		return int(char - 'A' + 27)
	} else {
		return 0
	}
}

func findDuplicate(firstCompartment string, secondCompartment string) byte {
	occurrence_list_first_compartment := make(map[rune]int)

	for _, character := range firstCompartment {
		occurrence_list_first_compartment[character] += 1
	}

	for _, character := range secondCompartment {
		if occurrence_list_first_compartment[character] > 0 {
			return byte(character)
		}
	}

	return byte('a')
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
