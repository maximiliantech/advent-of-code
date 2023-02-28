package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/maximiliantech/advent-of-code/2022-go/common"
)

func main() {
	lines := common.ReadAllLines("./input.txt")
	log.Println("Day 05 Part 01")
	partOne(lines)
	log.Println("Day 05 Part 02")
	partOne(lines)
}

func partOne(lines []string) {
	max_height_stacks := 0
	stacks_count := 0
	procedure_string := []string{}
	collect_procedure := false
	procedure_count := 0

	// check input for dimensions
	for i, line := range lines {
		if collect_procedure {
			procedure_string = append(procedure_string, line)
			procedure_count += 1
		}

		if len(line) == 0 {
			collect_procedure = true
			max_height_stacks = i - 1
			l := lines[i-1]
			stacks_count, _ = strconv.Atoi(string(l[len(l)-2]))
		}
	}

	stacks := getCrates(lines, stacks_count, max_height_stacks)

	procedure := transformRerangeProdecure(procedure_string, procedure_count)

	stacks = rerange(stacks, procedure)
	top_crates := getTopCrates(stacks)
	log.Println(top_crates)
}

func partTwo(lines []string) {
	max_height_stacks := 0
	stacks_count := 0
	procedure_string := []string{}
	collect_procedure := false
	procedure_count := 0

	// check input for dimensions
	for i, line := range lines {
		if collect_procedure {
			procedure_string = append(procedure_string, line)
			procedure_count += 1
		}

		if len(line) == 0 {
			collect_procedure = true
			max_height_stacks = i - 1
			l := lines[i-1]
			stacks_count, _ = strconv.Atoi(string(l[len(l)-2]))
		}
	}

	stacks := getCrates(lines, stacks_count, max_height_stacks)

	procedure := transformRerangeProdecure(procedure_string, procedure_count)

	stacks = rerangeWithAtOnce(stacks, procedure)
	top_crates := getTopCrates(stacks)
	log.Println(top_crates)
}

func rerange(stacks [][]byte, procedure [][]int) [][]byte {
	for _, proc := range procedure {
		for i := 0; i < proc[0]; i++ {
			stacks[proc[2]-1] = append(stacks[proc[2]-1], stacks[proc[1]-1][len(stacks[proc[1]-1])-1])
			stacks[proc[1]-1] = stacks[proc[1]-1][:len(stacks[proc[1]-1])-1]
		}
	}
	return stacks
}

func rerangeWithAtOnce(stacks [][]byte, procedure [][]int) [][]byte {
	for _, proc := range procedure {
		for i := 0; i < proc[0]; i++ {
			stacks[proc[2]-1] = append(stacks[proc[2]-1], stacks[proc[1]-1][len(stacks[proc[1]-1])-1])
			stacks[proc[1]-1] = stacks[proc[1]-1][:len(stacks[proc[1]-1])-1]
		}
	}
	return stacks
}

func getTopCrates(stacks [][]byte) string {
	s := ""
	for _, stack := range stacks {
		s += string(stack[len(stack)-1])
	}
	return s
}

func transformRerangeProdecure(lines []string, count int) [][]int {
	procedure := make([][]int, count)
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
				stacks[j] = append(stacks[j], s[j*4+1])
			}
		}
	}
	return stacks
}
