package main

import (
	"aoc-in-go/pkg/util"
	"fmt"
	"strings"
)

func main() {
	//aoc.Harness(run)
	util.Run(run, "2023/14/input-user.txt", true)
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

func (g Grid) copy() Grid {
	newG := make(Grid, len(g))
	for i, row := range g {
		newG[i] = make([]rune, len(row))
		copy(newG[i], row)
	}
	return newG
}

func (g Grid) invert() Grid {
	inverted := make(Grid, len(g[0]))
	for i := 0; i < len(g[0]); i++ {
		inverted[i] = make([]rune, len(g))
		for j := 0; j < len(g); j++ {
			inverted[i][j] = g[j][i]
		}
	}
	return inverted
}

func (g Grid) flip() Grid {
	newG := g.copy()
	for i, j := 0, len(g)-1; i < j; i, j = i+1, j-1 {
		newG[i] = g[j]
		newG[j] = g[i]
	}
	return newG
}

func (g Grid) RunCycle() Grid {
	// North
	g = g.TiltNorth()
	// West
	g = g.invert().TiltNorth().invert()
	// South
	g = g.flip().TiltNorth().flip()
	// East
	g = g.invert().flip().TiltNorth().flip().invert()

	return g
}

func (g Grid) TiltNorth() Grid {
	g = g.copy()
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
	return g
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

func (g Grid) Key() string {
	var sb strings.Builder
	for _, row := range g {
		sb.WriteString(string(row))
	}
	return sb.String()
}

const totalCycles = 1000000000

func run(part2 bool, input string) any {
	if !part2 {
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

	seenGrids := make(map[string]int)
	var cycleStart, nextCycleStart int
	for cycle := 0; cycle < totalCycles; cycle++ {
		key := grid.Key()
		if lastCycle, ok := seenGrids[key]; ok {
			fmt.Printf("last saw this grid on cycle %d (now cycle %d)\n", lastCycle, cycle)
			cycleStart, nextCycleStart = lastCycle, cycle
			break
		} else {
			seenGrids[key] = cycle
		}

		grid = grid.RunCycle()
	}

	cycleLength := nextCycleStart - cycleStart
	cyclesLeft := totalCycles - nextCycleStart
	for i := 0; i < cyclesLeft%cycleLength; i++ {
		grid = grid.RunCycle()
	}

	return grid.TotalLoad()
}
