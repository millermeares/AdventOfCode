package seven

func SolveCamelCardWinnings(input []string) int {
	hands := parseHands(input)
	SortHands(hands)
	return calculateTotalWinnings(hands)
}

// i'm pretty sure that i have implemented as i intended, need to read the thing again.
func calculateTotalWinnings(hands []Hand) int {
	sum := 0
	for i, hand := range hands {
		rank := len(hands) - i
		sum += rank * hand.bid
	}
	return sum
}
