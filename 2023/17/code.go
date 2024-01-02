package main

import (
	"aoc-in-go/pkg/collections"
	"aoc-in-go/pkg/util"
	"math"
	"strings"
)

func main() {
	//aoc.Harness(run)
	util.Run(run, "2023/17/input-user.txt", true)
}

type Grid [][]rune

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
	deltaI, deltaJ int
}

func (dir Direction) Left() Direction {
	return Direction{-dir.deltaJ, dir.deltaI}
}

func (dir Direction) Right() Direction {
	return Direction{dir.deltaJ, -dir.deltaI}
}

var North = Direction{-1, 0}
var East = Direction{0, 1}
var South = Direction{1, 0}
var West = Direction{0, -1}

type PathState struct {
	i, j         int
	movesThisDir int
	dir          Direction
}

func (ps PathState) Move() PathState {
	newI := ps.i + ps.dir.deltaI
	newJ := ps.j + ps.dir.deltaJ
	return PathState{newI, newJ, ps.movesThisDir + 1, ps.dir}
}

func (ps PathState) Left() PathState {
	return PathState{ps.i, ps.j, 0, ps.dir.Left()}
}

func (ps PathState) Right() PathState {
	return PathState{ps.i, ps.j, 0, ps.dir.Right()}
}

func distance(distances map[PathState]int, state PathState) int {
	if dist, ok := distances[state]; ok {
		return dist
	}
	return math.MaxInt
}

func filterInvalid(states []PathState, grid Grid) []PathState {
	var validStates []PathState
	for _, state := range states {
		if state.i >= 0 && state.i < len(grid) && state.j >= 0 && state.j < len(grid[0]) {
			validStates = append(validStates, state)
		}
	}
	return validStates
}

func shortestPath(grid Grid, source PathState) (map[PathState]int, map[PathState]PathState) {
	distances := map[PathState]int{source: 0}
	previous := map[PathState]PathState{}
	queue := collections.NewPriorityQueue[PathState](func(a, b PathState) bool {
		return distances[a] < distances[b]
	})
	queue.Add(source)

	for !queue.Empty() {
		state := queue.Poll()
		var nextStates []PathState
		if state.movesThisDir < 10 {
			nextStates = append(nextStates, state.Move())
		}
		if state.movesThisDir >= 4 {
			nextStates = append(nextStates, state.Left().Move())
			nextStates = append(nextStates, state.Right().Move())
		}
		nextStates = filterInvalid(nextStates, grid)
		for _, nextState := range nextStates {
			val := int(grid[nextState.i][nextState.j] - '0')
			alt := distance(distances, state) + val
			if alt < distance(distances, nextState) {
				distances[nextState] = alt
				previous[nextState] = state
				if queue.Contains(nextState) {
					queue.Update()
				} else {
					queue.Add(nextState)
				}
			}
		}
	}
	return distances, previous
}

func run(part2 bool, input string) any {
	if !part2 {
		return "not implemented"
	}

	lines := util.Lines(input)
	var grid Grid = make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}

	initialState := PathState{0, 0, 0, East}
	distances, _ := shortestPath(grid, initialState)

	minDistance := math.MaxInt
	for state, dist := range distances {
		if state.i == len(grid)-1 && state.j == len(grid[0])-1 && state.movesThisDir >= 4 {
			//fmt.Printf("%v -> %d\n", state, dist)
			minDistance = min(minDistance, dist)
		}
	}

	//fmt.Println(path)
	//fmt.Println()
	//for _, state := range path {
	//	var ch rune
	//	switch state.dir {
	//	case North:
	//		ch = '^'
	//	case East:
	//		ch = '>'
	//	case South:
	//		ch = 'v'
	//	case West:
	//		ch = '<'
	//	}
	//	grid[state.i][state.j] = ch
	//}
	//fmt.Println(grid)

	return minDistance
}
