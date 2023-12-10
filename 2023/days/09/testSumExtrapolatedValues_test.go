package nine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtrapoledValuesExample(t *testing.T) {
	input := []string{
		"0 3 6 9 12 15",
	}
	extrapolated := SolveExtrapolatedValues(input)
	assert.Equal(t, 18, extrapolated)
}

func TestExtrapolatedValuesExample2(t *testing.T) {
	input := []string{
		"0 3 6 9 12 15",
		"1 3 6 10 15 21",
		"10 13 16 21 30 45",
	}
	extrapolated := SolveExtrapolatedValues(input)
	assert.Equal(t, 114, extrapolated)
}
