package main

import (
	"github.com/jpillora/puzzler/harness/aoc"
	"strings"
)

func main() {
	aoc.Harness(run)
}

func firstDigit(s string) int {
	for _, c := range s {
		if c >= '0' && c <= '9' {
			return int(c - '0')
		}
	}
	panic("no digits found!")
}

func reverse(s string) string {
	chars := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

func value(line string) int {
	first := firstDigit(line)
	last := firstDigit(reverse(line))
	return 10*first + last
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return "not implemented"
	}

	lines := strings.Split(input, "\n")
	calibrationValue := 0

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		calibrationValue += value(line)
	}

	return calibrationValue
}
