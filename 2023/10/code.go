package main

import (
	"aoc-in-go/pkg/util"
	"fmt"
)

func main() {
	//aoc.Harness(run)
	util.Run(run, "2023/10/input-user.txt", false)
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

var pipes = map[rune]*pipe{
	'|': {dir(0, 1), dir(0, -1)},
	'-': {dir(-1, 0), dir(1, 0)},
	'L': {dir(0, 1), dir(1, 0)},
	'J': {dir(-1, 0), dir(0, 1)},
	'7': {dir(-1, 0), dir(0, -1)},
	'F': {dir(1, 0), dir(0, -1)},
}

func connected(one cell, two cell) bool {
	onePipe, twoPipe := pipes[one.char], pipes[two.char]
	delta := dirIJ(two.i-one.i, two.j-one.j)
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

func midpointDistance(grid [][]rune, start cell) int {
	neighbors := findNeighbors(grid, start)
	previous, current := start, neighbors[0]
	distance := 1
	for current.char != 'S' {
		next := nextCell(grid, previous, current)
		previous = current
		current = next
		distance++
	}
	return distance / 2
}

func run(part2 bool, input string) any {
	if part2 {
		return "not implemented"
	}

	var grid [][]rune
	for _, line := range util.Lines(input) {
		grid = append(grid, []rune(line))
	}

	for i, row := range grid {
		for j, ch := range row {
			if ch == 'S' {
				start := cell{i, j, ch}
				return midpointDistance(grid, start)
			}
		}
	}
	panic("how'd we get here?!")
}
