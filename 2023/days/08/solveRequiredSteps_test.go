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
