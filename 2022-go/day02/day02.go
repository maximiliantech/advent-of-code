package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	lines := readAllLines("./input.txt")
	log.Println("Day 02 Part 01")
	partOne(lines)
	log.Println("Day 02 Part 02")
	partTwo(lines)
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

func partTwo(inputLines []string) {
	var opponent []string
	var me []string

	for _, line := range inputLines {
		lineArray := strings.Split(line, " ")
		opponent = append(opponent, lineArray[0])
		me = append(me, lineArray[1])
	}

	solution := playRPCWithTarget(opponent, me)
	log.Println(solution)
}

func playRPC(opponent []string, me []string) int {
	rock := Rock{Score: 1, Opponent: "A", Me: "X"}
	paper := Paper{Score: 2, Opponent: "B", Me: "Y"}
	scissor := Scissor{Score: 3, Opponent: "C", Me: "Z"}
	score := Score{Lose: 0, Draw: 3, Win: 6}
	totalScore := 0
	for round := 0; round < len(opponent); round++ {
		//Rock draw
		if (opponent[round] == rock.Opponent) && (me[round] == rock.Me) {
			totalScore += rock.Score + score.Draw
		}
		//Paper draw
		if (opponent[round] == paper.Opponent) && (me[round] == paper.Me) {
			totalScore += paper.Score + score.Draw
		}
		//Scissor draw
		if (opponent[round] == scissor.Opponent) && (me[round] == scissor.Me) {
			totalScore += scissor.Score + score.Draw
		}
		//------------------------------------------
		// Rock beats Scissor
		if (opponent[round] == scissor.Opponent) && (me[round] == rock.Me) {
			totalScore += rock.Score + score.Win
		}
		// Scissor beats Paper
		if (opponent[round] == paper.Opponent) && (me[round] == scissor.Me) {
			totalScore += scissor.Score + score.Win
		}
		// Paper beats Rock
		if (opponent[round] == rock.Opponent) && (me[round] == paper.Me) {
			totalScore += paper.Score + score.Win
		}
		//------------------------------------------
		// Scissor lost to Rock
		if (opponent[round] == rock.Opponent) && (me[round] == scissor.Me) {
			totalScore += scissor.Score + score.Lose
		}
		// Paper lost to Scissor
		if (opponent[round] == scissor.Opponent) && (me[round] == paper.Me) {
			totalScore += paper.Score + score.Lose
		}
		// Rock lost to Paper
		if (opponent[round] == paper.Opponent) && (me[round] == rock.Me) {
			totalScore += rock.Score + score.Lose
		}
	}
	return totalScore
}

func playRPCWithTarget(opponent []string, me []string) int {
	rock := Rock{Score: 1, Opponent: "A", Me: "X"}
	paper := Paper{Score: 2, Opponent: "B", Me: "Y"}
	scissor := Scissor{Score: 3, Opponent: "C", Me: "Z"}
	decision := Decision{Lose: "X", Draw: "Y", Win: "Z"}
	score := Score{Lose: 0, Draw: 3, Win: 6}
	totalScore := 0
	for round := 0; round < len(opponent); round++ {
		// --------------- Opponent chooses Rock ----------------
		if (opponent[round] == rock.Opponent) && (me[round] == decision.Lose) {
			totalScore += scissor.Score + score.Lose
		}
		if (opponent[round] == rock.Opponent) && (me[round] == decision.Draw) {
			totalScore += rock.Score + score.Draw
		}
		if (opponent[round] == rock.Opponent) && (me[round] == decision.Win) {
			totalScore += paper.Score + score.Win
		}
		// ---------------- Opponent chooses Scissor --------------------
		if (opponent[round] == scissor.Opponent) && (me[round] == decision.Lose) {
			totalScore += paper.Score + score.Lose
		}
		if (opponent[round] == scissor.Opponent) && (me[round] == decision.Draw) {
			totalScore += scissor.Score + score.Draw
		}
		if (opponent[round] == scissor.Opponent) && (me[round] == decision.Win) {
			totalScore += rock.Score + score.Win
		}
		// ------------ Opponent chooses Paper -----------------
		if (opponent[round] == paper.Opponent) && (me[round] == decision.Lose) {
			totalScore += rock.Score + score.Lose
		}
		if (opponent[round] == paper.Opponent) && (me[round] == decision.Draw) {
			totalScore += paper.Score + score.Draw
		}
		if (opponent[round] == paper.Opponent) && (me[round] == decision.Win) {
			totalScore += scissor.Score + score.Win
		}
	}
	return totalScore
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

type Decision struct {
	Lose string
	Draw string
	Win  string
}

type Score struct {
	Lose int
	Draw int
	Win  int
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
