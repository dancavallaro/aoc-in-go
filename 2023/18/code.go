package main

import (
	"aoc-in-go/pkg/util"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	//aoc.Harness(run)
	util.Run(run, "2023/18/input-small.txt", false)
}

type Coord struct {
	i, j int
}

func (c Coord) Move(direction Direction, distance int) Coord {
	return Coord{c.i + distance*direction.deltaI, c.j + distance*direction.deltaJ}
}

type Direction struct {
	deltaI, deltaJ int
}

var Up = Direction{-1, 0}
var Right = Direction{0, 1}
var Down = Direction{1, 0}
var Left = Direction{0, -1}
var Directions = map[string]Direction{
	"U": Up,
	"R": Right,
	"D": Down,
	"L": Left,
}

func determinant(one Coord, two Coord) int {
	return one.j*two.i - one.i*two.j
}

var digPlanEntryRegex = regexp.MustCompile("([A-Z]) ([0-9]) .+")

func run(part2 bool, input string) any {
	if part2 {
		return "not implemented"
	}

	coord := Coord{0, 0}
	digPath := []Coord{coord}
	for _, line := range util.Lines(input) {
		digPlanParts := digPlanEntryRegex.FindStringSubmatch(line)
		direction, distanceStr := digPlanParts[1], digPlanParts[2]
		distance, err := strconv.Atoi(distanceStr)
		if err != nil {
			panic(err)
		}
		coord = coord.Move(Directions[direction], distance)
		digPath = append(digPath, coord)
	}
	fmt.Println(digPath)

	doubleArea := 0
	for i, pathNode := range digPath {
		if i >= len(digPath)-1 {
			continue
		}
		doubleArea += determinant(pathNode, digPath[i+1])
	}

	return doubleArea / 2
}
