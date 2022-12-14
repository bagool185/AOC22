package main

import (
	common "aoc22/main"
	"bufio"
	"fmt"
	"sort"
	"strconv"
)

func main() {

	var inputFile = common.ReadTestFile("day1_input.txt")
	fileScanner := bufio.NewScanner(inputFile)
	fileScanner.Split(bufio.ScanLines)

	var totals []int
	var currentTotal int = 0

	for fileScanner.Scan() {
		var line = fileScanner.Text()

		if line != "" {
			lineValue, err := strconv.Atoi(line)

			if err == nil {
				currentTotal += lineValue
			}
		} else {
			totals = append(totals, currentTotal)

			currentTotal = 0
		}
	}

	inputFile.Close()

	sort.Slice(totals, func(a, b int) bool {
		return totals[a] > totals[b]
	})

	fmt.Print(totals[0] + totals[1] + totals[2])
}
