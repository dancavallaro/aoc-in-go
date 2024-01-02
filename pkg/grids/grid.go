package grids

import "strings"

type Grid [][]rune

func NewWithFill(height, width int, fillChar rune) Grid {
	grid := make(Grid, height)
	for i := range grid {
		grid[i] = make([]rune, width)
		for j := range grid[i] {
			grid[i][j] = fillChar
		}
	}
	return grid
}

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

type Direction struct {
	DeltaI, DeltaJ int
}

func (dir Direction) Left() Direction {
	return Direction{-dir.DeltaJ, dir.DeltaI}
}

func (dir Direction) Right() Direction {
	return Direction{dir.DeltaJ, -dir.DeltaI}
}

var (
	North = Direction{-1, 0}
	East  = Direction{0, 1}
	South = Direction{1, 0}
	West  = Direction{0, -1}
)
