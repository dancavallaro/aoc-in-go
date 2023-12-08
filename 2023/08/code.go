package main

import (
	"aoc-in-go/pkg/util"
	"fmt"
)

func main() {
	//aoc.Harness(run)
	util.Run(run, "2023/08/input-user.txt", true)
}

type Node struct {
	label       string
	left, right string
}

func findNodes(network map[string]Node, starts []string, instructions string) ([]string, int, bool) {
	currNodes := starts
	stepsTaken := 0
	foundEnd := false

	for _, ins := range []rune(instructions) {
		stepsTaken++
		for i, currNode := range currNodes {
			node := network[currNode]
			if ins == 'L' {
				currNode = node.left
			} else if ins == 'R' {
				currNode = node.right
			}
			currNodes[i] = currNode
		}
		if allEndNodes(currNodes) {
			foundEnd = true
			break
		}
	}

	return currNodes, stepsTaken, foundEnd
}

func allEndNodes(nodes []string) bool {
	for _, label := range nodes {
		if label[2] != 'Z' {
			return false
		}
	}
	return true
}

func run(part2 bool, input string) any {
	if !part2 {
		return "not implemented"
	}

	lines := util.AllLines(input)
	instructions := lines[0]

	network := map[string]Node{}
	var startNodes []string
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

		lastChar := label[2]
		if lastChar == 'A' {
			startNodes = append(startNodes, label)
		}
	}

	for _, startNode := range startNodes {
		numSteps := 0
		currNodes := []string{startNode}
		for {
			nextNodes, stepsTaken, foundEnd := findNodes(network, currNodes, instructions)
			currNodes = nextNodes
			numSteps += stepsTaken
			if foundEnd {
				break
			}
		}
		fmt.Printf("%s to %s took %d steps\n", startNode, currNodes[0], numSteps)
	}

	return -1
}
