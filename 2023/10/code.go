package main

import (
	"aoc-in-go/pkg/util"
	"fmt"
)

func main() {
	//aoc.Harness(run)
	util.Run(run, "2023/10/input-user.txt", true)
}

type direction struct {
	deltaX, deltaY int
}

func (d direction) deltaI() int {
	return -d.deltaY
}

func (d direction) deltaJ() int {
	return d.deltaX
}

func (d direction) negative() direction {
	return dir(-d.deltaX, -d.deltaY)
}

func dir(deltaX int, deltaY int) direction {
	return direction{deltaX, deltaY}
}

func dirIJ(deltaI int, deltaJ int) direction {
	return direction{deltaJ, -deltaI}
}

func directionBetween(one cell, two cell) direction {
	return dirIJ(two.i-one.i, two.j-one.j)
}

type pipe struct {
	previous, next direction
}

type cell struct {
	i, j int
	char rune
}

func (c cell) plus(grid [][]rune, d direction) cell {
	i, j := c.i+d.deltaI(), c.j+d.deltaJ()
	return cell{i, j, grid[i][j]}
}

func (c cell) String() string {
	return fmt.Sprintf("[i=%d, j=%d, char=%c]", c.i, c.j, c.char)
}

func (d direction) String() string {
	return fmt.Sprintf("[%d %d]", d.deltaX, d.deltaY)
}

var pipes = map[rune]pipe{
	'|': {dir(0, 1), dir(0, -1)},
	'-': {dir(-1, 0), dir(1, 0)},
	'L': {dir(0, 1), dir(1, 0)},
	'J': {dir(-1, 0), dir(0, 1)},
	'7': {dir(-1, 0), dir(0, -1)},
	'F': {dir(1, 0), dir(0, -1)},
}
var pipeDirs = util.InvertMap(pipes)

func connected(one cell, two cell) bool {
	onePipe, twoPipe := pipes[one.char], pipes[two.char]
	delta := directionBetween(one, two)
	if one.char == 'S' {
		return twoPipe.previous == delta.negative() || twoPipe.next == delta.negative()
	} else if two.char == 'S' {
		return onePipe.previous == delta || onePipe.next == delta
	}
	return (onePipe.previous == delta || onePipe.next == delta) &&
		(twoPipe.previous == delta.negative() || twoPipe.next == delta.negative())
}

func findNeighbors(grid [][]rune, c cell) []cell {
	var neighbors []cell
	moves := []direction{dir(1, 0), dir(0, -1)}
	for _, move := range moves {
		newI, newJ := c.i+move.deltaI(), c.j+move.deltaJ()
		if newI < 0 || newI >= len(grid) || newJ < 0 || newJ >= len(grid[0]) {
			continue
		}
		neighbor := cell{newI, newJ, grid[newI][newJ]}
		if connected(c, neighbor) {
			neighbors = append(neighbors, neighbor)
		}
	}
	return neighbors
}

func nextCell(grid [][]rune, previous cell, start cell) cell {
	startPipe := pipes[start.char]
	if start.plus(grid, startPipe.previous) == previous {
		return start.plus(grid, startPipe.next)
	} else {
		return start.plus(grid, startPipe.previous)
	}
}

func findPipe(grid [][]rune, start cell) []cell {
	neighbors := findNeighbors(grid, start)
	pipeCells := []cell{start}
	previous, current := start, neighbors[0]
	distance := 1
	for current.char != 'S' {
		pipeCells = append(pipeCells, current)
		next := nextCell(grid, previous, current)
		previous = current
		current = next
		distance++
	}
	return pipeCells
}

func findStartPipe(pipeCells []cell) rune {
	start, next, last := pipeCells[0], pipeCells[1], pipeCells[len(pipeCells)-1]
	lastDir, nextDir := directionBetween(start, last), directionBetween(start, next)
	startPipe := pipeDirs[pipe{lastDir, nextDir}]
	if startPipe == 0 {
		startPipe = pipeDirs[pipe{nextDir, lastDir}]
	}
	return startPipe
}

func countCrossings(startI, startJ int, grid [][]rune, cellIsPipe map[cell]bool) int {
	crossings := 0
	incomingDirection := 0
	for j := startJ + 1; j < len(grid[startI]); j++ {
		char := grid[startI][j]
		c := cell{startI, j, char}

		if !cellIsPipe[c] {
			continue
		}

		if char == '|' {
			crossings++
		} else if char != '-' {
			if char == 'L' {
				incomingDirection = 1
			} else if char == 'F' {
				incomingDirection = -1
			} else if char == 'J' {
				if incomingDirection == -1 {
					crossings++
					incomingDirection = 0
				}
			} else if char == '7' {
				if incomingDirection == 1 {
					crossings++
					incomingDirection = 0
				}
			}
		}
	}
	return crossings
}

func countEnclosedTiles(grid [][]rune, cellIsPipe map[cell]bool) int {
	enclosedTiles := 0
	for i, row := range grid {
		for j := range row {
			c := cell{i, j, grid[i][j]}
			if cellIsPipe[c] {
				continue
			}
			crossings := countCrossings(i, j, grid, cellIsPipe)

			if crossings%2 == 1 {
				enclosedTiles++
			}
		}
	}
	return enclosedTiles
}

func run(part2 bool, input string) any {
	if !part2 {
		return "not implemented"
	}

	var grid [][]rune
	for _, line := range util.Lines(input) {
		grid = append(grid, []rune(line))
	}

	var pipeCells []cell
	for i, row := range grid {
		for j, ch := range row {
			if ch == 'S' {
				start := cell{i, j, ch}
				pipeCells = findPipe(grid, start)
			}
		}
	}
	if pipeCells == nil {
		panic("how'd we get here?!")
	}

	startPipe := findStartPipe(pipeCells)
	pipeCells[0].char = startPipe
	grid[pipeCells[0].i][pipeCells[0].j] = startPipe

	cellIsPipe := map[cell]bool{}
	for _, c := range pipeCells {
		cellIsPipe[c] = true
	}

	return countEnclosedTiles(grid, cellIsPipe)
}
