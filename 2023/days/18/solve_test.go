package eighteen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1_Example(t *testing.T) {
	input := []string{
		"R 6 (#70c710)",
		"D 5 (#0dc571)",
		"L 2 (#5713f0)",
		"D 2 (#d2c081)",
		"R 2 (#59c680)",
		"D 2 (#411b91)",
		"L 5 (#8ceee2)",
		"U 2 (#caa173)",
		"L 1 (#1b58a2)",
		"U 2 (#caa171)",
		"R 2 (#7807d2)",
		"U 3 (#a77fa3)",
		"L 2 (#015232)",
		"U 2 (#7a21e3)",
	}
	assert.Equal(t, 62, Part1(input))
}

func TestPart1_Inverted_Example(t *testing.T) {
	input := []string{
		"L 6 (#70c710)",
		"D 5 (#0dc571)",
		"R 2 (#5713f0)",
		"D 2 (#d2c081)",
		"L 2 (#59c680)",
		"D 2 (#411b91)",
		"R 5 (#8ceee2)",
		"U 2 (#caa173)",
		"R 1 (#1b58a2)",
		"U 2 (#caa171)",
		"L 2 (#7807d2)",
		"U 3 (#a77fa3)",
		"R 2 (#015232)",
		"U 2 (#7a21e3)",
	}
	assert.Equal(t, 62, Part1(input))
}

func TestPart1_UpsideDown_Example(t *testing.T) {
	input := []string{
		"R 6 (#70c710)",
		"U 5 (#0dc571)",
		"L 2 (#5713f0)",
		"U 2 (#d2c081)",
		"R 2 (#59c680)",
		"U 2 (#411b91)",
		"L 5 (#8ceee2)",
		"D 2 (#caa173)",
		"L 1 (#1b58a2)",
		"D 2 (#caa171)",
		"R 2 (#7807d2)",
		"D 3 (#a77fa3)",
		"L 2 (#015232)",
		"D 2 (#7a21e3)",
	}
	assert.Equal(t, 62, Part1(input))
}

func TestPart1_Playground(t *testing.T) {
	l1 := Line{
		start: Point{x: 0, y: 0}, end: Point{x: 6, y: 0},
	}
	l2 := Line{
		start: Point{x: 0, y: 2}, end: Point{x: 2, y: 2},
	}
	assert.Equal(t, true, linesOverlap(l1, l2))
}

func TestCountEnclosed_Rectangle(t *testing.T) {
	l1 := Line{
		start: Point{x: 0, y: 0}, end: Point{x: 6, y: 0},
	}
	l2 := Line{
		start: Point{x: 0, y: 2}, end: Point{x: 6, y: 2},
	}
	assert.Equal(t, 21, countEnclosed([]Line{l1, l2}))
}

func TestCountEnclosed_2(t *testing.T) {
	assert.Equal(t, 36, countEnclosed([]Line{
		{
			start: Point{x: 6, y: 0}, end: Point{x: 0, y: 0},
		},
		{
			start: Point{x: 0, y: 0}, end: Point{x: 0, y: 2},
		},
		{
			start: Point{x: 0, y: 2}, end: Point{x: 2, y: 2},
		},
		{
			start: Point{x: 2, y: 2}, end: Point{x: 2, y: 5},
		},
		{
			start: Point{x: 2, y: 5}, end: Point{x: 6, y: 5},
		},
		{
			start: Point{x: 6, y: 5}, end: Point{x: 6, y: 0},
		},
	}))
}

func TestCountEnclosed_2UpsideDown(t *testing.T) {
	assert.Equal(t, 36, countEnclosed([]Line{
		{
			start: Point{x: 6, y: 0}, end: Point{x: 0, y: 0},
		},
		{
			start: Point{x: 0, y: 0}, end: Point{x: 0, y: -2},
		},
		{
			start: Point{x: 0, y: -2}, end: Point{x: 2, y: -2},
		},
		{
			start: Point{x: 2, y: -2}, end: Point{x: 2, y: -5},
		},
		{
			start: Point{x: 2, y: -5}, end: Point{x: 6, y: -5},
		},
		{
			start: Point{x: 6, y: -5}, end: Point{x: 6, y: 0},
		},
	}))
}

func TestCountEnclosed_Inverted2(t *testing.T) {
	assert.Equal(t, 36, countEnclosed([]Line{
		{
			start: Point{x: -6, y: 0}, end: Point{x: 0, y: 0},
		},
		{
			start: Point{x: 0, y: 0}, end: Point{x: 0, y: 2},
		},
		{
			start: Point{x: 0, y: 2}, end: Point{x: -2, y: 2},
		},
		{
			start: Point{x: -2, y: 2}, end: Point{x: -2, y: 5},
		},
		{
			start: Point{x: -2, y: 5}, end: Point{x: -6, y: 5},
		},
		{
			start: Point{x: -6, y: 5}, end: Point{x: -6, y: 0},
		},
	}))
}

func TestPart2_Example(t *testing.T) {
	input := []string{
		"R 6 (#70c710)",
		"D 5 (#0dc571)",
		"L 2 (#5713f0)",
		"D 2 (#d2c081)",
		"R 2 (#59c680)",
		"D 2 (#411b91)",
		"L 5 (#8ceee2)",
		"U 2 (#caa173)",
		"L 1 (#1b58a2)",
		"U 2 (#caa171)",
		"R 2 (#7807d2)",
		"U 3 (#a77fa3)",
		"L 2 (#015232)",
		"U 2 (#7a21e3)",
	}
	assert.Equal(t, 952408144115, Part2(input))
}

func TestParseHex(t *testing.T) {
	hex := "#70c71"
	parsed := parseHexLength(hex)
	assert.Equal(t, parsed, 461937)
}
