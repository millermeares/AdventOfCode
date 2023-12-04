package three

import (
	"unicode"
)

type EngineSchematic struct {
	schematic [][]rune
}

func parseEngineSchematic(input []string) EngineSchematic {
	e := EngineSchematic{
		schematic: make([][]rune, len(input)),
	}
	for i, line := range input {
		e.addSchematicLine(i, line)
	}
	return e
}

func (e EngineSchematic) addSchematicLine(lineNum int, input string) {
	var line []rune
	for _, val := range input {
		line = append(line, val)
	}
	e.schematic[lineNum] = line
}

func (e EngineSchematic) getEndOfNumberIndex(y, xStart int) int {
	row := e.schematic[y]
	xEnd := xStart
	for ; xEnd < len(row) && isDigit(row[xEnd]); xEnd++ {

	}
	return xEnd - 1
}

func (e EngineSchematic) isAdjacentToSpecialCharacter(y, xStart, xEnd int) bool {
	return len(e.adjacentSpecialCharacters(y, xStart, xEnd)) != 0
}

func (e EngineSchematic) specialCharacterLeft(y, start int) []Point {
	if start == 0 || !isSpecialCharacter(e.schematic[y][start-1]) {
		return []Point{}
	}
	return []Point{
		{
			x: start - 1,
			y: y,
		},
	}
}

func (e EngineSchematic) specialCharacterRight(y, end int) []Point {
	if end == len(e.schematic[y])-1 || !isSpecialCharacter(e.schematic[y][end+1]) {
		return []Point{}
	}
	return []Point{
		{
			x: end + 1,
			y: y,
		},
	}
}

func (e EngineSchematic) specialCharacterAbove(y, xStart, xEnd int) []Point {
	if y == 0 {
		// top row
		return []Point{}
	}
	return e.specialCharacterInSection(y-1, xStart-1, xEnd+2)
}

func (e EngineSchematic) specialCharacterBelow(y, xStart, xEnd int) []Point {
	if y == len(e.schematic)-1 {
		// bottom row
		return []Point{}
	}
	return e.specialCharacterInSection(y+1, xStart-1, xEnd+2)
}

func (e EngineSchematic) specialCharacterInSection(y, xStart, xEnd int) []Point {
	row := e.schematic[y]
	var points []Point
	for i := xStart; i < xEnd && i < len(row); i++ {
		if i < 0 {
			continue
		}
		if isSpecialCharacter(row[i]) {
			points = append(points, Point{
				y: y,
				x: i,
			})
		}
	}
	return points
}

func isSpecialCharacter(char rune) bool {
	// if not . and not digit, true
	return !unicode.IsDigit(char) && string(char) != "."
}

func isDigit(char rune) bool {
	return unicode.IsDigit(char)
}
