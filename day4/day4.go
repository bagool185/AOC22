package main

import (
	common "aoc22/main"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Interval struct {
	lowerBound int
	upperBound int
}

func parseInterval(group string) Interval {
	var interval = strings.Split(group, "-")

	var lowerBound, _ = strconv.Atoi(interval[0])
	var upperBound, _ = strconv.Atoi(interval[1])

	return Interval{lowerBound, upperBound}
}

func main() {
	var inputFile = common.ReadTestFile("day4_input.txt")
	fileScanner := bufio.NewScanner(inputFile)
	fileScanner.Split(bufio.ScanLines)

	part1Total := 0
	part2Total := 0

	for fileScanner.Scan() {
		var line = fileScanner.Text()
		var groups = strings.Split(line, ",")

		var interval1 = parseInterval(groups[0])
		var interval2 = parseInterval(groups[1])

		if interval1.lowerBound >= interval2.lowerBound &&
			interval1.upperBound <= interval2.upperBound {
			part1Total++
		} else if interval2.lowerBound >= interval1.lowerBound &&
			interval2.upperBound <= interval1.upperBound {
			part1Total++
		}

		if interval1.lowerBound >= interval2.lowerBound &&
			interval1.lowerBound <= interval2.upperBound {
			part2Total++
		} else if interval2.lowerBound >= interval1.lowerBound &&
			interval2.lowerBound <= interval1.upperBound {
			part2Total++
		}
	}

	fmt.Printf("Part 1: total %d\n", part1Total)
	fmt.Printf("Part 2: total %d\n", part2Total)
}
