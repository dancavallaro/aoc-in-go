package util

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func AllLines(input string) []string {
	return strings.Split(input, "\n")
}

func Lines(input string) []string {
	nonEmptyLines := make([]string, 0)
	for _, line := range AllLines(input) {
		if len(line) > 0 {
			nonEmptyLines = append(nonEmptyLines, line)
		}
	}
	return nonEmptyLines
}

func Run(runFunc func(bool, string) any, inputPath string, part2 bool) {
	bytes, err := os.ReadFile(inputPath)
	if err != nil {
		panic(err)
	}

	input := string(bytes)
	fmt.Println(runFunc(part2, input))
}

func ParseInts(line string) []int {
	var ints []int
	for _, intStr := range strings.Split(line, " ") {
		if intStr == "" {
			continue
		}
		theInt, err := strconv.Atoi(intStr)
		if err != nil {
			panic(err)
		}
		ints = append(ints, theInt)
	}
	return ints
}

func InvertMap[K comparable, V comparable](m map[K]V) map[V]K {
	inverted := map[V]K{}
	for k, v := range m {
		inverted[v] = k
	}
	return inverted
}
