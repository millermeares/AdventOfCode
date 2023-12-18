package eighteen

import (
	"fmt"
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

func TestPart1_Playground(t *testing.T) {
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
	p := Point{y: 1, x: 3}
	lines := toLines(parseInput(input))
	fmt.Println(lines)
	_, maxX := xRange(lines)
	fmt.Println("max x", maxX)
	enclosed := pointEnclosed(p, maxX, lines)
	assert.Equal(t, true, enclosed)
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
