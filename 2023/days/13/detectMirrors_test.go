package thirteen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetectMirrors_Example(t *testing.T) {
	input := []string{
		"#.##..##.",
		"..#.##.#.",
		"##......#",
		"##......#",
		"..#.##.#.",
		"..##..##.",
		"#.#.##.#.",
		"",
		"#...##..#",
		"#....#..#",
		"..##..###",
		"#####.##.",
		"#####.##.",
		"..##..###",
		"#....#..#",
	}
	assert.Equal(t, 405, DetectMirrors(input))
}

func TestDetectMirrors_Single(t *testing.T) {
	input := []string{
		//...LR
		"#.##..##.",
		"..#.##.#.",
		"##......#",
		"##......#",
		"..#.##.#.",
		"..##..##.",
		"#.#.##.#.",
	}
	idx := getIndexOfVerticalReflectionLine(input, -1)
	assert.Equal(t, 4, idx)

	input = []string{
		"#...##..#",
		"#....#..#",
		"..##..###",
		"#####.##.", //T
		"#####.##.", //B
		"..##..###",
		"#....#..#",
	}
	inverted := invert(input)
	idx = getIndexOfVerticalReflectionLine(inverted, -1)
	assert.Equal(t, 3, idx)
}

func TestDetectSmudgedMirrors_Example(t *testing.T) {
	input := []string{
		"#.##..##.",
		"..#.##.#.",
		"##......#",
		"##......#",
		"..#.##.#.",
		"..##..##.",
		"#.#.##.#.",
		"",
		"#...##..#",
		"#....#..#",
		"..##..###",
		"#####.##.",
		"#####.##.",
		"..##..###",
		"#....#..#",
	}
	assert.Equal(t, 400, DetectMirrorsWithSmudge(input))
}

func TestDetectSmudgeMirrors_Individual(t *testing.T) {
	input := []string{
		"#.##..##.",
		"..#.##.#.",
		"##......#",
		"##......#",
		"..#.##.#.",
		"..##..##.",
		"#.#.##.#.",
	}
	assert.Equal(t, 300, DetectMirrorsWithSmudge(input))
	input = []string{
		"#...##..#",
		"#....#..#",
		"..##..###",
		"#####.##.",
		"#####.##.",
		"..##..###",
		"#....#..#",
	}
	assert.Equal(t, 100, DetectMirrorsWithSmudge(input))
}
