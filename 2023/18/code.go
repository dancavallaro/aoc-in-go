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
	util.Run(run, "2023/18/input-example.txt", false)
}

type Coord struct {
	i, j             int
	prevDir, nextDir *grids.Direction
}

func (c Coord) Move(direction grids.Direction, distance int) Coord {
	return Coord{c.i + distance*direction.DeltaI, c.j + distance*direction.DeltaJ, nil, nil}
}

var Directions = map[string]grids.Direction{
	"U": grids.North,
	"R": grids.East,
	"D": grids.South,
	"L": grids.West,
}

type DirectionChange struct {
	prevDir, nextDir grids.Direction
}

func (dc DirectionChange) Equal(other DirectionChange) bool {
	return dc.prevDir == other.prevDir && dc.nextDir == other.nextDir
}

func reverse(d grids.Direction) grids.Direction {
	return grids.Direction{-d.DeltaI, -d.DeltaJ}
}

func (dc DirectionChange) Reverse() DirectionChange {
	return DirectionChange{reverse(dc.nextDir), reverse(dc.prevDir)}
}

var digPlanEntryRegex = regexp.MustCompile("([A-Z]) ([0-9]+) .+")

func run(part2 bool, input string) any {
	if part2 {
		return "not implemented"
	}

	currCoord := Coord{0, 0, nil, nil}
	var pathCoords []Coord
	var firstDirection, lastDirection *grids.Direction
	maxI, maxJ, minI, minJ := 0, 0, 0, 0
	for _, line := range util.Lines(input) {
		digPlanParts := digPlanEntryRegex.FindStringSubmatch(line)
		directionStr, distanceStr := digPlanParts[1], digPlanParts[2]
		direction := Directions[directionStr]

		if firstDirection == nil {
			firstDirection = &direction
		}

		if lastDirection != nil {
			//fmt.Printf("Changing direction from %v to %v\n", *lastDirection, direction)
			pathCoords[len(pathCoords)-1].prevDir = lastDirection
			pathCoords[len(pathCoords)-1].nextDir = &direction
		}

		distance, err := strconv.Atoi(distanceStr)
		if err != nil {
			panic(err)
		}

		for delta := 1; delta <= distance; delta++ {
			nextCoord := currCoord.Move(direction, delta)

			if nextCoord.i == 0 && nextCoord.j == 0 {
				nextCoord.prevDir = &direction
				nextCoord.nextDir = firstDirection
			}
			if delta < distance {
				nextCoord.nextDir = &direction
			}

			pathCoords = append(pathCoords, nextCoord)
		}

		currCoord = currCoord.Move(direction, distance)
		maxI, maxJ, minI, minJ = max(maxI, currCoord.i), max(maxJ, currCoord.j), min(minI, currCoord.i), min(minJ, currCoord.j)
		lastDirection = &direction
	}

	cornerSymbols := map[DirectionChange]rune{
		{grids.Direction{-1, 0}, grids.Direction{0, 1}}:  'F',
		{grids.Direction{1, 0}, grids.Direction{0, 1}}:   'L',
		{grids.Direction{1, 0}, grids.Direction{0, -1}}:  'J',
		{grids.Direction{-1, 0}, grids.Direction{0, -1}}: '7',
	}

	grid := grids.NewWithFill(maxI-minI+1, maxJ-minJ+1, '.')
	for _, coord := range pathCoords {
		coord.i -= minI
		coord.j -= minJ
		fillChar := '#'

		if coord.prevDir != nil && coord.nextDir != nil {
			fmt.Printf("Changing direction from %v to %v\n", coord.prevDir, coord.nextDir)
			dc := DirectionChange{*coord.prevDir, *coord.nextDir}
			if sym, ok := cornerSymbols[dc]; ok {
				fillChar = sym
			}
			if sym, ok := cornerSymbols[dc.Reverse()]; ok {
				fillChar = sym
			}
		} else if coord.nextDir != nil {
			if coord.nextDir.DeltaI == 0 {
				fillChar = '-'
			} else {
				fillChar = '|'
			}
		}

		grid[coord.i][coord.j] = fillChar
	}
	//markInterior(grid)
	fmt.Println(grid)

	// TODO: then find the midpoint between the first two vertices
	// TODO: then find the line perpendicular to that
	// TODO: then go in both directions, and figure out which side is the outside
	// TODO: then circumnavigate the path, recording the vertices of the real exterior boundary

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
