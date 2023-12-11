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

func TestSolveEncodedArea_ManualExpansion(t *testing.T) {
	// input := []string{
	// 	".o.o.o.o.o.o.o.o.o.o.o",
	// 	"oooooooooooooooooooooo",
	// 	".oS---------------7o.o",
	// 	"oo|ooooooooooooooo|ooo",
	// 	".o|oF-----------7o|o.o",
	// 	"oo|o|ooooooooooo|o|ooo",
	// 	".o|o|o.o.o.o.o.o|o|o.o",
	// 	"oo|o|ooooooooooo|o|ooo",
	// 	".o|o|o.o.o.o.o.o|o|o.o",
	// 	"oo|o|ooooooooooo|o|ooo",
	// 	".o|oL---7o.oF---Jo|o.o",
	// 	"oo|ooooo|ooo|ooooo|ooo",
	// 	".o|o.o.o|o.o|o.o.o|o.o",
	// 	"oo|ooooo|ooo|ooooo|ooo",
	// 	".oL-----Jo.oL-----Jo.o",
	// 	"oooooooooooooooooooooo",
	// 	".o.o.o.o.o.o.o.o.o.o.o",
	// 	"oooooooooooooooooooooo",
	// }
	// this broke. it worked until i added 5 dots across in the trapped area. interesting.
	input := []string{
		"....................",
		"..S-------------7...",
		"..|.............|...",
		"..|...F-------7.|...",
		"..|...|.......|.|...",
		"..|...|.......|.|...",
		"..|...L-7...F-J.|...",
		"..|.....|...|...|...",
		"..|.....|...|...|...",
		"..|.....|...|...|...",
		"..L-----J...L---J...",
		"....................",
	}
	enclosed := countTrappedSquares(input)
	assert.Equal(t, 37, enclosed)
}

func TestPg(t *testing.T) {
	input := []string{
		".........",
		".S-----7.",
		".|.....|.",
		".|.....|.",
		".|.....|.",
		".L-----J.",
		".........",
	}
	enclosed := countTrappedSquares(input)
	assert.Equal(t, 15, enclosed)
}
