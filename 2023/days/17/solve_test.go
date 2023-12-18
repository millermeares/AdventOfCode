package seventeen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1_Example(t *testing.T) {
	input := []string{
		"2413432311323",
		"3215453535623",
		"3255245654254",
		"3446585845452",
		"4546657867536",
		"1438598798454",
		"4457876987766",
		"3637877979653",
		"4654967986887",
		"4564679986453",
		"1224686865563",
		"2546548887735",
		"4322674655533",
	}
	assert.Equal(t, 102, Part1(input))
}
