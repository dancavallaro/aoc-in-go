package main

import (
	"aoc-in-go/pkg/util"
	"math"
	"strconv"
	"strings"
)

func main() {
	//aoc.Harness(run)
	util.Run(run, "2023/04/input-user.txt", false)
}

type card struct {
	id             int
	winningNumbers map[int]bool
	yourNumbers    map[int]bool
}

func makeSet[T comparable](elems []T) map[T]bool {
	elemMap := map[T]bool{}
	for _, elem := range elems {
		elemMap[elem] = true
	}
	return elemMap
}

func newCard(id int, winningNumbers []int, yourNumbers []int) card {
	return card{id, makeSet(winningNumbers), makeSet(yourNumbers)}
}

func (c card) points() int {
	matchingNumbers := 0
	for num := range c.winningNumbers {
		if c.yourNumbers[num] {
			matchingNumbers++
		}
	}

	if matchingNumbers == 0 {
		return 0
	} else {
		return int(math.Pow(2, float64(matchingNumbers-1)))
	}
}

func parseNumbers(line string) []int {
	var numbers []int
	for _, numberStr := range strings.Split(line, " ") {
		if numberStr == "" {
			continue
		}
		number, err := strconv.Atoi(strings.TrimSpace(numberStr))
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, number)
	}
	return numbers
}

func parseCard(line string) card {
	colonIdx := strings.IndexRune(line, ':')
	cardId, err := strconv.Atoi(strings.TrimSpace(line[5:colonIdx])) // Lines all begin with "Card X: "
	if err != nil {
		panic(err)
	}
	restOfLine := line[colonIdx+2:]
	numbers := strings.Split(restOfLine, "|")

	return newCard(cardId, parseNumbers(numbers[0]), parseNumbers(numbers[1]))
}

func run(part2 bool, input string) any {
	if part2 {
		return "not implemented"
	}

	totalPoints := 0
	for _, line := range util.Lines(input) {
		card := parseCard(line)
		totalPoints += card.points()
	}

	return totalPoints
}
