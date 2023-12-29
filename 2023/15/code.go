package main

import (
	"aoc-in-go/pkg/util"
	"strings"
)

func main() {
	//aoc.Harness(run)
	util.Run(run, "2023/15/input-user.txt", false)
}

func Hash(input string) int {
	hash := 0
	for _, ch := range []rune(input) {
		hash += int(ch)
		hash *= 17
		hash %= 256
	}
	return hash
}

func run(part2 bool, input string) any {
	if part2 {
		return "not implemented"
	}

	input = util.Lines(input)[0]
	steps := strings.Split(input, ",")
	total := 0
	for _, step := range steps {
		total += Hash(step)
	}

	return total
}
