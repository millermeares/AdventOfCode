package eleven

import (
	"days"
	"strings"
)

func GetDay() days.Day {
	return days.MakeDay(SolveGalaxyDistance, Part2, "11")
}

type Point struct {
	id int
	x  int
	y  int
}

func SolveGalaxyDistance(input []string) int {
	return solveDynamicGalaxyWeight(input, 2)
}

func getGalaxies(input []string) []Point {
	var points []Point
	for y, line := range input {
		for x, c := range line {
			if c == '#' {
				points = append(points, Point{x: x, y: y, id: len(points) + 1})
				// p := points[len(points)-1]
				// fmt.Println("Id", p.id, "is at y", p.y, "x", p.x)
			}
		}
	}
	return points
}

func columnHasGalaxy(col int, input []string) bool {
	for _, line := range input {
		if line[col] == '#' {
			return true
		}
	}
	return false
}

func rowHasGalaxy(line string) bool {
	return strings.Contains(line, "#")
}

func Part2(input []string) int {
	return solveDynamicGalaxyWeight(input, 1000000)
}

func solveDynamicGalaxyWeight(input []string, emptyWeight int) int {
	rowCosts := getRowCosts(input, emptyWeight)
	// fmt.Println("r costs", rowCosts)
	columnCosts := getColumnCosts(input, emptyWeight)
	// fmt.Println("c costs", columnCosts)
	galaxies := getGalaxies(input)
	sum := 0
	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			galaxyOne := galaxies[i]
			galaxyTwo := galaxies[j]
			weightedCost := galaxyOne.getWeightedCost(galaxyTwo, rowCosts, columnCosts)
			// fmt.Println("Cost from", galaxyOne.id, "to", galaxyTwo.id, "is", weightedCost)
			sum += weightedCost
		}
	}
	return sum
}

func (p Point) getWeightedCost(other Point, rowCosts []int, colCosts []int) int {
	return getColTravelCosts(p, other, colCosts) + getRowTravelCosts(p, other, rowCosts)
}

func getColTravelCosts(p1 Point, p2 Point, colCosts []int) int {
	smallerX := p1.x
	largerX := p2.x
	if smallerX > largerX {
		temp := smallerX
		smallerX = largerX
		largerX = temp
	}
	sum := 0
	for i := smallerX + 1; i <= largerX; i++ {
		// fmt.Println("Adding cost", colCosts[i], "for traversal on column", i)
		sum += colCosts[i]
	}
	// fmt.Println("Column travel", sum)
	return sum
}

func getRowTravelCosts(p1 Point, p2 Point, rowCosts []int) int {
	smallerY := p1.y
	largerY := p2.y
	if smallerY > largerY {
		temp := smallerY
		smallerY = largerY
		largerY = temp
	}
	sum := 0
	for i := smallerY + 1; i <= largerY; i++ {
		// fmt.Println("Adding cost", rowCosts[i], "for traversal on row", i)
		sum += rowCosts[i]
	}
	// fmt.Println("Row travel", sum)
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

// func printLines(input []string) {
// 	for _, line := range input {
// 		fmt.Println(line)
// 	}
// }
