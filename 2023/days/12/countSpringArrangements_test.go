package twelve

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountSpringArrangements_Example1(t *testing.T) {
	input := "???.### 1,1,3"
	count := countSpringArrangements(input)
	assert.Equal(t, 1, count)
}

// func TestCalculateState(t *testing.T) {
// 	broken := []int{1, 1, 3}
// 	input := "???.###"
// 	state := calculateState(input, broken)
// 	assert.Equal(t, State{unmatchedBrokenIndex: 0, unmatchedInputIndex: 0}, state)
// }

func TestCountSpringArrangements_Example2(t *testing.T) {
	input := "?###???????? 3,2,1"
	assert.Equal(t, 10, countSpringArrangements(input))

	input = ".??..??...?##. 1,1,3"
	assert.Equal(t, 4, countSpringArrangements(input))

	input = "?#?#?#?#?#?#?#? 1,3,1,6"
	assert.Equal(t, 1, countSpringArrangements(input))

	input = "????.######..#####. 1,6,5"
	assert.Equal(t, 4, countSpringArrangements(input))
}

func TestCountSpringArrangements_Playground(t *testing.T) {
	input := "??.? 1,1"
	assert.Equal(t, 2, countSpringArrangements(input))

	input = "#???.? 1,1"
	assert.Equal(t, 3, countSpringArrangements(input))
}

func TestCountSpringArrangements_SumExample(t *testing.T) {
	input := []string{
		"???.### 1,1,3",
		".??..??...?##. 1,1,3",
		"?#?#?#?#?#?#?#? 1,3,1,6",
		"????.#...#... 4,1,1",
		"????.######..#####. 1,6,5",
		"?###???????? 3,2,1",
	}
	sum := SumCountSpringArrangements(input)
	assert.Equal(t, 21, sum)
}

func TestCountSpringArrangements_ExpandedExample1(t *testing.T) {
	input := "???.### 1,1,3"
	count := countExpandedSpringArrangement(input)
	assert.Equal(t, 1, count)

	input = ".??..??...?##. 1,1,3"
	count = countExpandedSpringArrangement(input)
	assert.Equal(t, 16384, count)
}
