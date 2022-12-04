package main

import (
	common "aoc23/main"
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func getItemPriority(item rune) int {
	if unicode.IsUpper(item) {
		return int(item) - 'A' + 26 + 1
	}

	return int(item) - 'a' + 1
}

func part1(inputFile *os.File) {
	fileScanner := bufio.NewScanner(inputFile)
	fileScanner.Split(bufio.ScanLines)

	total := 0

	for fileScanner.Scan() {
		var line = fileScanner.Text()

		var items map[rune]bool = make(map[rune]bool)

		for index, char := range line {

			if index < len(line)/2 {
				items[char] = true
			} else {
				if items[char] {

					total += getItemPriority(char)

					break
				}
			}
		}
	}

	fmt.Printf("Part 1: total %d\n", total)
}

func dedupeRunes(text string) string {
	uniqueValues := ""

	for _, char := range text {
		if !strings.ContainsRune(uniqueValues, char) {
			uniqueValues += string(char)
		}
	}

	return uniqueValues
}

func part2(inputFile *os.File) {
	fileScanner := bufio.NewScanner(inputFile)
	fileScanner.Split(bufio.ScanLines)

	total := 0

	for fileScanner.Scan() {
		var lines []string

		var items map[rune]int = make(map[rune]int)

		lines = append(lines, dedupeRunes(fileScanner.Text()))

		for i := 0; i < 2; i++ {
			fileScanner.Scan()
			lines = append(lines, dedupeRunes(fileScanner.Text()))
		}

		for _, line := range lines {
			for _, char := range line {
				items[char]++

				if items[char] == 3 {
					total += getItemPriority(char)
				}
			}
		}
	}

	fmt.Printf("Part 2: total %d\n", total)
}

func main() {
	var inputFile = common.ReadTestFile("day3_input.txt")

	// part1(inputFile)
	part2(inputFile)

	inputFile.Close()
}
