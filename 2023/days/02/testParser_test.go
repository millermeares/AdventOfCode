package two

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGameParser_MaxValue(t *testing.T) {
	input := "Game 66: 1 green, 6 blue; 7 blue, 8 green; 2 blue, 9 red, 14 green"
	game := buildGame(input)
	assert.Equal(t, 66, game.id)
	assert.Equal(t, 3, len(game.handfuls))
	assert.Equal(t, 14, game.maxValueOfColor(Green))
	assert.Equal(t, 7, game.maxValueOfColor(Blue))
	assert.Equal(t, 9, game.maxValueOfColor(Red))
}

func TestHandfulParser(t *testing.T) {
	input := "1 green, 6 blue"
	handful := buildHandful(input)

	// make sure keys are valid
	for k := range handful.colorCount {
		assert.True(t, k == Blue || k == Green || k == Red, "unexpected key: "+k)
	}

	assert.Equal(t, 2, len(handful.colorCount))

	assert.Equal(t, 1, handful.getValueOfColor(Green))
	assert.Equal(t, 6, handful.getValueOfColor(Blue))
}

func TestAddColorToMap(t *testing.T) {
	input := "6 green"
	colorCount := make(map[string]int)
	addColorToMap(colorCount, input)
	assert.Equal(t, 1, len(colorCount))
	assert.Equal(t, 6, colorCount[Green])
}
