package seven

import "days"

func GetDay() days.Day {
	return days.MakeDay(SolveCamelCardWinnings, SolveCamelCardWinningsWithJoker, "07")
}

func SolveCamelCardWinnings(input []string) int {
	hands := parseHands(input)
	SortHands(hands)
	return calculateTotalWinnings(hands)
}

func SolveCamelCardWinningsWithJoker(input []string) int {
	hands := parseHands(input)
	for _, h := range hands {
		makeJacksJokers(&h)
	}
	SortHands(hands)
	return calculateTotalWinnings(hands)
}

func makeJacksJokers(h *Hand) {
	for i := 0; i < len(h.cards); i++ {
		if h.cards[i] == 11 {
			h.cards[i] = jokerKey
		}
	}
}

func calculateTotalWinnings(hands []Hand) int {
	sum := 0
	for i, hand := range hands {
		rank := len(hands) - i
		sum += rank * hand.bid
	}
	return sum
}
