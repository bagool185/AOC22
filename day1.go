package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {

	inputFile, err := os.Open("./day1_input.txt")

	if err != nil {
		fmt.Printf("bronk AAAAAAAAAAAA\n%s", err)
	}

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
