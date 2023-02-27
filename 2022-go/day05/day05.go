package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/maximiliantech/advent-of-code/2022-go/common"
)

func main() {
	lines := common.ReadAllLines("./input_test.txt")
	log.Println("Day 05 Part 01")
	partOne(lines)
}

func partOne(lines []string) {
	input_index := 0

	//max_height_stacks := 0
	//stacks_count := 0

	// check input for dimensions
	for i, line := range lines {
		log.Println(line, len(line), i, input_index)
		if len(line) == 0 {
			//max_height_stacks = i
			//stacks_count = len(strings.Split(lines[i-1], " "))
			log.Println(strings.Split(lines[i-1], " "))
		}
	}

	//stacks := make([][]string, stacks_count)

	//stacks := rerange()
	//top_crates := getTopCrates(stacks)
	//log.Println(top_crates)
}

func rerange(stacks []string, procedure [][]int) []string {
	return []string{}
}

func getTopCrates(stacks []string) string {
	return ""
}

func transformRerangeProdecure(lines []string) [][]int {
	var procedure [][]int
	for i, line := range lines {
		p := make([]int, 3)
		fmt.Sscanf(line, "move %d from %d to %d", &p[0], &p[1], &p[2])
		procedure[i] = p
	}
	return procedure
}
