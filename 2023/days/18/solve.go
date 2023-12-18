package eighteen

import (
	"days"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func GetDay() days.Day {
	return days.MakeDay(Part1, Part2, "18")
}

func Part1(input []string) int {
	digs := parseInput(input)
	points := digsToPoints(digs)
	grid := pointsToGrid(points)
	fillInEnclosedPoints(grid)
	return sum(grid)
}

func Part2(input []string) int {
	digs := parseHexadecimalInput(input)
	lines := toLines(digs)
}

func countEnclosed(lines []Line) int {

}

func parseHexadecimalInput(input []string) []Dig {
	var digs []Dig
	for _, line := range input {
		line = strings.Fields(line)[2]
		direction := directionFromHexRune(rune(line[7]))
		hexLength := line[1:7]
		digs = append(digs, Dig{
			direction: direction,
			meters:    parseHexLength(hexLength),
		})
	}
	return digs
}

func parseHexLength(hexLength string) int {
	fmt.Println("parsing", hexLength)
	hexLength = strings.Replace(hexLength, "#", "", -1)
	parsed, e := strconv.ParseInt(hexLength, 16, 64)
	if e != nil {
		panic(e)
	}
	fmt.Println("Parsed length", parsed, "from", hexLength)
	return int(parsed)
}

func directionFromHexRune(r rune) string {
	if r == '0' {
		return "R"
	}
	if r == '1' {
		return "D"
	}
	if r == '2' {
		return "L"
	}
	if r == '3' {
		return "U"
	}
	panic("Unexpected hex direction: " + string(r))
}

func sum(grid [][]int) int {
	sum := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			sum += grid[y][x]
		}
	}
	return sum
}

func fillInEnclosedPoints(grid [][]int) {
	for y := len(grid) - 1; y >= 0; y-- {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] != 0 {
				continue
			}
			if isEnclosed(y, x, grid) {
				grid[y][x] = 1 // i bet part 2 will be manipulating this
			}
		}
	}
}

func isEnclosed(y, x int, grid [][]int) bool {
	if y == 0 || y == len(grid)-1 {
		return false // top and bottom edge of grid cannot be enclosed.
	}
	// count amount of times that it crosses a line. if even, it is not enclosed.
	// if odd, it is enclosed.
	count := 0
	for i := x; i < len(grid[y]); i++ {
		// if this point is a "wall" going upwards at least 1 square, count it.
		if grid[y-1][i] != 0 && grid[y][i] != 0 {
			count++
		}
	}
	return count%2 == 1
}

// points not on grid should have value 1. on grid should have value 0
func pointsToGrid(points []Point) [][]int {
	points = offsetPoints(minX(points), minY(points), points)
	maxX := maxX(points)
	maxY := maxY(points)
	grid := make([][]int, maxY+1)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]int, maxX+1)
	}

	for _, p := range points {
		grid[p.y][p.x] = 1
	}
	return grid
}

func offsetPoints(xOffset, yOffset int, points []Point) []Point {
	var offsetPoints []Point
	for _, p := range points {
		offsetPoints = append(offsetPoints, Point{x: p.x - xOffset, y: p.y - yOffset})
	}
	return offsetPoints
}

func minX(points []Point) int {
	min := math.MaxInt
	for _, p := range points {
		if min > p.x {
			min = p.x
		}
	}
	return min
}

func minY(points []Point) int {
	min := math.MaxInt
	for _, p := range points {
		if min > p.y {
			min = p.y
		}
	}
	return min
}

func maxX(points []Point) int {
	max := 0
	for _, p := range points {
		if p.x > max {
			max = p.x
		}
	}
	return max
}

func maxY(points []Point) int {
	max := 0
	for _, p := range points {
		if p.y > max {
			max = p.y
		}
	}
	return max
}

func digsToPoints(digs []Dig) []Point {
	start := Point{x: 0, y: 0}
	perimeter := []Point{}
	for _, dig := range digs {
		perimeter = append(perimeter, dig.toPoints(start)...)
		start = perimeter[len(perimeter)-1]
	}
	return perimeter
}

func (d Dig) toPoints(cur Point) []Point {
	xChange, yChange := getXYChange(d.direction)
	var points []Point
	for i := 0; i < d.meters; i++ {
		nextPoint := Point{x: cur.x + (xChange * (i + 1)), y: cur.y + (yChange * (i + 1))}
		points = append(points, nextPoint)
	}
	return points
}

func toLines(digs []Dig) []Line {
	var lines []Line
	start := Point{x: 0, y: 0}
	for _, dig := range digs {
		lines = append(lines, dig.toLine(start))
		start = lines[len(lines)-1].end
	}
	return lines
}

func (d Dig) toLine(cur Point) Line {
	xChange, yChange := getXYChange(d.direction)
	return Line{
		start: cur,
		end:   Point{x: cur.x + xChange, y: cur.y + yChange},
	}
}

func getXYChange(direction string) (int, int) {
	if direction == "U" {
		return 0, -1
	}
	if direction == "D" {
		return 0, 1
	}
	if direction == "R" {
		return 1, 0
	}
	if direction == "L" {
		return -1, 0
	}
	panic("Unexpected input " + direction)
}

type Point struct {
	x int
	y int
}

type Dig struct {
	meters    int
	direction string
}

type Line struct {
	start Point
	end   Point
}

func (l Line) isVertical() bool {
	return l.start.x == l.end.y
}

func parseInput(input []string) []Dig {
	var digs []Dig
	for _, line := range input {
		split := strings.Fields(line)
		meters, _ := strconv.Atoi(split[1])
		digs = append(digs, Dig{meters: meters, direction: split[0]})
	}
	return digs
}
