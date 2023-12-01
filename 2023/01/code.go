package main

import (
	"github.com/jpillora/puzzler/harness/aoc"
	"strings"
)

func main() {
	aoc.Harness(run)
	//b, err := os.ReadFile("/Users/dan/workspace/aoc-in-go/2023/01/input-user.txt")
	//if err != nil {
	//	panic(err)
	//}
	//str := string(b)
	//fmt.Println(runPart2(str))
}

func firstDigit(s string) int {
	for _, c := range s {
		if c >= '0' && c <= '9' {
			return int(c - '0')
		}
	}
	return -1
}

func reverse(s string) string {
	chars := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

func value(line string) int {
	first := firstDigit(line)
	last := firstDigit(reverse(line))
	return 10*first + last
}

func valuePart2(line string) int {
	first := firstNum(line, digitPrefixes, digitStrings)
	last := firstNum(reverse(line), revDigitPrefixes, revDigitStrings)
	//fmt.Printf("%s: first = %d, last = %d\n", line, first, last)
	return 10*first + last
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return runPart2(input)
	} else {
		return runPart1(input)
	}
}

func firstNum(s string, prefixes Trie, digits map[string]int) int {
	for i, c := range s {
		if c >= '0' && c <= '9' {
			return int(c - '0')
		} else {
			maybeDigitString := ""

			for j := i; j < len(s); j++ {
				if s[j] >= '0' && s[j] <= '9' {
					break
				}
				maybeDigitString += string(s[j])
				prefix, word := prefixes.Contains(maybeDigitString)

				if word {
					return digits[maybeDigitString]
				} else if !prefix {
					break
				}
			}
		}
	}
	panic("no valid nums found")
}

var digitStrings = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}
var digitPrefixes = NewTrie()
var revDigitStrings = make(map[string]int)
var revDigitPrefixes = NewTrie()

func runPart2(input string) any {
	lines := strings.Split(input, "\n")
	calibrationValue := 0

	for str, digit := range digitStrings {
		revStr := reverse(str)
		revDigitStrings[revStr] = digit
		digitPrefixes.Insert(str)
		revDigitPrefixes.Insert(revStr)
	}

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		calibrationValue += valuePart2(line)
	}

	return calibrationValue
}

func runPart1(input string) any {
	lines := strings.Split(input, "\n")
	calibrationValue := 0

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		calibrationValue += value(line)
	}

	return calibrationValue
}

type Trie struct {
	root *node
}

type node struct {
	value    rune
	isWord   bool
	children [26]*node
}

func newNode(v rune) *node {
	return &node{value: v, isWord: false}
}

func NewTrie() Trie {
	return Trie{root: newNode(0)}
}

func (trie Trie) Insert(word string) {
	curNode := trie.root
	for _, c := range []rune(word) {
		charNum := int(c - 'a')
		if curNode.children[charNum] == nil {
			curNode.children[charNum] = newNode(c)
		}
		curNode = curNode.children[charNum]
	}
	curNode.isWord = true
}

func (trie Trie) Contains(s string) (prefix bool, word bool) {
	curNode := trie.root
	for _, c := range []rune(s) {
		charNum := int(c - 'a')
		if curNode.children[charNum] == nil {
			prefix, word = false, false
			return
		}
		curNode = curNode.children[charNum]
	}
	prefix, word = true, curNode.isWord
	return
}

func (trie Trie) ContainsPrefix(s string) bool {
	prefix, _ := trie.Contains(s)
	return prefix
}

func (trie Trie) ContainsWord(s string) bool {
	_, word := trie.Contains(s)
	return word
}
