package main

import (
	"github.com/maximiliantech/advent-of-code/2022-go/common"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	lines := common.ReadAllLines("../2024-go/day01/input_test.txt")
	log.Println("Day 01 Part 01")
	partOne(lines)
	log.Println("Day 01 Part 02")
}

func partOne(lines []string) {
	leftList := []float64{}
	rightList := []float64{}
	for _, line := range lines {
		s := strings.Split(line, "   ") // split at 3 spaces
		left, _ := strconv.Atoi(s[0])
		right, _ := strconv.Atoi(s[1])
		leftList = append(leftList, float64(left))
		rightList = append(rightList, float64(right))
	}
	sort.Float64s(leftList)
	sort.Float64s(rightList)
	diffs := float64(0)
	for i := range leftList {
		diffs += math.Abs(leftList[i] - rightList[i])
	}
	log.Println(diffs)
}
