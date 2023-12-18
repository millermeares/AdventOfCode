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
	lines := toLines(digs)
	for _, line := range lines {
		fmt.Println(line)
	}
	return countEnclosed(lines)
}

func Part2(input []string) int {
	digs := parseHexadecimalInput(input)
	lines := toLines(digs)
	return countEnclosed(lines)
}

func countEnclosed(lines []Line) int {
	count := 0
	minX, maxX := xRange(lines)
	minY, maxY := yRange(lines)
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			p := Point{x: x, y: y}
			fmt.Println("Evaluating", p)
			if pointEnclosed(p, maxX, lines) {
				count++
			}
		}
	}
	return count
}

func pointEnclosed(p Point, maxX int, lines []Line) bool {
	linesCrossed := 0
	endPoint := Point{y: p.y, x: maxX + 1}
	lineToEdge := Line{start: p, end: endPoint}
	for _, line := range lines {
		if line.isOnLine(p) {
			return true // if on line, functionally enclosed. horizontal should be handled?
		}
		if !line.isVertical() {
			continue // only count vertical lines that we crossed.
		}
		if line.crossesLine(lineToEdge) {
			// fmt.Println(lineToEdge, "crosses line", line)
			linesCrossed++
		} else {
			// fmt.Println(line, "does not cross", lineToEdge)
		}
	}
	// fmt.Println("Crosses", linesCrossed, lineToEdge)
	return linesCrossed%2 != 0
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
		end:   Point{x: cur.x + (xChange * d.meters), y: cur.y + (yChange * d.meters)},
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

func (l Line) minX() int {
	if l.start.x > l.end.x {
		return l.end.x
	}
	return l.start.x
}

func (l Line) maxX() int {
	if l.start.x > l.end.x {
		return l.start.x
	}
	return l.end.x
}

func (l Line) minY() int {
	if l.start.y > l.end.y {
		return l.end.y
	}
	return l.start.y
}

func (l Line) maxY() int {
	if l.start.y < l.end.y {
		return l.end.y
	}
	return l.start.y
}

func xRange(lines []Line) (int, int) {
	minX := math.MaxInt
	maxX := math.MinInt
	for _, line := range lines {
		lineMinX := line.minX()
		lineMaxX := line.maxX()
		if minX > lineMinX {
			minX = lineMinX
		}
		if lineMaxX > maxX {
			maxX = lineMaxX
		}
	}
	return minX, maxX
}

func yRange(lines []Line) (int, int) {
	minY := math.MaxInt
	maxY := math.MinInt
	for _, line := range lines {
		lineMinY := line.minY()
		lineMaxX := line.maxY()
		if minY > lineMinY {
			minY = lineMinY
		}
		if lineMaxX > maxY {
			maxY = lineMaxX
		}
	}
	return minY, maxY
}

func (l Line) isVertical() bool {
	return l.start.x == l.end.x
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

// smaller y
func (l Line) top() Point {
	if l.start.y > l.end.y {
		return l.end
	}
	return l.start
}

// bigger y
func (l Line) bottom() Point {
	if l.start.y > l.end.y {
		return l.start
	}
	return l.end
}

// smaller x
func (l Line) left() Point {
	if l.start.x > l.end.x {
		return l.end
	}
	return l.start
}

// bigger x
func (l Line) right() Point {
	if l.start.x > l.end.x {
		return l.start
	}
	return l.end
}

func (vertical Line) crossesLine(horizontal Line) bool {
	// horizontal Y should be between bottom Y(inclusive) and top Y (exclusive)
	if !(horizontal.start.y <= vertical.bottom().y && horizontal.start.y > vertical.top().y) {
		return false
	}

	// vertical X should be between horizontal left() and right()
	if !(vertical.start.x >= horizontal.left().x && vertical.start.x < horizontal.right().x) {
		return false
	}
	return true
}

func (l Line) isOnLine(p Point) bool {
	if l.isVertical() {
		if p.x != l.start.x {
			return false
		}
		return l.bottom().y >= p.y && l.top().y <= p.y
	} else {
		// l is horizontal (y is same for each point)
		if p.y != l.start.y {
			return false
		}
		return l.left().x <= p.x && l.right().x >= p.x
	}
}
