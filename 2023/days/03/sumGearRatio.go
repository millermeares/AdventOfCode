package three

import "strconv"

func SolveGearRatio(input []string) int {
	e := parseEngineSchematic(input)
	return e.calculateGearRatio()
}

type Point struct {
	x int
	y int
}

func (e EngineSchematic) calculateGearRatio() int {
	specialPoints := e.getSpecialPointsMap()
	return calculateGearRatioFromPoints(specialPoints)
}

func (e EngineSchematic) getSpecialPointsMap() map[Point][]string {
	specialPoints := make(map[Point][]string)
	for y, row := range e.schematic {
		for x := 0; x < len(row); x++ {
			val := row[x]
			if !isDigit(val) {
				continue
			}
			endIdx := e.getEndOfNumberIndex(y, x)
			num := string(row[x : endIdx+1])

			adjacentSpecialPoints := e.adjacentSpecialCharacters(y, x, endIdx)
			addToSpecialPointsMap(specialPoints, adjacentSpecialPoints, num)
			x = endIdx // dont count this number again.
		}
	}
	return specialPoints
}

func addToSpecialPointsMap(specialPoints map[Point][]string, points []Point, num string) {
	for _, point := range points {
		_, exists := specialPoints[point]
		if !exists {
			specialPoints[point] = []string{}
		}
		specialPoints[point] = append(specialPoints[point], num)
	}
}

func calculateGearRatioFromPoints(specialPoints map[Point][]string) int {
	sum := 0
	for _, v := range specialPoints {
		if len(v) != 2 {
			continue
		}
		numOne, _ := strconv.Atoi(v[0])
		numTwo, _ := strconv.Atoi(v[1])
		sum += (numOne * numTwo)
	}
	return sum
}

func (e EngineSchematic) adjacentSpecialCharacters(y, xStart, xEnd int) []Point {
	points := e.specialCharacterAbove(y, xStart, xEnd)
	points = append(points, e.specialCharacterBelow(y, xStart, xEnd)...)
	points = append(points, e.specialCharacterRight(y, xEnd)...)
	points = append(points, e.specialCharacterLeft(y, xStart)...)
	return points
}
