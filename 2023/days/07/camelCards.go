package seven

import (
	"sort"
	"strconv"
	"strings"
	"unicode"
)

const (
	jokerKey = 1
)

type Hand struct {
	cards []int // could also do based on int. convert to int from A,K,Q,J in parsing?
	bid   int
}

func SortHands(hands []Hand) {
	// sort ascending so that bad hands are first.
	sort.SliceStable(hands, func(i, j int) bool {
		h1s := hands[i].getStrength()
		h2s := hands[j].getStrength()
		if h1s == h2s {
			return tiebreak(hands[i], hands[j])
		}
		return h1s > h2s
	})
}

// number value reflects
func (h *Hand) getStrength() int {
	// maybe count by card?
	countByCard := map[int]int{}
	for _, c := range h.cards {
		v, exists := countByCard[c]
		if !exists {
			v = 0
		}
		countByCard[c] = v + 1
	}

	_, jackIsJoker := countByCard[jokerKey]
	// only care about if jack is joker *if* there are cards of "joker" value.
	if jackIsJoker {
		addJokerCountToHighestCard(countByCard)
	}

	if len(countByCard) == 1 || len(countByCard) == 0 { // only 1 other card or 1 non-joker.
		// five of a kind.
		return 6
	}
	if len(countByCard) == 2 {
		max := maxVal(countByCard)
		if max == 4 {
			return 5 // 4 of a kind
		}
		return 4 // full house
	}
	if len(countByCard) == 3 {
		max := maxVal(countByCard)
		if max == 3 {
			return 3 // 3 of a kind
		}
		return 2 // two pair
	}

	if len(countByCard) == 4 {
		return 1 // one pair
	}
	if len(countByCard) == 5 {
		return 0 // high card
	}
	panic("Unexpected hand strength")
}

func maxVal(counts map[int]int) int {
	max := 0
	for _, v := range counts {
		if v > max {
			max = v
		}
	}
	return max
}

func tiebreak(h1, h2 Hand) bool {
	for i := 0; i < 5; i++ {
		c1 := h1.cards[i]
		c2 := h2.cards[i]
		if c1 == c2 {
			continue
		}
		return c1 > c2 // todo: maybe <
	}
	panic("Same hand")
}

func parseHands(input []string) []Hand {
	var hands []Hand
	for _, line := range input {
		hands = append(hands, parseHand(line))
	}
	return hands
}

func parseHand(line string) Hand {
	fields := strings.Fields(line)
	bid, _ := strconv.Atoi(fields[1])
	return Hand{
		bid:   bid,
		cards: parseCards(fields[0]),
	}
}

func parseCards(input string) []int {
	var cards []int
	for _, l := range input {
		cards = append(cards, cardToInt(l))
	}
	return cards
}

func cardToInt(r rune) int {
	if unicode.IsDigit(r) {
		return int(r - '0') // 2-9
	}
	if r == 'A' {
		return 14
	}
	if r == 'K' {
		return 13
	}
	if r == 'Q' {
		return 12
	}
	if r == 'J' {
		return 11
	}
	if r == 'T' {
		return 10
	}
	panic("Unexpected input" + string(r))
}

func addJokerCountToHighestCard(countByCard map[int]int) {
	jokerCount := countByCard[jokerKey]
	delete(countByCard, jokerKey) // don't directly consider it.

	if jokerCount != 0 && len(countByCard) == 0 {
		countByCard[jokerKey] = jokerCount
	}
	maxKey := 0
	maxValue := 0
	for k, v := range countByCard {
		if v > maxValue {
			maxKey = k
			maxValue = v
		}
	}
	countByCard[maxKey] += jokerCount
}
