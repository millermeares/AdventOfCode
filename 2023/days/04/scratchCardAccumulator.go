package four

func SolveScratchCardAccumulator(input []string) int {
	cardCounts := make([]int, len(input))
	for i := range cardCounts {
		cardCounts[i] = 1 // each card has a value of 1.
	}
	for i, line := range input {
		sc := parseScratchCard(line)
		incrementCounts(cardCounts, i, sc)
	}
	// sum of cardCounts
	sum := 0
	for _, val := range cardCounts {
		sum += val
	}
	return sum
}

func incrementCounts(counts []int, cardIndex int, sc ScratchCard) {
	matches := sc.numberOfMatches()
	cardCount := counts[cardIndex]
	for i := cardIndex + 1; i < len(counts) && i-cardIndex <= matches; i++ {
		counts[i] += cardCount
	}
}
