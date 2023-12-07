package three

import (
	"days"
	"strconv"
)

func GetDay() days.Day {
	return days.MakeDay(SolveEngineSchematic, SolveGearRatio, "03")
}

func SolveEngineSchematic(input []string) int {
	e := parseEngineSchematic(input)
	return e.calculateSchematicSum()
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
