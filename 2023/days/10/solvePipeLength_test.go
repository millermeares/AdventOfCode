package ten

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolvePipeLength_Example(t *testing.T) {
	input := []string{
		"-L|F7",
		"7S-7|",
		"L|7||",
		"-L-J|",
		"L|-JF",
	}
	l := SolvePipeLength(input)
	assert.Equal(t, 4, l)
}

func TestSolveEncodedArea_Example1(t *testing.T) {
	input := []string{
		"...........",
		".S-------7.",
		".|F-----7|.",
		".||.....||.",
		".||.....||.",
		".|L-7.F-J|.",
		".|..|.|..|.",
		".L--J.L--J.",
		"...........",
	}
	enclosed := Part2(input)
	assert.Equal(t, 4, enclosed)
}

func TestSolveEncodedArea_BetweenPipes(t *testing.T) {
	input := []string{
		"..........",
		".S------7.",
		".|F----7|.",
		".||....||.",
		".||....||.",
		".|L-7F-J|.",
		".|..||..|.",
		".L--JL--J.",
		"..........",
	}
	enclosed := Part2(input)
	assert.Equal(t, 4, enclosed)
}
