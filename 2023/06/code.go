package main

import (
	"aoc-in-go/pkg/util"
	"strings"
)

func main() {
	//aoc.Harness(run)
	util.Run(run, "2023/06/input-user.txt", true)
}

func numWinningWays(time int, distance int) int {
	low, high := findSolutionBounds(time, distance)

	return int(high-low) + 1
}

func findSolutionBounds(time int, distance int) (int, int) {
	solutionFunc := func(holdTime int) int {
		return holdTime*holdTime - time*holdTime + distance
	}

	midpoint := time / 2
	lowZero := findZero(solutionFunc, 0, midpoint)
	highZero := time - lowZero

	return lowZero, highZero
}

func findZero(f func(int) int, low int, high int) int {
	if low == high {
		return low
	}

	mid := (low + high) / 2
	if f(mid) < 0 {
		return findZero(f, low, mid)
	} else {
		return findZero(f, mid+1, high)
	}
}

func parseNums(line string) []int {
	numStr := strings.Split(line, ":")[1]
	return util.ParseInts(numStr)
}

func parseBigNum(line string) []int {
	lineWithoutSpace := strings.ReplaceAll(line, " ", "")
	return parseNums(lineWithoutSpace)
}

func run(part2 bool, input string) any {
	lines := util.Lines(input)
	var times, distances []int
	if part2 {
		times = parseBigNum(lines[0])
		distances = parseBigNum(lines[1])
	} else {
		times = parseNums(lines[0])
		distances = parseNums(lines[1])
	}
	result := 1

	for i, time := range times {
		distance := distances[i]
		result *= numWinningWays(time, distance)
	}

	return result
}
