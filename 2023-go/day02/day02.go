package main

import (
	"github.com/maximiliantech/advent-of-code/2022-go/common"
	"log"
	"strconv"
	"strings"
)

func main() {
	lines := common.ReadAllLines("../2023-go/day02/input.txt")
	games := readGames(lines)
	log.Println("Day 02 Part 01")
	partOne(games)
	log.Println("Day 02 Part 02")
	partTwo(games)
}

func partOne(games []game) {
	possibleGames := countPossibleGames(games)
	log.Println(possibleGames)
}

func countPossibleGames(games []game) int {
	possibleGames := 0
	for _, game := range games {
		if isGamePossible(game) {
			possibleGames += game.id
		}
	}
	return possibleGames
}

func isGamePossible(game game) bool {
	for _, subset := range game.subsets {
		if subset.red > 12 || subset.green > 13 || subset.blue > 14 {
			return false
		}
	}
	return true
}

func partTwo(games []game) {
	var sum int
	for _, game := range games {
		sub := findMinimumSubset(game)
		sum += sub.red * sub.green * sub.blue
	}
	log.Println(sum)
}

func findMinimumSubset(game game) subset {
	minimumSubset := subset{red: 0, green: 0, blue: 0} // start low to find the minimum
	for _, subset := range game.subsets {
		if subset.green > minimumSubset.green {
			minimumSubset.green = subset.green
		}
		if subset.red > minimumSubset.red {
			minimumSubset.red = subset.red
		}
		if subset.blue > minimumSubset.blue {
			minimumSubset.blue = subset.blue
		}
	}
	return minimumSubset
}

// ------- Helper functions to parse the input lines into structs of games and subsets -------

type game struct {
	id      int
	subsets []subset
}

type subset struct {
	red   int
	green int
	blue  int
}

func readGames(lines []string) []game {
	var games []game
	for _, line := range lines {
		games = append(games, readGame(line))
	}
	return games
}

func readGame(line string) game {
	g := game{}
	s := strings.Split(line, ": ")
	g.id = readGameId(s[0])
	g.subsets = readSubsets(s[1])
	return g
}

func readGameId(s string) int {
	s = strings.TrimPrefix(s, "Game ")
	n, _ := strconv.Atoi(s)
	return n
}

func readSubsets(subsetsRaw string) []subset {
	var subsets []subset
	sub := strings.Split(subsetsRaw, "; ")
	for _, subsetRaw := range sub {
		subsets = append(subsets, readSubset(subsetRaw))
	}
	return subsets
}

func readSubset(subsetRaw string) subset {
	subset := subset{}
	s := strings.Split(subsetRaw, ", ")
	for _, color := range s {
		c := strings.Split(color, " ")
		switch c[1] {
		case "red":
			subset.red, _ = strconv.Atoi(c[0])
		case "green":
			subset.green, _ = strconv.Atoi(c[0])
		case "blue":
			subset.blue, _ = strconv.Atoi(c[0])
		}
	}
	return subset
}
