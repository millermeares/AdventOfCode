package three

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumGearRatio_Example(t *testing.T) {
	input := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}
	sum := SolveGearRatio(input)
	assert.Equal(t, 467835, sum)
}
