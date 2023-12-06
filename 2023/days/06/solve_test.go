package six

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveWaysToBeatRecord_Example(t *testing.T) {
	input := []string{
		"Time:      7  15   30",
		"Distance:  9  40  200",
	}
	waysToBeat := SolveWaysToBeatRecord(input)
	assert.Equal(t, 288, waysToBeat)
}

func TestWaysToBeatRecord_SingleRace(t *testing.T) {
	input := []string{
		"Time:      7  15   30",
		"Distance:  9  40  200",
	}
	waysToBeat := SolveWaysToBeatRecordSingleRace(input)
	assert.Equal(t, 71503, waysToBeat)
}
