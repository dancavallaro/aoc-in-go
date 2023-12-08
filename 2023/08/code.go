package main

import (
	"aoc-in-go/pkg/util"
)

func main() {
	//aoc.Harness(run)
	util.Run(run, "2023/08/input-user.txt", false)
}

type Node struct {
	label       string
	left, right string
}

func findNode(network map[string]Node, start string, instructions string) string {
	currNode := start

	for _, ins := range []rune(instructions) {
		node := network[currNode]
		if ins == 'L' {
			currNode = node.left
		} else if ins == 'R' {
			currNode = node.right
		}
	}

	return currNode
}

func run(part2 bool, input string) any {
	if part2 {
		return "not implemented"
	}

	lines := util.AllLines(input)
	instructions := lines[0]

	network := map[string]Node{}
	for _, line := range lines[2:] {
		if line == "" {
			continue
		}
		label := line[0:3]
		node := Node{
			label: label,
			left:  line[7:10],
			right: line[12:15],
		}
		network[label] = node
	}

	numSteps := 0
	currNode := "AAA"
	for {
		currNode = findNode(network, currNode, instructions)
		numSteps += len(instructions)

		if currNode == "ZZZ" {
			break
		}
	}

	return numSteps
}
