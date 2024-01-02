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

var digPlanEntryRegex = regexp.MustCompile("([A-Z]) ([0-9]) .+")

func run(part2 bool, input string) any {
	if part2 {
		return "not implemented"
	}

	currCoord := Coord{0, 0}
	digPath := []Coord{currCoord}
	maxI, maxJ := 0, 0
	for _, line := range util.Lines(input) {
		digPlanParts := digPlanEntryRegex.FindStringSubmatch(line)
		direction, distanceStr := digPlanParts[1], digPlanParts[2]
		distance, err := strconv.Atoi(distanceStr)
		if err != nil {
			panic(err)
		}
		currCoord = currCoord.Move(Directions[direction], distance)
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
	for _, coord := range digPath {
		grid[coord.i][coord.j] = '#'
	}
	//fmt.Println(grid)

	return doubleArea / 2
}
