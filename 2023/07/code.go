package main

import (
	"aoc-in-go/pkg/util"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	//aoc.Harness(run)
	util.Run(run, "input-user.txt", false)
}

type Rank int

const (
	HighCard Rank = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func maxValue(m map[rune]int) int {
	maxVal := 0
	for _, v := range m {
		maxVal = max(maxVal, v)
	}
	return maxVal
}

func rankHand(hand Hand) Rank {
	counts := map[rune]int{}
	for _, c := range hand.cards {
		counts[c]++
	}
	maxCount := maxValue(counts)

	if len(counts) == 1 {
		return FiveOfAKind
	} else if len(counts) == 2 && maxCount == 4 {
		return FourOfAKind
	} else if len(counts) == 2 && maxCount == 3 {
		return FullHouse
	} else if maxCount == 3 {
		return ThreeOfAKind
	} else if len(counts) == 3 {
		return TwoPair
	} else if maxCount == 2 {
		return OnePair
	}

	return HighCard
}

func rankCards(ordering []rune) map[rune]int {
	ranking := map[rune]int{}
	for i, r := range ordering {
		ranking[r] = len(ordering) - 1 - i // I'm lazy and don't want to reverse the CardOrder list
	}
	return ranking
}

var CardOrder = []rune{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'}
var CardRankings = rankCards(CardOrder)

// TODO: are hands of type HighCard compared based on their highest card, or
// in order of cards like the regular tiebreaker?
func strongerThan(hand1 Hand, hand2 Hand) bool {
	for i := 0; i < 5; i++ {
		rank1, rank2 := CardRankings[hand1.cards[i]], CardRankings[hand2.cards[i]]
		if rank1 > rank2 {
			return true
		} else if rank1 < rank2 {
			return false
		}
		// If these cards are the same, keep going and compare the next card
	}

	panic("hands are an exact tie!")
}

type Cards []rune

func (c Cards) String() string {
	return string(c)
}

type Hand struct {
	cards Cards
	bid   int
}

func (h Hand) String() string {
	return fmt.Sprintf("{cards=%s, bid=%d}", h.cards, h.bid)
}

func parseHand(line string) Hand {
	parts := strings.Split(line, " ")
	cards := Cards(parts[0])
	bid, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}
	hand := Hand{cards, bid}
	//fmt.Printf("hand = %v, bid = %d; rank = %d\n", cards, bid, rankHand(cards))

	return hand
}

func run(part2 bool, input string) any {
	if part2 {
		return "not implemented"
	}

	var hands []Hand
	for _, line := range util.Lines(input) {
		hands = append(hands, parseHand(line))
	}

	slices.SortFunc(hands, func(h1 Hand, h2 Hand) int {
		rank1, rank2 := rankHand(h1), rankHand(h2)
		if rank1 > rank2 {
			return 1
		} else if rank1 < rank2 {
			return -1
		} else {
			if strongerThan(h1, h2) {
				return 1
			}
			return -1
		}
	})

	totalWinnings := 0
	for i, hand := range hands {
		totalWinnings += hand.bid * (i + 1)
	}
	return totalWinnings
}
