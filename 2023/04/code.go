package main

import (
	"aoc-in-go/pkg/util"
	"strconv"
	"strings"
)

func main() {
	//aoc.Harness(run)
	util.Run(run, "2023/04/input-user.txt", true)
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

func (c card) matchingNumbers() int {
	matchingNumbers := 0
	for num := range c.winningNumbers {
		if c.yourNumbers[num] {
			matchingNumbers++
		}
	}
	return matchingNumbers
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
	if !part2 {
		return "not implemented"
	}

	lines := util.Lines(input)
	totalPoints := len(lines)
	matchesPerCard := make([]int, len(lines)+1)
	for _, line := range lines {
		card := parseCard(line)
		matchingNums := card.matchingNumbers()
		matchesPerCard[card.id] = matchingNums
	}

	cardsPerGame := map[int]int{}
	for game := len(lines); game > 0; game-- {
		cardsThisGame := matchesPerCard[game]
		for i := 1; i <= matchesPerCard[game]; i++ {
			cardsThisGame += cardsPerGame[game+i]
		}
		cardsPerGame[game] = cardsThisGame
		totalPoints += cardsThisGame
	}

	return totalPoints
}
