package twentyone

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	input := []string{
		"...........",
		".....###.#.",
		".###.##..#.",
		"..#.#...#..",
		"....#.#....",
		".##..S####.",
		".##..#...#.",
		".......##..",
		".##.#.####.",
		".##..##.##.",
		"...........",
	}
	assert.Equal(t, 16, countViableSpots(input, 6, false))
}

func TestPart2(t *testing.T) {
	input := []string{
		"...........",
		".....###.#.",
		".###.##..#.",
		"..#.#...#..",
		"....#.#....",
		".##..S####.",
		".##..#...#.",
		".......##..",
		".##.#.####.",
		".##..##.##.",
		"...........",
	}
	assert.Equal(t, 16, countViableSpots(input, 6, true))
	assert.Equal(t, 50, countViableSpots(input, 10, true))
	// assert.Equal(t, 1594, countViableSpots(input, 50, true))
	// assert.Equal(t, 6536, countViableSpots(input, 100, true))
}
