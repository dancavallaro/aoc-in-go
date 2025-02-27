package main

import (
	"aoc-in-go/pkg/grids"
	"aoc-in-go/pkg/util"
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

var digPlanEntryRegex = regexp.MustCompile("([A-Z]) ([0-9]+) .+")

func run(part2 bool, input string) any {
	if part2 {
		return "not implemented"
	}

	currCoord := Coord{0, 0}
	var pathCoords []Coord
	maxI, maxJ, minI, minJ := 0, 0, 0, 0
	for _, line := range util.Lines(input) {
		digPlanParts := digPlanEntryRegex.FindStringSubmatch(line)
		directionStr, distanceStr := digPlanParts[1], digPlanParts[2]
		direction := Directions[directionStr]
		distance, err := strconv.Atoi(distanceStr)
		if err != nil {
			panic(err)
		}

		for delta := 1; delta <= distance; delta++ {
			nextCoord := currCoord.Move(direction, delta)
			pathCoords = append(pathCoords, nextCoord)
		}

		currCoord = currCoord.Move(direction, distance)
		maxI, maxJ, minI, minJ = max(maxI, currCoord.i), max(maxJ, currCoord.j), min(minI, currCoord.i), min(minJ, currCoord.j)
	}

	grid := grids.NewWithFill(maxI-minI+1, maxJ-minJ+1, '.')
	for _, coord := range pathCoords {
		coord.i -= minI
		coord.j -= minJ
		grid[coord.i][coord.j] = '#'
	}
	markInterior(grid)

	totalArea := (maxI - minI + 1) * (maxJ - minJ + 1)
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
