package main

import (
	common "aoc23/main"
	"bufio"
	"fmt"
	"strings"
)

const PlayerRock = "X"
const PlayerPaper = "Y"
const PlayerScissors = "Z"

const OpponentRock = "A"
const OpponentPaper = "B"
const OpponentScissors = "C"

func isWinningMove(opponent string, player string) bool {
	playerWinsMoves := map[string]string{
		OpponentRock:     PlayerPaper,
		OpponentPaper:    PlayerScissors,
		OpponentScissors: PlayerRock,
	}

	return playerWinsMoves[opponent] == player
}

func isDraw(opponent string, player string) bool {
	drawMoves := map[string]string{
		OpponentRock:     PlayerRock,
		OpponentPaper:    PlayerPaper,
		OpponentScissors: PlayerScissors,
	}

	return drawMoves[opponent] == player
}

func main() {
	var inputFile = common.ReadTestFile("day2_input.txt")

	fileScanner := bufio.NewScanner(inputFile)
	fileScanner.Split(bufio.ScanLines)

	shapeScoreBonuses := map[string]int{
		PlayerRock:     1,
		PlayerPaper:    2,
		PlayerScissors: 3,
	}

	total := 0

	for fileScanner.Scan() {

		var line = fileScanner.Text()

		if line == "" {
			break
		}

		var shapes = strings.Split(fileScanner.Text(), " ")

		total += shapeScoreBonuses[shapes[1]]

		if isDraw(shapes[0], shapes[1]) {
			total += 3
		} else if isWinningMove(shapes[0], shapes[1]) {
			total += 6
		}
	}

	fmt.Print(total)
}
