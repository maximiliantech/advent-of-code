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

func rerange(stacks [][]byte, procedure [][]int) [][]byte {
	return [][]byte{}
}

func getTopCrates(stacks [][]byte) string {
	s := ""
	for _, stack := range stacks {
		s += string(stack[len(stack)-1])
	}
	return s
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

func getCrates(lines []string, stacks_count, max_height int) [][]byte {
	stacks := make([][]byte, stacks_count)
	for i := max_height - 1; i >= 0; i-- {
		s := lines[i]
		for j := 0; j < stacks_count; j++ {
			if s[j*4+1] != 32 {
				stacks[max_height-1-i] = append(stacks[max_height-1-i], s[j*4+1])
				log.Println(string(s[j*4+1]), j+1, max_height-1-i)
			}
		}
	}
	return stacks
}
