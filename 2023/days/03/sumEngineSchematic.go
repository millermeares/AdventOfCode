package three

import (
	"strconv"
	"unicode"
)

func SolveEngineSchematic(input []string) int {
	e := parseEngineSchematic(input)
	return e.calculateSchematicSum()
}

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

func (e EngineSchematic) calculateSchematicSum() int {
	sum := 0
	for y, row := range e.schematic {
		for x := 0; x < len(row); x++ {
			val := row[x]
			if !isDigit(val) {
				continue
			}
			endIdx := e.getEndOfNumberIndex(y, x)
			num := string(row[x : endIdx+1])
			if !e.isAdjacentToSpecialCharacter(y, x, endIdx) {
				x = endIdx // dont count this number again.
				continue
			}
			// add to sum.
			parsedInt, _ := strconv.Atoi(num)
			sum = sum + parsedInt
			x = endIdx // dont count this number again.
		}
	}
	return sum
}

func (e EngineSchematic) getEndOfNumberIndex(y, xStart int) int {
	row := e.schematic[y]
	xEnd := xStart
	for ; xEnd < len(row) && isDigit(row[xEnd]); xEnd++ {

	}
	return xEnd - 1
}

func (e EngineSchematic) isAdjacentToSpecialCharacter(y, xStart, xEnd int) bool {
	return e.specialCharacterAbove(y, xStart, xEnd) || e.specialCharacterBelow(y, xStart, xEnd) ||
		e.specialCharacterRight(y, xEnd) || e.specialCharacterLeft(y, xStart)
}

func (e EngineSchematic) specialCharacterLeft(y, start int) bool {
	if start == 0 {
		return false
	}
	return isSpecialCharacter(e.schematic[y][start-1])
}

func (e EngineSchematic) specialCharacterRight(y, end int) bool {
	if end == len(e.schematic[y])-1 {
		return false
	}
	return isSpecialCharacter(e.schematic[y][end+1])
}

func (e EngineSchematic) specialCharacterAbove(y, xStart, xEnd int) bool {
	if y == 0 {
		// top row
		return false
	}
	return specialCharacterInSection(e.schematic[y-1], xStart-1, xEnd+2)
}

func (e EngineSchematic) specialCharacterBelow(y, xStart, xEnd int) bool {
	if y == len(e.schematic)-1 {
		// bottom row
		return false
	}
	return specialCharacterInSection(e.schematic[y+1], xStart-1, xEnd+2)
}

func specialCharacterInSection(row []rune, xStart, xEnd int) bool {
	for i := xStart; i < xEnd && i < len(row); i++ {
		if i < 0 {
			continue
		}
		if isSpecialCharacter(row[i]) {
			return true
		}
	}
	return false
}

func isSpecialCharacter(char rune) bool {
	// if not . and not digit, true
	return !unicode.IsDigit(char) && string(char) != "."
}

func isDigit(char rune) bool {
	return unicode.IsDigit(char)
}
