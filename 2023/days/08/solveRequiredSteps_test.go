package eight

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveRequiredSteps_Example(t *testing.T) {
	input := []string{
		"RL",
		"",
		"AAA = (BBB, CCC)",
		"BBB = (DDD, EEE)",
		"CCC = (ZZZ, GGG)",
		"DDD = (DDD, DDD)",
		"EEE = (EEE, EEE)",
		"GGG = (GGG, GGG)",
		"ZZZ = (ZZZ, ZZZ",
	}
	steps := SolveRequiredSteps(input)
	assert.Equal(t, 2, steps)
}
func TestSolveRequiredSteps_LoopingExample(t *testing.T) {
	input := []string{
		"LLR",
		"",
		"AAA = (BBB, BBB)",
		"BBB = (AAA, ZZZ)",
		"ZZZ = (ZZZ, ZZZ",
	}
	steps := SolveRequiredSteps(input)
	assert.Equal(t, 6, steps)
}

func TestSolveRequiredSteps_Ghost(t *testing.T) {
	input := []string{
		"LR",
		"",
		"11A = (11B, XXX)",
		"11B = (XXX, 11Z)",
		"11Z = (11B, XXX)",
		"22A = (22B, XXX)",
		"22B = (22C, 22C)",
		"22C = (22Z, 22Z)",
		"22Z = (22B, 22B)",
		"XXX = (XXX, XXX)",
	}
	steps := Part2(input)
	assert.Equal(t, 6, steps)
}
