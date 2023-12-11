package eleven

import (
	"days"
	"math"
	"strings"
)

func GetDay() days.Day {
	return days.MakeDay(SolveGalaxyDistance, Part2, "11")
}

type Point struct {
	x int
	y int
}

func (p Point) getDistanceBetween(other Point) int {
	xDiff := int(math.Abs(float64(p.x - other.x)))
	yDiff := int(math.Abs(float64(p.y - other.y)))
	return xDiff + yDiff
}

func SolveGalaxyDistance(input []string) int {
	// any rows or columns that contain no galaxies should be twice as big.
	rowsExpanded := expandAbsentGalaxyRows(input)
	columnsExpanded := expandAbsentGalaxyColumns(rowsExpanded)
	galaxies := getGalaxies(columnsExpanded)
	sum := 0
	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			galaxyOne := galaxies[i]
			galaxyTwo := galaxies[j]
			sum += galaxyOne.getDistanceBetween(galaxyTwo)
		}
	}
	return sum
}

func getGalaxies(input []string) []Point {
	var points []Point
	for y, line := range input {
		for x, c := range line {
			if c == '#' {
				points = append(points, Point{x: x, y: y})
			}
		}
	}
	return points
}

func expandAbsentGalaxyColumns(input []string) []string {
	for i := 0; i < len(input[0]); i++ {
		// for each character
		if columnHasGalaxy(i, input) {
			continue // no modifications needed
		}
		// need to add a column here.
		input = addPeriodColumn(i, input)
		i++ // prevent double counting none-galaxy columns.
	}
	return input
}

func addPeriodColumn(col int, input []string) []string {
	for i := 0; i < len(input); i++ {
		input[i] = input[i][:col] + "." + input[i][col:]
	}
	return input
}

func columnHasGalaxy(col int, input []string) bool {
	for _, line := range input {
		if line[col] == '#' {
			return true
		}
	}
	return false
}

func expandAbsentGalaxyRows(input []string) []string {
	var expanded []string
	for _, line := range input {
		if !rowHasGalaxy(line) {
			expanded = append(expanded, getRowOfPeriods(len(line)))
		}
		expanded = append(expanded, line)
	}
	return expanded
}

func getRowOfPeriods(l int) string {
	str := ""
	for i := 0; i < l; i++ {
		str += "."
	}
	return str
}

func rowHasGalaxy(line string) bool {
	return strings.Contains(line, "#")
}

func Part2(input []string) int {
	return solveDynamicGalaxyWeight(input, 1000000)
}

func solveDynamicGalaxyWeight(input []string, emptyWeight int) int {
	rowCosts := getRowCosts(input, emptyWeight)
	columnCosts := getColumnCosts(input, emptyWeight)
	galaxies := getGalaxies(input)
	sum := 0
	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			galaxyOne := galaxies[i]
			galaxyTwo := galaxies[j]
			sum += galaxyOne.getWeightedCost(galaxyTwo, rowCosts, columnCosts)
		}
	}
	return sum
}

func (p Point) getWeightedCost(other Point, rowCosts []int, colCosts []int) int {
	return getColTravelCosts(p, other, colCosts) + getRowTravelCosts(p, other, rowCosts)
}

func getColTravelCosts(p1 Point, p2 Point, colCosts []int) int {
	smallerY := p1.y
	largerY := p2.y
	if p2.y < p1.y {
		largerY = p1.y
		smallerY = p2.y
	}
	sum := 0
	for i := smallerY; i < largerY; i++ {
		sum += colCosts[i]
	}
	return sum
}

func getRowTravelCosts(p1 Point, p2 Point, rowCosts []int) int {
	smallerX := p1.x
	largerX := p2.x
	if p2.x < p1.x {
		largerX = p1.x
		smallerX = p2.x
	}
	sum := 0
	for i := smallerX; i < largerX; i++ {
		sum += rowCosts[i]
	}
	return sum
}

func getRowCosts(input []string, emptyGalaxyWeight int) []int {
	costs := make([]int, len(input))
	for i, line := range input {
		costs[i] = 1
		if !rowHasGalaxy(line) {
			costs[i] = emptyGalaxyWeight
		}
	}
	return costs
}

func getColumnCosts(input []string, emptyGalaxyWeight int) []int {
	costs := make([]int, len(input[0]))
	for i := 0; i < len(input[0]); i++ {
		costs[i] = 1
		if !columnHasGalaxy(i, input) {
			costs[i] = emptyGalaxyWeight
		}
	}
	return costs
}
