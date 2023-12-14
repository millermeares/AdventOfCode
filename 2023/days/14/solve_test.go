package fourteen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1_Example(t *testing.T) {
	input := []string{
		"O....#....",
		"O.OO#....#",
		".....##...",
		"OO.#O....O",
		".O.....O#.",
		"O.#..O.#.#",
		"..O..#O..O",
		".......O..",
		"#....###..",
		"#OO..#....",
	}
	assert.Equal(t, 136, Part1(input))
}

func TestPart2_Example(t *testing.T) {
	input := []string{
		"O....#....",
		"O.OO#....#",
		".....##...",
		"OO.#O....O",
		".O.....O#.",
		"O.#..O.#.#",
		"..O..#O..O",
		".......O..",
		"#....###..",
		"#OO..#....",
	}
	oneSpin := spinCycle(input)
	expected := []string{
		".....#....",
		"....#...O#",
		"...OO##...",
		".OO#......",
		".....OOO#.",
		".O#...O#.#",
		"....O#....",
		"......OOOO",
		"#...O###..",
		"#..OO#....",
	}

	assert.Equal(t, expected, oneSpin)
	twoSpins := spinCycle(oneSpin)
	expected = []string{
		".....#....",
		"....#...O#",
		".....##...",
		"..O#......",
		".....OOO#.",
		".O#...O#.#",
		"....O#...O",
		".......OOO",
		"#..OO###..",
		"#.OOO#...O",
	}
	assert.Equal(t, expected, twoSpins)

	threeSpins := spinCycle(twoSpins)
	expected = []string{
		".....#....",
		"....#...O#",
		".....##...",
		"..O#......",
		".....OOO#.",
		".O#...O#.#",
		"....O#...O",
		".......OOO",
		"#...O###.O",
		"#.OOO#...O",
	}
	assert.Equal(t, expected, threeSpins)
	assert.False(t, true)
}
