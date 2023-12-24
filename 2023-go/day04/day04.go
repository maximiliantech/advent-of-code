package main

import (
	"github.com/maximiliantech/advent-of-code/2022-go/common"
	"log"
	"strconv"
	"strings"
)

func main() {
	lines := common.ReadAllLines("../2023-go/day04/input.txt")
	cards := readCards(lines)
	log.Println("Day 04 Part 01")
	partOne(cards)
	log.Println("Day 04 Part 02")
}

func partOne(cards []card) {
	var sum int
	for _, c := range cards {
		sum += c.determinePoints()
	}
	log.Println(sum)
}

// ------- Helper functions and structs -------
type card struct {
	id             int
	winningNumbers []int
	myNumbers      []int
}

// determinePoints determines the points for a card.
// The points are determined whether myNumbers appear in winningNumbers.
func (c *card) determinePoints() int {
	var points int
	for _, n := range c.myNumbers {
		for _, w := range c.winningNumbers {
			if n == w {
				points *= 2
				if points == 0 {
					points = 1
				}
				break
			}
		}
	}
	return points
}

// readCards reads a slice of strings and returns a slice of cards.
func readCards(lines []string) []card {
	cards := []card{}
	for _, line := range lines {
		newCard := &card{}
		meta := strings.Split(line, ": ")
		newCard.id, _ = strconv.Atoi(strings.Split(meta[0], " ")[1])
		numbers := strings.Split(meta[1], " | ")
		newCard.winningNumbers = readNumbers(numbers[0])
		newCard.myNumbers = readNumbers(numbers[1])
		cards = append(cards, *newCard)
	}
	return cards
}

// readNumbers reads a string of numbers and returns a slice of int.
func readNumbers(line string) []int {
	numbers := []int{}
	for _, number := range strings.Split(line, " ") {
		n, _ := strconv.Atoi(number)
		if n != 0 {
			numbers = append(numbers, n)
		}
	}
	return numbers
}
