package main

import (
	"aoc-in-go/pkg/util"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//aoc.Harness(run)
	util.Run(run, "2023/12/input-user.txt", false)
}

type condition int

const (
	operational condition = iota
	damaged
	unknown
)

type groupCounts []int

func (gc groupCounts) decrement() groupCounts {
	newGroupCounts := make([]int, len(gc))
	copy(newGroupCounts, gc)
	newGroupCounts[0]--
	return newGroupCounts
}

func (gc groupCounts) isEmpty() bool {
	return len(gc) == 0 || (len(gc) == 1 && gc[0] == 0)
}

func (gc groupCounts) justFinishedGroup() bool {
	return gc[0] == 0
}

func (gc groupCounts) dropEmpty() groupCounts {
	if len(gc) > 0 && gc[0] == 0 {
		return gc[1:]
	}
	return gc
}

/*
State is represented by springs remaining (conditions slice), damaged group counts remaining,
and condition of the previous spring in this arrangement.

We start with the full original conditions slice and damaged group counts, and "previous spring state"
unknown (this is the only time we'll use this state, since when we choose an arrangement all springs
will be either damaged or operational, never unknown).

Base case is when the conditions slice is empty, i.e. we've used all the springs. If damaged
group counts is empty (or only has a 0 entry), return 1, otherwise return 0. Damaged group
counts being empty is not a base case, it just means there can't be any remaining damaged
springs (could be either 1 or 0, depending).

State transitions depend on the current spring status:

- operational:
  - if the last spring was damaged and we did NOT just end a group, return 0
  - otherwise, recurse and go to the next spring
- damaged:
  - if damaged group counts is 0 or empty, OR we just completed a group, return 0
  - decrement damaged group count, and recurse to next spring
- unknown:
  - if the last spring wasn't damaged, or it was but we ended a group, then we can consider this operational
  - if damaged group counts is non-empty/zero, and we did NOT just complete a group, we can consider this damaged
  - return the sum of the two optional parts
*/

func countArrangementsRecurse(conditions []condition, damagedSprings groupCounts, lastSpringState condition, indent, s string) int {
	//fmt.Printf("%sconditions = %v, damagedSprings = %v, lastCondition = %v\n", indent, conditions, damagedSprings, lastSpringState)
	indent += "  "

	if len(conditions) == 0 {
		if damagedSprings.isEmpty() {
			fmt.Printf(" ==> Found valid arrangement: %s\n", s)
			return 1
		} else {
			return 0
		}
	}

	considerDamaged, considerOperational := false, false
	if conditions[0] == operational || conditions[0] == unknown {
		considerOperational = lastSpringState != damaged || damagedSprings.justFinishedGroup()
	}
	if conditions[0] == damaged || conditions[0] == unknown {
		considerDamaged = !damagedSprings.isEmpty() && !damagedSprings.justFinishedGroup()
	}

	damagedSprings = damagedSprings.dropEmpty()
	arrangements := 0
	if considerOperational {
		arrangements += countArrangementsRecurse(conditions[1:], damagedSprings, operational, indent, s+".")
	}
	if considerDamaged {
		damagedSprings = damagedSprings.decrement()
		arrangements += countArrangementsRecurse(conditions[1:], damagedSprings, damaged, indent, s+"#")
	}
	return arrangements
}

func countArrangements(conditions []condition, damagedGroupCounts []int) int {
	return countArrangementsRecurse(conditions, damagedGroupCounts, unknown, "", "")
}

func run(part2 bool, input string) any {
	if part2 {
		return "not implemented"
	}

	totalArrangements := 0
	for _, line := range util.Lines(input) {
		parts := strings.Split(line, " ")

		var conditions []condition
		for _, ch := range []rune(parts[0]) {
			var cond condition
			switch ch {
			case '.':
				cond = operational
			case '#':
				cond = damaged
			case '?':
				cond = unknown
			}
			conditions = append(conditions, cond)
		}

		var damagedGroupCounts []int
		for _, countStr := range strings.Split(parts[1], ",") {
			count, err := strconv.Atoi(countStr)
			if err != nil {
				panic(err)
			}
			damagedGroupCounts = append(damagedGroupCounts, count)
		}

		arrangements := countArrangements(conditions, damagedGroupCounts)
		fmt.Printf("%d arrangements for line: %s\n", arrangements, line)
		totalArrangements += arrangements
	}

	return totalArrangements
}
