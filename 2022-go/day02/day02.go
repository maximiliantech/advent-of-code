package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	lines := readAllLines("./2022-go/day02/input.txt")
	log.Println("Day 02 Part 01")
	partOne(lines)
}

func partOne(inputLines []string) {
	var opponent []string
	var me []string

	for _, line := range inputLines {
		lineArray := strings.Split(line, " ")
		opponent = append(opponent, lineArray[0])
		me = append(me, lineArray[1])
	}

	solution := playRPC(opponent, me)
	log.Println(solution)
}

func playRPC(opponent []string, me []string) int {
	rock := Rock{Score: 1, Opponent: "A", Me: "X"}
	paper := Paper{Score: 2, Opponent: "B", Me: "Y"}
	scissor := Scissor{Score: 3, Opponent: "C", Me: "Z"}
	lose := 0
	draw := 3
	win := 6
	score := 0
	for round := 0; round < len(opponent); round++ {
		//Rock draw
		if (opponent[round] == rock.Opponent) && (me[round] == rock.Me) {
			score += rock.Score + draw
		}
		//Paper draw
		if (opponent[round] == paper.Opponent) && (me[round] == paper.Me) {
			score += paper.Score + draw
		}
		//Scissor draw
		if (opponent[round] == scissor.Opponent) && (me[round] == scissor.Me) {
			score += scissor.Score + draw
		}
		//------------------------------------------
		// Rock beats Scissor
		if (opponent[round] == scissor.Opponent) && (me[round] == rock.Me) {
			score += rock.Score + win
		}
		// Scissor beats Paper
		if (opponent[round] == paper.Opponent) && (me[round] == scissor.Me) {
			score += scissor.Score + win
		}
		// Paper beats Rock
		if (opponent[round] == rock.Opponent) && (me[round] == paper.Me) {
			score += paper.Score + win
		}
		//------------------------------------------
		// Scissor lost to Rock
		if (opponent[round] == rock.Opponent) && (me[round] == scissor.Me) {
			score += scissor.Score + lose
		}
		// Paper lost to Scissor
		if (opponent[round] == scissor.Opponent) && (me[round] == paper.Me) {
			score += paper.Score + lose
		}
		// Rock lost to Paper
		if (opponent[round] == paper.Opponent) && (me[round] == rock.Me) {
			score += rock.Score + lose
		}
	}
	return score
}

type Rock struct {
	Score    int
	Opponent string
	Me       string
}

type Paper struct {
	Score    int
	Opponent string
	Me       string
}

type Scissor struct {
	Score    int
	Opponent string
	Me       string
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
