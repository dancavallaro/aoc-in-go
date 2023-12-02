package util

import (
	"fmt"
	"os"
	"strings"
)

func Lines(input string) []string {
	nonEmptyLines := make([]string, 0)
	for _, line := range strings.Split(input, "\n") {
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
