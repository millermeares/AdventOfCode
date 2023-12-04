package four

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScratchCard_Example1(t *testing.T) {
	line := "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"
	sc := parseScratchCard(line)
	value := sc.calculatePointValue()
	assert.Equal(t, 8, value)
}
