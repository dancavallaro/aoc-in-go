package main

import (
	"aoc-in-go/pkg/util"
	"math"
	"strconv"
	"strings"
)

func main() {
	//aoc.Harness(run)
	util.Run(run, "2023/05/input-user.txt", false)
}

type MapEntry struct {
	sourceRangeStart int
	rangeLength      int
	destDelta        int
}

func NewMapEntry(destRangeStart int, sourceRangeStart int, rangeLength int) MapEntry {
	return MapEntry{
		sourceRangeStart: sourceRangeStart,
		rangeLength:      rangeLength,
		destDelta:        destRangeStart - sourceRangeStart,
	}
}

func (entry MapEntry) Convert(source int) (int, bool) {
	if source >= entry.sourceRangeStart && source < entry.sourceRangeStart+entry.rangeLength {
		return source + entry.destDelta, true
	} else {
		return source, false
	}
}

type Map struct {
	entries []MapEntry
}

func (m Map) Convert(source int) int {
	for _, entry := range m.entries {
		dest, converted := entry.Convert(source)
		if converted {
			return dest
		}
	}
	return source
}

type Almanac struct {
	seeds []int
	maps  []Map
}

func (al Almanac) MapLocation(seed int) int {
	result := seed
	for _, nextMap := range al.maps {
		result = nextMap.Convert(result)
	}
	return result
}

func parseInts(line string) []int {
	var ints []int
	for _, intStr := range strings.Split(line, " ") {
		theInt, err := strconv.Atoi(intStr)
		if err != nil {
			panic(err)
		}
		ints = append(ints, theInt)
	}
	return ints
}

func parseAlmanac(input string) Almanac {
	almanac := Almanac{}
	var currMap *Map
	for _, line := range util.AllLines(input) {
		if strings.HasPrefix(line, "seeds:") {
			almanac.seeds = parseInts(line[7:])
		}

		if strings.Contains(line, " map:") {
			currMap = &Map{}
		} else if currMap != nil {
			if line == "" {
				almanac.maps = append(almanac.maps, *currMap)
				currMap = nil
			} else {
				entry := parseInts(line)
				destRangeStart, sourceRangeStart, rangeLength := entry[0], entry[1], entry[2]
				currMap.entries = append(currMap.entries, NewMapEntry(destRangeStart, sourceRangeStart, rangeLength))
			}
		}

	}
	return almanac
}

func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return "not implemented"
	}

	almanac := parseAlmanac(input)
	lowestLocation := math.MaxInt
	for _, seed := range almanac.seeds {
		location := almanac.MapLocation(seed)
		lowestLocation = min(lowestLocation, location)
	}

	return lowestLocation
}
