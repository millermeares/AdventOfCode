package four

import (
	"math"
	"strings"
)

type ScratchCard struct {
	id             string
	winningNumbers []string
	cardNumbers    []string
}

func (sc ScratchCard) calculatePointValue() int {
	totalMatches := sc.numberOfMatches()
	if totalMatches == 0 {
		return 0
	}
	return int(math.Pow(float64(2), float64(totalMatches-1)))
}

func (sc ScratchCard) numberOfMatches() int {
	totalMatches := 0
	for _, cardNumber := range sc.cardNumbers {
		if contains(sc.winningNumbers, cardNumber) {
			totalMatches++
		}
	}
	return totalMatches
}

func parseScratchCard(line string) ScratchCard {
	idSep := strings.Split(line, ":")
	id := strings.Split(idSep[0], " ")[1]

	numberSep := strings.Split(idSep[1], "|")
	winningNumbers := strings.Fields(numberSep[0])
	cardNumbers := strings.Fields(numberSep[1])
	return ScratchCard{
		id:             id,
		winningNumbers: winningNumbers,
		cardNumbers:    cardNumbers,
	}
}

func contains(list []string, val string) bool {
	for _, val2 := range list {
		if val2 == val {
			return true
		}
	}
	return false
}
