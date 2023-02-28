package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/maximiliantech/advent-of-code/2022-go/common"
)

func main() {
	lines := common.ReadAllLines("./input_test.txt")
	log.Println("Day 05 Part 01")
	partOne(lines)
}

func partOne(lines []string) {
	max_height_stacks := 0
	stacks_count := 0
	procedure_string := []string{}

	// check input for dimensions
	for i, line := range lines {
		collect_procedure := false

		if collect_procedure {
			procedure_string = append(procedure_string, line)
		}

		if len(line) == 0 {
			collect_procedure = true
			max_height_stacks = i - 1
			l := lines[i-1]
			stacks_count, _ = strconv.Atoi(string(l[len(l)-2]))
		}
	}

	stacks := getCrates(lines, stacks_count, max_height_stacks)

	procedure := transformRerangeProdecure(procedure_string)

	stacks = rerange(stacks, procedure)
	top_crates := getTopCrates(stacks)
	log.Println(top_crates)
}

func rerange(stacks [][]string, procedure [][]int) [][]string {
	return [][]string{}
}

func getTopCrates(stacks [][]string) string {
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

func getCrates(lines []string, stacks_count, max_height int) [][]string {
	stacks := make([][]string, stacks_count)
	for i := 1; i <= max_height; i++ {
		s := lines[i-1]

	}
	return stacks
}
