package main

import (
	"aoc-in-go/pkg/grids"
	"aoc-in-go/pkg/util"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	//aoc.Harness(run)
	util.Run(run, "2023/18/input-user.txt", false)
}

type Coord struct {
	i, j int
}

func (c Coord) Move(direction grids.Direction, distance int) Coord {
	return Coord{c.i + distance*direction.DeltaI, c.j + distance*direction.DeltaJ}
}

var Directions = map[string]grids.Direction{
	"U": grids.North,
	"R": grids.East,
	"D": grids.South,
	"L": grids.West,
}

func determinant(one Coord, two Coord) int {
	return one.j*two.i - one.i*two.j
}

var digPlanEntryRegex = regexp.MustCompile("([A-Z]) ([0-9]+) .+")

func run(part2 bool, input string) any {
	if part2 {
		return "not implemented"
	}

	currCoord := Coord{0, 0}
	digPath := []Coord{currCoord}
	pathLines := map[Coord]rune{}
	maxI, maxJ := 0, 0
	for _, line := range util.Lines(input) {
		digPlanParts := digPlanEntryRegex.FindStringSubmatch(line)
		directionStr, distanceStr := digPlanParts[1], digPlanParts[2]
		direction := Directions[directionStr]
		distance, err := strconv.Atoi(distanceStr)
		if err != nil {
			panic(err)
		}

		pathLine := 'x' // TODO
		if direction == grids.East || direction == grids.West {
			pathLine = '-'
		} else if direction == grids.South || direction == grids.North {
			pathLine = '|'
		}
		for delta := 1; delta <= distance; delta++ {
			nextCoord := currCoord.Move(direction, delta)
			pathLines[nextCoord] = pathLine
		}

		currCoord = currCoord.Move(direction, distance)
		digPath = append(digPath, currCoord)
		maxI, maxJ = max(maxI, currCoord.i), max(maxJ, currCoord.j)
	}
	fmt.Println(digPath)

	doubleArea := 0
	for i, pathNode := range digPath {
		if i >= len(digPath)-1 {
			continue
		}
		doubleArea += determinant(pathNode, digPath[i+1])
	}

	grid := grids.NewWithFill(maxI+1, maxJ+1, '.')
	for coord, _ := range pathLines {
		grid[coord.i][coord.j] = '#' //char
	}
	fmt.Println(grid)
	markInterior(grid)
	fmt.Println(grid)

	totalArea := (maxI + 1) * (maxJ + 1)
	for _, row := range grid {
		for _, cell := range row {
			if cell == '*' {
				totalArea--
			}
		}
	}
	return totalArea
}

func markInterior(grid grids.Grid) {
	for i := range grid {
		floodFill(grid, i, 0)
		floodFill(grid, i, len(grid[i])-1)
	}
	for j := range grid[0] {
		floodFill(grid, 0, j)
		floodFill(grid, len(grid)-1, j)
	}
}

func floodFill(grid grids.Grid, i, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
		return
	}
	if grid[i][j] != '.' {
		return
	}

	grid[i][j] = '*'

	floodFill(grid, i, j+1)
	floodFill(grid, i+1, j)
	floodFill(grid, i, j-1)
	floodFill(grid, i-1, j)
}
