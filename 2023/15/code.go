package main

import (
	"aoc-in-go/pkg/util"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//aoc.Harness(run)
	util.Run(run, "2023/15/input-user.txt", true)
}

type Lense struct {
	label       string
	focalLength int
}

type Box []Lense

func (b *Box) Insert(l Lense) {
	for i, lense := range *b {
		if lense.label == l.label {
			(*b)[i] = l
			return
		}
	}
	*b = append(*b, l)
}

func (b *Box) Remove(label string) {
	for i, lense := range *b {
		if lense.label == label {
			for j := i; j < len(*b)-1; j++ {
				(*b)[j] = (*b)[j+1]
			}
			if len(*b) == 1 {
				*b = nil
			} else {
				*b = (*b)[0 : len(*b)-1]
			}
			break
		}
	}
}

type StorageFacility [256]Box

func NewStorageFacility() StorageFacility {
	return [256]Box{}
}

func (sf *StorageFacility) Perform(step string) {
	fmt.Printf("Performing step '%s'\n", step)
	idx := strings.IndexAny(step, "-=")
	label := step[0:idx]
	box := Hash(label)

	if step[idx] == '-' {
		sf[box].Remove(label)
	} else if step[idx] == '=' {
		focalLength, err := strconv.Atoi(step[idx+1:])
		if err != nil {
			panic(err)
		}
		lense := Lense{label, focalLength}
		sf[box].Insert(lense)
	}
}

func (sf *StorageFacility) FocusingPower() int {
	power := 0
	for boxNum, box := range sf {
		for slotNum, lense := range box {
			power += (boxNum + 1) * (slotNum + 1) * lense.focalLength
		}
	}
	return power
}

func (sf StorageFacility) String() string {
	var sb strings.Builder
	for i, box := range sf {
		if len(box) == 0 {
			continue
		}
		sb.WriteString(fmt.Sprintf("Box %d: %v\n", i, box))
	}
	return sb.String()
}

func Hash(input string) int {
	hash := 0
	for _, ch := range []rune(input) {
		hash += int(ch)
		hash *= 17
		hash %= 256
	}
	return hash
}

func run(part2 bool, input string) any {
	if !part2 {
		return "not implemented"
	}

	input = util.Lines(input)[0]
	steps := strings.Split(input, ",")
	sf := NewStorageFacility()
	for _, step := range steps {
		sf.Perform(step)
		//fmt.Println(sf)
	}

	return sf.FocusingPower()
}
