package sixteen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1_Example(t *testing.T) {
	input := []string{
		".|...\\....",
		"|.-.\\.....",
		".....|-...",
		"........|.",
		"..........",
		".........\\",
		"..../.\\\\..",
		".-.-/..|..",
		".|....-|.\\",
		"..//.|....",
	}
	assert.Equal(t, 46, Part1(input))
}

func TestPart2_Example(t *testing.T) {
	input := []string{
		".|...\\....",
		"|.-.\\.....",
		".....|-...",
		"........|.",
		"..........",
		".........\\",
		"..../.\\\\..",
		".-.-/..|..",
		".|....-|.\\",
		"..//.|....",
	}
	assert.Equal(t, 51, Part2(input))
}
