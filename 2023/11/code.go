package main

import (
	"aoc-in-go/pkg/util"
)

func main() {
	//aoc.Harness(run)
	util.Run(run, "2023/11/input-user.txt", true)
}

type galaxy struct {
	row, col int
}

var emptySpaceSize = 2

func findDistance(galA galaxy, galB galaxy, galRows []bool, galCols []bool) int {
	distance := 0

	startRow, endRow := min(galA.row, galB.row), max(galA.row, galB.row)
	for row := startRow; row < endRow; row++ {
		if galRows[row] {
			distance++
		} else {
			distance += emptySpaceSize
		}
	}

	startCol, endCol := min(galA.col, galB.col), max(galA.col, galB.col)
	for col := startCol; col < endCol; col++ {
		if galCols[col] {
			distance++
		} else {
			distance += emptySpaceSize
		}
	}

	return distance
}

func run(part2 bool, input string) any {
	if part2 {
		emptySpaceSize = 1000000
	}

	lines := util.Lines(input)
	numRows, numCols := len(lines), len(lines[0])
	galaxyRows, galaxyCols := make([]bool, numRows), make([]bool, numCols)
	var galaxies []galaxy

	for row, line := range lines {
		for col, ch := range []rune(line) {
			if ch == '#' {
				galaxies = append(galaxies, galaxy{row, col})
				galaxyRows[row] = true
				galaxyCols[col] = true
			}
		}
	}

	totalDistance := 0
	for i, galA := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			galB := galaxies[j]
			totalDistance += findDistance(galA, galB, galaxyRows, galaxyCols)
		}
	}
	return totalDistance
}
