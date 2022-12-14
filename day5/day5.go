package main

import (
	common "aoc22/main"
	"bufio"
	"fmt"
	"regexp"
	"strconv"
)

type Deque struct {
	elements []string
}

func (s *Deque) removeLast() string {
	lastElem := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return lastElem
}

func (s *Deque) removeFirst() string {
	firstElem := s.elements[0]
	s.elements = s.elements[1:len(s.elements)]
	return firstElem
}

func (s *Deque) pushFront(elem string) {
	s.elements = append([]string{elem}, s.elements...)
}

func (s *Deque) pushBack(elem string) {
	s.elements = append(s.elements, elem)
}

func (s *Deque) top() string {
	return s.elements[len(s.elements)-1]
}

func (s *Deque) isEmpty() bool {
	return len(s.elements) == 0
}

func sanitise(elem string) string {
	garboCharsRe := regexp.MustCompile("\\[|\\]")
	return garboCharsRe.ReplaceAllString(elem, "")
}

func movePartOne(stacks []Deque, crates int, fromStack int, toStack int) []Deque {

	for ; crates > 0; crates-- {
		if !stacks[fromStack].isEmpty() {
			pop := stacks[fromStack].removeLast()
			stacks[toStack].pushBack(pop)
		}
	}

	return stacks
}

func prettyPrint(stacks []Deque) {
	for index, stack := range stacks {
		fmt.Printf("Stack %d %s\n", index, stack.elements)
	}
}

func movePartTwo(stacks []Deque, crates int, fromStack int, toStack int) []Deque {

	length := len(stacks[fromStack].elements)
	cratesToMove := stacks[fromStack].elements[length-crates:]

	stacks[fromStack].elements = stacks[fromStack].elements[:length-crates]
	stacks[toStack].elements = append(stacks[toStack].elements, cratesToMove...)

	return stacks
}

func main() {

	stacks := make([]Deque, 9)

	inputFile := common.ReadTestFile("day5_input.txt")
	fileScanner := bufio.NewScanner(inputFile)
	fileScanner.Split(bufio.ScanLines)

	whitespaceRe := regexp.MustCompile("\\s{4}|\\s{1}")
	digitsRe := regexp.MustCompile("\\d")
	moveRe := regexp.MustCompile("move (\\d+) from (\\d+) to (\\d+)")

	isPartOne := false

	for fileScanner.Scan() {
		line := fileScanner.Text()

		hasDigits := digitsRe.FindString(line) != ""

		if !hasDigits {

			var groups []string = whitespaceRe.Split(line, -1)

			for index, elem := range groups {
				if elem != "" {
					stacks[index].pushFront(sanitise(elem))
				}
			}
		} else {
			matchingGroups := moveRe.FindStringSubmatch(line)

			if len(matchingGroups) == 4 {
				crates, _ := strconv.Atoi(matchingGroups[1])
				fromStack, _ := strconv.Atoi(matchingGroups[2])
				toStack, _ := strconv.Atoi(matchingGroups[3])

				if isPartOne {
					stacks = movePartOne(stacks, crates, fromStack-1, toStack-1)
				} else {
					stacks = movePartTwo(stacks, crates, fromStack-1, toStack-1)
				}
			}
		}
	}

	prettyPrint(stacks)

	for _, stack := range stacks {
		if !stack.isEmpty() {
			fmt.Printf("%s", stack.top())
		}
	}
}
