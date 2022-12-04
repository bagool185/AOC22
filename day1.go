package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	inputFile, err := os.Open("./day1_input.txt")

	if err != nil {
		fmt.Printf("bronk AAAAAAAAAAAA\n%s", err)
	}

	fileScanner := bufio.NewScanner(inputFile)
	fileScanner.Split(bufio.ScanLines)

	var maxTotal int = 0
	var currentTotal int = 0

	for fileScanner.Scan() {
		var line = fileScanner.Text()

		if line != "" {
			lineValue, err := strconv.Atoi(line)

			if err == nil {
				currentTotal += lineValue
			}
		} else {
			if currentTotal > maxTotal {
				maxTotal = currentTotal
			}

			currentTotal = 0
		}
	}

	inputFile.Close()

	fmt.Print(maxTotal)
}
