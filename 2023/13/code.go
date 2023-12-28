package main

import (
	"aoc-in-go/pkg/util"
	"fmt"
)

func main() {
	//aoc.Harness(run)
	util.Run(run, "2023/13/input-user.txt", false)
}

type pattern [][]rune

func (p pattern) encode() []int64 {
	var result []int64
	for _, line := range p {
		result = append(result, encode(line))
	}
	return result
}

func (p pattern) invert() pattern {
	inverted := make(pattern, len(p[0]))
	for i := 0; i < len(p[0]); i++ {
		inverted[i] = make([]rune, len(p))
		for j := 0; j < len(p); j++ {
			inverted[i][j] = p[j][i]
		}
	}
	return inverted
}

func (p pattern) summarize() int {
	lengthRows, numBeforeRows := longestPalindrome(p.encode())
	lengthCols, numBeforeCols := longestPalindrome(p.invert().encode())

	if lengthRows > lengthCols {
		return 100 * numBeforeRows
	} else {
		return numBeforeCols
	}
}

func encode(line []rune) int64 {
	result := int64(0)
	for _, cell := range line {
		var digit int64 = 0
		if cell == '#' {
			digit = 1
		}
		result = (result << 1) + digit
	}
	return result
}

func isPalindrome(code []int64) bool {
	if len(code)%2 == 1 {
		// These palindromes will all have even length
		return false
	}

	for i, j := 0, len(code)-1; i < j; i, j = i+1, j-1 {
		if code[i] != code[j] {
			return false
		}
	}
	return true
}

func longestPalindrome(code []int64) (int, int) {
	numBeforeStart := 0
	longestFromStart := 0
	for i := len(code); i > 0; i-- {
		if isPalindrome(code[0:i]) {
			longestFromStart = i
			numBeforeStart = i / 2
			break
		}
	}

	numBeforeEnd := 0
	longestFromEnd := 0
	for i := 0; i < len(code); i++ {
		if isPalindrome(code[i:]) {
			longestFromEnd = len(code) - i
			numBeforeEnd = (len(code) + i) / 2
			break
		}
	}

	if longestFromStart > longestFromEnd {
		return longestFromStart, numBeforeStart
	} else {
		return longestFromEnd, numBeforeEnd
	}
}

func run(part2 bool, input string) any {
	if part2 {
		return "not implemented"
	}

	var patterns []pattern
	var currentPattern pattern
	for _, line := range util.AllLines(input) {
		if len(line) == 0 {
			patterns = append(patterns, currentPattern)
			currentPattern = nil
		} else {
			currentPattern = append(currentPattern, []rune(line))
		}
	}

	totalSummary := 0
	for _, pat := range patterns {
		for _, line := range pat {
			fmt.Println(string(line))
		}

		fmt.Println("Rows: ", pat.encode())
		length, numBefore := longestPalindrome(pat.encode())
		fmt.Printf("Longest palindrome: %d (%d rows/cols before)\n", length, numBefore)
		fmt.Println("Columns: ", pat.invert().encode())
		length, numBefore = longestPalindrome(pat.invert().encode())
		fmt.Printf("Longest palindrome: %d (%d rows/cols before)\n", length, numBefore)
		totalSummary += pat.summarize()
		fmt.Println()
	}

	return totalSummary
}
