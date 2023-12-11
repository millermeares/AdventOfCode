package eleven

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveGalaxyDistance_Example(t *testing.T) {
	input := []string{
		"...#......",
		".......#..",
		"#.........",
		"..........",
		"......#...",
		".#........",
		".........#",
		"..........",
		".......#..",
		"#...#.....",
	}
	sum := SolveGalaxyDistance(input)
	assert.Equal(t, 374, sum)
}

func TestSolveDynamicGalaxyWeight_100(t *testing.T) {
	input := []string{
		"...#......",
		".......#..",
		"#.........",
		"..........",
		"......#...",
		".#........",
		".........#",
		"..........",
		".......#..",
		"#...#.....",
	}
	sum := solveDynamicGalaxyWeight(input, 100)
	assert.Equal(t, 8410, sum)
}

func TestSolveDynamicGalaxyWeight_10(t *testing.T) {
	input := []string{
		"...#......",
		".......#..",
		"#.........",
		"..........",
		"......#...",
		".#........",
		".........#",
		"..........",
		".......#..",
		"#...#.....",
	}
	sum := solveDynamicGalaxyWeight(input, 10)
	assert.Equal(t, 1030, sum)
}
