package main

import (
	"aoc-in-go/pkg/util"
	"math"
	"strconv"
	"strings"
)

func main() {
	//aoc.Harness(run)
	util.Run(run, "2023/05/input-user.txt", true)
}

type MapEntry struct {
	interval Interval
	delta    int
}

func NewMapEntry(destRangeStart int, sourceRangeStart int, rangeLength int) MapEntry {
	return MapEntry{
		interval: Interval{sourceRangeStart, sourceRangeStart + rangeLength - 1},
		delta:    destRangeStart - sourceRangeStart,
	}
}

func (entry MapEntry) Convert(source Interval) (Interval, []Interval, bool) {
	intersection, intersects := entry.interval.Intersection(source)
	if !intersects {
		return Interval{}, []Interval{source}, false
	}

	var leftoverIntervals []Interval
	if source.start < entry.interval.start {
		leftoverIntervals = append(leftoverIntervals, Interval{source.start, entry.interval.start - 1})
	}
	if source.end > entry.interval.end {
		leftoverIntervals = append(leftoverIntervals, Interval{entry.interval.end + 1, source.end})
	}

	return intersection.Add(entry.delta), leftoverIntervals, true
}

type Map struct {
	entries []MapEntry
}

func (m Map) Convert(sources []Interval) []Interval {
	var results []Interval
	remainingSources := sources
	for _, entry := range m.entries {
		var allLeftovers []Interval
		for _, interval := range remainingSources {
			converted, leftovers, didConvert := entry.Convert(interval)
			if didConvert {
				results = append(results, converted)
			}
			allLeftovers = append(allLeftovers, leftovers...)
		}
		remainingSources = allLeftovers
	}
	return append(results, remainingSources...)
}

type Interval struct {
	start, end int // Inclusive
}

func (i Interval) Intersection(o Interval) (Interval, bool) {
	latestStart := max(i.start, o.start)
	earliestEnd := min(i.end, o.end)

	if latestStart > earliestEnd {
		return Interval{}, false
	} else {
		return Interval{latestStart, earliestEnd}, true
	}
}

func (i Interval) Add(delta int) Interval {
	return Interval{i.start + delta, i.end + delta}
}

type Almanac struct {
	seeds []Interval
	maps  []Map
}

func (al Almanac) MinLocation() int {
	results := al.seeds
	for _, m := range al.maps {
		results = m.Convert(results)
	}

	minLocation := math.MaxInt
	for _, locationRange := range results {
		minLocation = min(minLocation, locationRange.start)
	}
	return minLocation
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

func parseSeeds(line string) []Interval {
	var seeds []Interval
	seedRanges := parseInts(line)
	numRanges := len(seedRanges) / 2

	for i := 0; i < numRanges; i++ {
		rangeStart := seedRanges[2*i]
		rangeSize := seedRanges[2*i+1]
		rangeEnd := rangeStart + rangeSize - 1
		seeds = append(seeds, Interval{rangeStart, rangeEnd})
	}

	return seeds
}

func parseAlmanac(input string) Almanac {
	almanac := Almanac{}
	var currMap *Map
	for _, line := range util.AllLines(input) {
		if strings.HasPrefix(line, "seeds:") {
			almanac.seeds = parseSeeds(line[7:])
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
	if !part2 {
		return "not implemented"
	}

	return parseAlmanac(input).MinLocation()
}
