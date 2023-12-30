package main

import (
	"aoc-in-go/pkg/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	//aoc.Harness(run)
	util.Run(run, "2023/19/input-user.txt", false)
}

var workflowRegex = regexp.MustCompile("([a-z]+){(.*)}")

type Operator int

const (
	Gt Operator = iota
	Lt
)

func (o Operator) String() string {
	if o == Gt {
		return ">"
	}
	return "<"
}

type Conditional struct {
	attribute string
	operator  Operator
	threshold int
}

func (c Conditional) String() string {
	return fmt.Sprintf("%s%v%d", c.attribute, c.operator, c.threshold)
}

type Rule struct {
	conditional *Conditional
	target      string
}

func (r Rule) String() string {
	if r.conditional == nil {
		return fmt.Sprintf("[%v]", r.target)
	}
	return fmt.Sprintf("[%v -> %s]", r.conditional, r.target)
}

type Workflow struct {
	label string
	rules []Rule
}

func (w Workflow) String() string {
	return fmt.Sprintf("%s: %v", w.label, w.rules)
}

func parseRule(rule string) Rule {
	var conditional *Conditional
	if idx := strings.Index(rule, ":"); idx != -1 {
		attribute, operator := rule[0:1], Lt
		if rule[1] == '>' {
			operator = Gt
		}
		threshold, err := strconv.Atoi(rule[2:idx])
		if err != nil {
			panic(err)
		}
		conditional = &Conditional{attribute, operator, threshold}
		rule = rule[idx+1:]
	}
	return Rule{conditional, rule}
}

func parseWorkflow(line string) Workflow {
	matches := workflowRegex.FindStringSubmatch(line)
	label, rulesStr := matches[1], matches[2]
	ruleParts := strings.Split(rulesStr, ",")
	var rules []Rule
	for _, ruleStr := range ruleParts {
		rules = append(rules, parseRule(ruleStr))
	}

	return Workflow{label, rules}
}

type Part map[string]int

func (p Part) Rating() int {
	return p["x"] + p["m"] + p["a"] + p["s"]
}

var partRegex = regexp.MustCompile("{x=([0-9]+),m=([0-9]+),a=([0-9]+),s=([0-9]+)}")

func parseInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return num
}

func parsePart(part string) Part {
	matches := partRegex.FindStringSubmatch(part)
	return Part{
		"x": parseInt(matches[1]),
		"m": parseInt(matches[2]),
		"a": parseInt(matches[3]),
		"s": parseInt(matches[4]),
	}
}

const (
	Accept = "A"
	Reject = "R"
)

func evaluatePart(workflows map[string]Workflow, part Part) int {
	target := "in"

	for {
		if target == Accept || target == Reject {
			break
		}
		workflow := workflows[target]
		for _, rule := range workflow.rules {
			if cond := rule.conditional; cond != nil {
				val := part[cond.attribute]
				if cond.operator == Gt && val <= cond.threshold {
					continue
				}
				if cond.operator == Lt && val >= cond.threshold {
					continue
				}
			}
			target = rule.target
			break
		}
	}

	if target == Accept {
		return part.Rating()
	}
	return 0
}

func run(part2 bool, input string) any {
	if part2 {
		return "not implemented"
	}

	workflows := map[string]Workflow{}
	lines := util.AllLines(input)
	var index int
	for index = 0; index < len(lines); index++ {
		line := lines[index]
		if line == "" {
			break
		}
		workflow := parseWorkflow(line)
		workflows[workflow.label] = workflow
	}
	index++

	totalRating := 0
	for ; index < len(lines)-1; index++ {
		part := parsePart(lines[index])
		totalRating += evaluatePart(workflows, part)
	}

	return totalRating
}
