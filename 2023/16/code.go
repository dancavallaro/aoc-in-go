package main

import (
	"aoc-in-go/pkg/util"
	"fmt"
	"strings"
)

func main() {
	//aoc.Harness(run)
	util.Run(run, "2023/16/input-user.txt", true)
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

func (g Grid) Copy() Grid {
	newG := make(Grid, len(g))
	for i, row := range g {
		newG[i] = make([]rune, len(row))
		copy(newG[i], row)
	}
	return newG
}

func (g Grid) CountSymbols(sym rune) int {
	count := 0
	for _, row := range g {
		for _, ch := range row {
			if ch == sym {
				count++
			}
		}
	}
	return count
}

type Direction struct {
	deltaI, deltaJ int
}

var (
	North = Direction{-1, 0}
	East  = Direction{0, 1}
	South = Direction{1, 0}
	West  = Direction{0, -1}
)

func (dir Direction) Symbol() rune {
	if dir == North {
		return '^'
	} else if dir == East {
		return '>'
	} else if dir == South {
		return 'v'
	} else if dir == West {
		return '<'
	}
	return '#'
}

type Beam struct {
	i, j int
	dir  Direction
	Done bool
}

type Beams []*Beam

func (bs Beams) String() string {
	var sb strings.Builder
	sb.WriteString("[ ")
	for _, beam := range bs {
		sb.WriteString(fmt.Sprintf("%v ", *beam))
	}
	sb.WriteString("]")
	return sb.String()
}

func (b *Beam) Step(grid Grid, energized Grid) {
	newI, newJ := b.i+b.dir.deltaI, b.j+b.dir.deltaJ
	if newI < 0 || newI >= len(grid) || newJ < 0 || newJ >= len(grid[0]) {
		b.Done = true
	} else {
		b.i, b.j = newI, newJ
		energized[newI][newJ] = b.dir.Symbol()
	}
}

func NextMove(beam *Beam, grid Grid, energized Grid) *Beam {
	if beam.Done {
		return nil
	}
	tile := grid[beam.i][beam.j]
	var splitBeam *Beam

	if tile == '-' && beam.dir.deltaJ == 0 {
		beam.dir = Direction{0, 1}
		splitBeam = &Beam{i: beam.i, j: beam.j, dir: Direction{0, -1}, Done: false}
	} else if tile == '|' && beam.dir.deltaI == 0 {
		beam.dir = Direction{1, 0}
		splitBeam = &Beam{i: beam.i, j: beam.j, dir: Direction{-1, 0}, Done: false}
	} else if tile == '/' {
		beam.dir = Direction{-beam.dir.deltaJ, -beam.dir.deltaI}
	} else if tile == '\\' {
		beam.dir = Direction{beam.dir.deltaJ, beam.dir.deltaI}
	}

	beam.Step(grid, energized)
	if splitBeam != nil {
		splitBeam.Step(grid, energized)
	}
	return splitBeam
}

func countEnergized(grid Grid, initialBeam Beam) int {
	energized := make(Grid, len(grid))
	for i, row := range grid {
		energized[i] = make([]rune, len(row))
		for j := range row {
			energized[i][j] = '.'
		}
	}
	energized[initialBeam.i][initialBeam.j] = initialBeam.dir.Symbol()

	var beams Beams = []*Beam{&initialBeam}
	seen := make(map[string]bool)
	for {
		for _, beam := range beams {
			if !beam.Done {
				splitBeam := NextMove(beam, grid, energized)
				//fmt.Println("Split beam: ", splitBeam)
				if splitBeam != nil {
					beams = append(beams, splitBeam)
				}
			}
		}
		if seen[energized.String()] {
			break
		}
		seen[energized.String()] = true
	}

	return (len(grid) * len(grid[0])) - energized.CountSymbols('.')
}

func run(part2 bool, input string) any {
	if !part2 {
		return "not implemented"
	}

	lines := util.Lines(input)
	var grid Grid = make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}

	var initialBeams []Beam
	for i := range grid {
		initialBeams = append(initialBeams, Beam{i, 0, East, false})
		initialBeams = append(initialBeams, Beam{i, len(grid[i]) - 1, West, false})
	}
	for j := range grid[0] {
		initialBeams = append(initialBeams, Beam{0, j, South, false})
		initialBeams = append(initialBeams, Beam{len(grid) - 1, j, North, false})
	}

	maxEnergized := 0
	for _, initialBeam := range initialBeams {
		energized := countEnergized(grid, initialBeam)
		maxEnergized = max(maxEnergized, energized)
		fmt.Printf("%v -> %d\n", initialBeam, energized)
	}
	return maxEnergized
}
