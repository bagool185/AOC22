package main

import (
	common "aoc23/main"
	"bufio"
	"fmt"
	"strings"
)

const (
	PlayerRock     = "X"
	PlayerPaper    = "Y"
	PlayerScissors = "Z"
)

const (
	OpponentRock     = "A"
	OpponentPaper    = "B"
	OpponentScissors = "C"
)

var shapeScoreBonuses = map[string]int{
	PlayerRock:     1,
	PlayerPaper:    2,
	PlayerScissors: 3,
}

func getWinningMove(opponentMove string) string {
	playerWinsMoves := map[string]string{
		OpponentRock:     PlayerPaper,
		OpponentPaper:    PlayerScissors,
		OpponentScissors: PlayerRock,
	}

	return playerWinsMoves[opponentMove]
}

func getDrawMove(opponentMove string) string {
	drawMoves := map[string]string{
		OpponentRock:     PlayerRock,
		OpponentPaper:    PlayerPaper,
		OpponentScissors: PlayerScissors,
	}

	return drawMoves[opponentMove]
}

func getLosingMove(opponentMove string) string {
	playerLosesMoves := map[string]string{
		OpponentRock:     PlayerScissors,
		OpponentPaper:    PlayerRock,
		OpponentScissors: PlayerPaper,
	}

	return playerLosesMoves[opponentMove]
}

func part1() {
	var inputFile = common.ReadTestFile("day2_input.txt")

	fileScanner := bufio.NewScanner(inputFile)
	fileScanner.Split(bufio.ScanLines)

	total := 0

	for fileScanner.Scan() {

		var line = fileScanner.Text()

		if line == "" {
			break
		}

		var shapes = strings.Split(fileScanner.Text(), " ")

		total += shapeScoreBonuses[shapes[1]]

		if getDrawMove(shapes[0]) == shapes[1] {
			total += 3
		} else if getWinningMove(shapes[0]) == shapes[1] {
			total += 6
		}
	}

	inputFile.Close()

	fmt.Printf("Part 1: total %d\n", total)
}

func part2() {
	var inputFile = common.ReadTestFile("day2_input.txt")

	fileScanner := bufio.NewScanner(inputFile)
	fileScanner.Split(bufio.ScanLines)

	total := 0

	for fileScanner.Scan() {

		var line = fileScanner.Text()

		if line == "" {
			break
		}

		var shapes = strings.Split(fileScanner.Text(), " ")

		switch shapes[1] {
		case "X":
			var losingMove = getLosingMove(shapes[0])
			total += shapeScoreBonuses[losingMove]
		case "Y":
			var drawMove = getDrawMove(shapes[0])
			total += 3 + shapeScoreBonuses[drawMove]
		case "Z":
			var winningMove = getWinningMove(shapes[0])
			total += 6 + shapeScoreBonuses[winningMove]
		}
	}

	inputFile.Close()

	fmt.Printf("Part 2: total %d\n", total)
}

func main() {
	part1()
	part2()
}
