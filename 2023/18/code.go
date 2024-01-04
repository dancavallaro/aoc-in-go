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
	i, j             int
	prevDir, nextDir *grids.Direction
}

func (c Coord) String() string {
	return fmt.Sprintf("[%d, %d %v %v]", c.i, c.j, c.prevDir, c.nextDir)
}

func (c Coord) Midpoint(c2 Coord) Coord {
	i := (c.i + c2.i) / 2
	j := (c.j + c2.j) / 2
	return Coord{i, j, nil, c.nextDir}
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

	var cornerCoords []Coord
	validCoords := map[Coord]Coord{}
	grid := grids.NewWithFill(maxI-minI+1, maxJ-minJ+1, '.')
	for _, coord := range pathCoords {
		validCoords[Coord{coord.i, coord.j, nil, nil}] = coord
		coord.i -= minI
		coord.j -= minJ
		fillChar := '#'

		if coord.prevDir != nil && coord.nextDir != nil {
			cornerCoords = append(cornerCoords, coord)
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
	fmt.Println(grid)

	for i := 0; i < len(cornerCoords)-1; i++ {
		// Find a horizontal normal because I'm lazy
		midpoint := cornerCoords[i].Midpoint(cornerCoords[i+1])
		fmt.Println(midpoint)
		normal, invNormal := midpoint.nextDir.Left(), midpoint.nextDir.Right()
		fmt.Println(normal, invNormal)
		if normal.DeltaI != 0 {
			continue
		}
		fmt.Println(countCrossings(midpoint, normal, validCoords, grid), countCrossings(midpoint, invNormal, validCoords, grid))
		break
	}

	// TODO: then go in both directions, and figure out which side is the outside
	// TODO: then circumnavigate the path, recording the vertices of the real exterior boundary

	return 42
}

func countCrossings(start Coord, normal grids.Direction, validCoords map[Coord]Coord, grid grids.Grid) int {
	crossings := 0
	var crossedCorner *rune
	for delta := 1; delta <= 10; delta++ { // TODO: 10 -> 1000000
		nextPoint := start.Move(normal, delta)
		var ok bool
		var nextCoord Coord
		if nextCoord, ok = validCoords[Coord{nextPoint.i, nextPoint.j, nil, nil}]; !ok {
			continue
		}
		char := grid[nextCoord.i][nextCoord.j]

		if char == '|' {
			crossings++
		} else if char != '-' {
			if crossedCorner == nil {
				crossedCorner = &char
			} else {
				if (*crossedCorner == 'L' && char == '7') || (*crossedCorner == '7' && char == 'L') || (*crossedCorner == 'F' && char == 'J') || (*crossedCorner == 'J' && char == 'F') {
					crossings++
					crossedCorner = nil
				}
			}
		}
	}
	return crossings
}
