package main

import (
	"aoc-in-go/pkg/util"
)

func main() {
	//aoc.Harness(run)
	util.Run(run, "2023/09/input-user.txt", false)
}

func allZeroes(nums []int) bool {
	for _, num := range nums {
		if num != 0 {
			return false
		}
	}
	return true
}

func findDifferences(history []int) [][]int {
	differences := [][]int{history}
	for !allZeroes(history) {
		var nextDiffs []int
		for i := 0; i < len(history)-1; i++ {
			diff := history[i+1] - history[i]
			nextDiffs = append(nextDiffs, diff)
		}
		history = nextDiffs
		differences = append(differences, history)
	}
	return differences
}

func extrapolateValue(differences [][]int) int {
	val := 0
	for i := len(differences) - 2; i >= 0; i-- {
		row := differences[i]
		val = val + row[len(row)-1]
	}
	return val
}

func forecast(history []int) int {
	differences := findDifferences(history)
	return extrapolateValue(differences)
}

func run(part2 bool, input string) any {
	if part2 {
		return "not implemented"
	}

	sum := 0
	for _, line := range util.Lines(input) {
		history := util.ParseInts(line)
		sum += forecast(history)
	}
	return sum
}
