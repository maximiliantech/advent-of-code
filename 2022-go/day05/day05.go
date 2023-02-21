package main

import (
	"fmt"
	"log"

	"github.com/maximiliantech/advent-of-code/2022-go/common"
)

func main() {
	lines := common.ReadAllLines("./input_test.txt")
	log.Println("Day 05 Part 01")
	partOne(lines)
}

func partOne(lines []string) {

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
