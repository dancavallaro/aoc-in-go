package main

import (
	"aoc-in-go/pkg/util"
	"strconv"
	"strings"
)

func main() {
	//aoc.Harness(run)
	util.Run(run, "2023/02/input-user.txt", false)
}

func parseCubeSet(s string) cubeSet {
	cubeStrings := strings.Split(s, ", ")
	red, green, blue := 0, 0, 0

	for _, cubeString := range cubeStrings {
		parts := strings.Split(cubeString, " ")
		numCubes, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}

		if parts[1] == "red" {
			red = numCubes
		} else if parts[1] == "green" {
			green = numCubes
		} else if parts[1] == "blue" {
			blue = numCubes
		}
	}
	return cubeSet{red, green, blue}
}

func parseCubeSets(s string) []cubeSet {
	cubeSetStrings := strings.Split(s, "; ")
	cubeSets := make([]cubeSet, 0)

	for _, cubeSetStr := range cubeSetStrings {
		cubeSets = append(cubeSets, parseCubeSet(cubeSetStr))
	}

	return cubeSets
}

func parseGame(line string) game {
	colonIdx := strings.IndexRune(line, ':')
	gameId, err := strconv.Atoi(line[5:colonIdx]) // Lines all begin with "Game X: "
	if err != nil {
		panic(err)
	}
	restOfLine := line[colonIdx+2:]

	return game{id: gameId, cubes: parseCubeSets(restOfLine)}
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return "not implemented"
	}

	cubeBag := cubeSet{12, 13, 14}
	sumOfMatchingIds := 0

outerLoop:
	for _, line := range util.Lines(input) {
		game := parseGame(line)
		for _, cubes := range game.cubes {
			if !cubeBag.contains(cubes) {
				continue outerLoop
			}
		}
		sumOfMatchingIds += game.id
	}

	return sumOfMatchingIds
}

type cubeSet struct {
	red, green, blue int
}

func (cubes cubeSet) contains(other cubeSet) bool {
	return cubes.red >= other.red && cubes.green >= other.green && cubes.blue >= other.blue
}

type game struct {
	id    int
	cubes []cubeSet
}
