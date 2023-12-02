package two

import (
	"strings"
)

type Handful struct {
	colorCount map[string]int
}

func buildHandfuls(input string) []Handful {
	eachHandful := strings.Split(input, ";")
	var handfuls []Handful
	for _, handfulInput := range eachHandful {
		handfuls = append(handfuls, buildHandful(handfulInput))
	}
	return handfuls
}

func buildHandful(input string) Handful {
	colorCounts := make(map[string]int)
	colors := strings.Split(input, ",")
	for _, color := range colors {
		processedColor := strings.TrimLeft(color, " ")
		addColorToMap(colorCounts, processedColor)
	}
	return Handful{
		colorCount: colorCounts,
	}
}

func (h Handful) getValueOfColor(color string) int {
	val, present := h.colorCount[color]
	if !present {
		return 0
	}
	return val
}
