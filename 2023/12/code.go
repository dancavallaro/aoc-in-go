package main

import (
	"aoc-in-go/pkg/util"
	"strconv"
	"strings"
)

func main() {
	//aoc.Harness(run)
	util.Run(run, "2023/12/input-user.txt", true)
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

func (gc groupCounts) sum() int {
	total := 0
	for _, count := range gc {
		total += count
	}
	return total
}

type cacheKey struct {
	springsUsed, damagedSpringsUsed int
	lastSpringState                 condition
}

type cache struct {
	originalConditions  []condition
	damagedSpringGroups groupCounts
	cacheStore          map[cacheKey]*int
}

func newCache(originalConditions []condition, damagedSpringGroups groupCounts) cache {
	return cache{originalConditions, damagedSpringGroups, map[cacheKey]*int{}}
}

func (c cache) makeKey(conditions []condition, damagedSprings groupCounts, lastSpringState condition) cacheKey {
	springsUsed := len(c.originalConditions) - len(conditions)
	damagedSpringsUsed := c.damagedSpringGroups.sum() - damagedSprings.sum()
	return cacheKey{springsUsed, damagedSpringsUsed, lastSpringState}
}

func countArrangementsRecurse(c cache, conditions []condition, damagedSprings groupCounts, lastSpringState condition, indent, s string) int {
	//fmt.Printf("%sconditions = %v, damagedSprings = %v, lastCondition = %v\n", indent, conditions, damagedSprings, lastSpringState)
	indent += "  "

	if len(conditions) == 0 {
		if damagedSprings.isEmpty() {
			//fmt.Printf(" ==> Found valid arrangement: %s\n", s)
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

	key := c.makeKey(conditions, damagedSprings, lastSpringState)
	if c.cacheStore[key] == nil {
		arrangements := 0
		if considerOperational {
			arrangements += countArrangementsRecurse(c, conditions[1:], damagedSprings, operational, indent, s+".")
		}
		if considerDamaged {
			damagedSprings = damagedSprings.decrement()
			arrangements += countArrangementsRecurse(c, conditions[1:], damagedSprings, damaged, indent, s+"#")
		}
		c.cacheStore[key] = &arrangements
	}
	return *c.cacheStore[key]
}

func countArrangements(conditions []condition, damagedGroupCounts []int) int {
	return countArrangementsRecurse(newCache(conditions, damagedGroupCounts), conditions, damagedGroupCounts, unknown, "", "")
}

const unfoldingMultiplier = 5

func repeat[T any](input []T, sep *T) []T {
	newLength := unfoldingMultiplier * len(input)
	if sep != nil {
		newLength += unfoldingMultiplier - 1
	}

	output := make([]T, newLength)
	currIdx := 0
	for i := 0; i < unfoldingMultiplier; i++ {
		copy(output[currIdx:], input)
		currIdx += len(input)

		if sep != nil && i < unfoldingMultiplier-1 {
			output[currIdx] = *sep
			currIdx++
		}
	}
	return output
}

func run(part2 bool, input string) any {
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

		if part2 {
			separator := unknown
			conditions, damagedGroupCounts = repeat(conditions, &separator), repeat(damagedGroupCounts, nil)
		}
		arrangements := countArrangements(conditions, damagedGroupCounts)
		//fmt.Printf("%d arrangements for line: %s\n", arrangements, line)
		totalArrangements += arrangements
	}

	return totalArrangements
}
