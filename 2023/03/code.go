package main

import (
	"aoc-in-go/pkg/util"
	"fmt"
	"strconv"
)

func main() {
	//aoc.Harness(run)
	util.Run(run, "2023/03/input-user.txt", false)
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
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return "not implemented"
	}

	lines := make([][]any, 0)
	for _, line := range util.Lines(input) {
		lines = append(lines, parseLine(line))
	}

	partNumSum := 0
	foundPartNums := make(map[*number]bool)
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[0]); j++ {
			switch lines[i][j].(type) {
			case symbol:
				sumNearbyPartNums := sumPartNumsNear(lines, foundPartNums, i, j)
				partNumSum += sumNearbyPartNums
			}
		}
	}

	return partNumSum
}

func sumPartNumsNear(lines [][]any, foundPartNums map[*number]bool, i int, j int) int {
	partNumSum := 0
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
					foundPartNums[num] = true
					partNumSum += int(*num)
				}
			}
		}
	}

	return partNumSum
}
