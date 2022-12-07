package main

import (
	common "aoc23/main"
	"bufio"
	"fmt"
	"regexp"
	"strconv"
)

type Stack struct {
	elements []string
}

func (s *Stack) pop() string {
	lastElem := s.elements[0]
	s.elements = s.elements[1:len(s.elements)]
	return lastElem
}

func (s *Stack) prepend(elem string) {
	s.elements = append([]string{elem}, s.elements...)
}

func (s *Stack) push(elem string) {
	s.elements = append(s.elements, elem)
}

func (s *Stack) top() string {
	return s.elements[0]
}

func (s *Stack) isEmpty() bool {
	return len(s.elements) == 0
}

func sanitise(elem string) string {
	garboCharsRe := regexp.MustCompile("\\[|\\]")
	return garboCharsRe.ReplaceAllString(elem, "")
}

func move(stacks []Stack, crates int, fromStack int, toStack int) []Stack {

	fmt.Printf("Moving %d from %d to %d\n", crates, fromStack, toStack)
	fmt.Println(stacks[fromStack-1])

	for ; crates > 0; crates-- {
		pop := stacks[fromStack-1].pop()
		fmt.Printf("pop %s\n", pop)

		stacks[toStack-1].prepend(pop)
	}

	prettyPrint(stacks)

	return stacks
}

func prettyPrint(stacks []Stack) {
	for _, stack := range stacks {
		for _, elem := range stack.elements {
			fmt.Printf("%s,", elem)
		}
		fmt.Println()
	}
}

func main() {

	stacks := make([]Stack, 20)

	inputFile := common.ReadTestFile("day5_input.txt")
	fileScanner := bufio.NewScanner(inputFile)
	fileScanner.Split(bufio.ScanLines)

	whitespaceRe := regexp.MustCompile("\\s{4}|\\s{1}")
	digitsRe := regexp.MustCompile("\\d")
	moveRe := regexp.MustCompile("move (\\d) from (\\d) to (\\d)")

	for fileScanner.Scan() {
		line := fileScanner.Text()

		hasDigits := digitsRe.FindString(line) != ""

		if !hasDigits {

			var groups []string = whitespaceRe.Split(line, -1)

			for index, elem := range groups {
				if elem != "" {
					stacks[index].push(sanitise(elem))
				}
			}
		} else {
			matchingGroups := moveRe.FindStringSubmatch(line)

			if len(matchingGroups) == 4 {
				crates, _ := strconv.Atoi(matchingGroups[1])
				fromStack, _ := strconv.Atoi(matchingGroups[2])
				toStack, _ := strconv.Atoi(matchingGroups[3])

				stacks = move(stacks, crates, fromStack, toStack)
			}
		}
	}

	prettyPrint(stacks)

	for _, stack := range stacks {
		if len(stack.elements) > 0 {
			fmt.Printf("%s", stack.top())
		}
	}
}
