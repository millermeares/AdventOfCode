package seven

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveCamelCardWinnings(t *testing.T) {
	input := []string{
		"32T3K 765",
		"T55J5 684",
		"KK677 28",
		"KTJJT 220",
		"QQQJA 483",
	}
	winnings := SolveCamelCardWinnings(input)
	assert.Equal(t, 6440, winnings)
}

func TestSolveCamelCardWinnings_WithJoker(t *testing.T) {
	input := []string{
		"32T3K 765",
		"T55J5 684",
		"KK677 28",
		"KTJJT 220",
		"QQQJA 483",
	}
	winnings := SolveCamelCardWinningsWithJoker(input)
	assert.Equal(t, 5905, winnings)
}
