package main

import (
	"aoc-in-go/pkg/util"
	"fmt"
	"strconv"
)

func main() {
	//aoc.Harness(run)
	util.Run(run, "2023/03/input-user.txt", true)
}

type symbol rune

func (sym symbol) String() string {
	return fmt.Sprintf("sym:%s", string(sym))
}

type number int

func (num number) String() string {
	return fmt.Sprintf("num:%s", strconv.Itoa(int(num)))
}

func parseLine(line string) []any {
	parsedLine := make([]any, len(line))
	var curNum = new(number)

	for i, c := range line {
		if c >= '0' && c <= '9' {
			*curNum = 10*(*curNum) + number(c-'0')
			parsedLine[i] = curNum
		} else {
			if *curNum > 0 {
				curNum = new(number)
			}
			if c != '.' {
				parsedLine[i] = symbol(c)
			}
		}
	}
	return parsedLine
}

func run(part2 bool, input string) any {
	if !part2 {
		return "not implemented"
	}

	lines := make([][]any, 0)
	for _, line := range util.Lines(input) {
		lines = append(lines, parseLine(line))
	}

	gearRatioSum := 0
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[0]); j++ {
			switch lines[i][j].(type) {
			case symbol:
				if lines[i][j] == symbol('*') {
					nearbyPartNums := findPartNumsNear(lines, i, j)
					if len(nearbyPartNums) == 2 {
						gearRatioSum += int(nearbyPartNums[0]) * int(nearbyPartNums[1])
					}
				}
			}
		}
	}

	return gearRatioSum
}

func findPartNumsNear(lines [][]any, i int, j int) []number {
	partNums := make([]number, 0)
	foundPartNums := make(map[*number]bool)
	adjacencies := [...][2]int{
		{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1},
	}

	for _, move := range adjacencies {
		newI := i + move[0]
		newJ := j + move[1]

		if newI >= 0 && newI < len(lines) && newJ >= 0 && newJ < len(lines[0]) {
			switch lines[newI][newJ].(type) {
			case *number:
				num := lines[newI][newJ].(*number)
				if !foundPartNums[num] {
					partNums = append(partNums, *num)
					foundPartNums[num] = true
				}
			}
		}
	}

	return partNums
}
