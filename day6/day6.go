package main

import (
	common "aoc23/main"
	"bufio"
	"fmt"
)

func isMarker(text string) bool {

	var flags [256]bool

	for _, ascii := range text {
		if flags[ascii] {
			return false
		}

		flags[ascii] = true
	}

	return true
}

func main() {
	inputFile := common.ReadTestFile("day6_input.txt")
	scanner := bufio.NewScanner((inputFile))

	scanner.Split(bufio.ScanRunes)

	var crtMarkerCandidate []rune
	index := 1

	for scanner.Scan() {
		char := scanner.Text()

		crtMarkerCandidate = append(crtMarkerCandidate, rune(char[0]))

		if len(crtMarkerCandidate) > 14 {
			crtMarkerCandidate = crtMarkerCandidate[1:]

			if isMarker(string(crtMarkerCandidate)) {
				fmt.Printf("Marker after character: #%d", index)
				break
			}
		}

		index++
	}
}
