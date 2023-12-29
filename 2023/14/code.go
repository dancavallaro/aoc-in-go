package main

import (
	"aoc-in-go/pkg/util"
	"fmt"
	"strings"
)

func main() {
	//aoc.Harness(run)
	util.Run(run, "2023/14/input-user.txt", false)
}

type Grid [][]rune

func (g Grid) String() string {
	var sb strings.Builder
	for _, row := range g {
		sb.WriteString(string(row))
		sb.WriteRune('\n')
	}
	return sb.String()
}

func (g Grid) TiltNorth() {
	for i, row := range g {
		for j, ch := range row {
			if ch != 'O' {
				continue
			}

			var newI int
			for newI = i - 1; newI >= 0; newI-- {
				if g[newI][j] != '.' {
					break
				}
			}
			newI++
			g[i][j] = '.'
			g[newI][j] = 'O'
		}
	}
}

func (g Grid) TotalLoad() int {
	load := 0
	for i, row := range g {
		for _, ch := range row {
			if ch == 'O' {
				load += len(g) - i
			}
		}
	}
	return load
}

func run(part2 bool, input string) any {
	if part2 {
		return "not implemented"
	}

	lines := util.Lines(input)
	var grid Grid = make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}

	fmt.Println("Before:")
	fmt.Println(grid)
	fmt.Println()
	grid.TiltNorth()
	fmt.Println("After:")
	fmt.Println(grid)

	return grid.TotalLoad()
}
