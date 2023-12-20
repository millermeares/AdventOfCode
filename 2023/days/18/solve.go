package eighteen

import (
	"days"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func GetDay() days.Day {
	return days.MakeDay(Part1, Part2, "18")
}

func Part1(input []string) int {
	digs := parseInput(input)
	lines := toLines(digs)
	return countEnclosed(lines)
}

func Part2(input []string) int {
	digs := parseHexadecimalInput(input)
	lines := toLines(digs)
	return countEnclosed(lines)
}

func countEnclosed(lines []Line) int {
	count := 0
	horizontalLines := sortedHorizontal(lines)
	verticalLines := sortedVertical(lines)
	for len(horizontalLines) > 0 {
		topLine := horizontalLines[0]
		horizontalLines = horizontalLines[1:] // pop. add unmatched part at end.

		for i := 0; i < len(horizontalLines); i++ {
			fmt.Println(horizontalLines)
			if !linesOverlap(topLine, horizontalLines[i]) {
				fmt.Println(topLine, "does not overlap with", horizontalLines[i])
				continue
			}
			bottomLine := horizontalLines[i]
			horizontalLines = append(horizontalLines[:i], horizontalLines[i+1:]...) // remove bottom line.

			// the problem with this approach is that there can sometimes be parts of vertical lines that are missed if below a "bottomLine" on the right side. (i think)
			yDiff := bottomLine.start.y - topLine.start.y + 1
			overlappingX := overlappingX(topLine, bottomLine) + 1
			areaToAdd := overlappingX * yDiff

			fmt.Println("Adding", areaToAdd, "as a result of overlap between", topLine, "and", bottomLine)
			count += areaToAdd
			descendingVerticalLineLength := descendingVerticalLineAtCornerLength(bottomLine, verticalLines) // this is not a valid technique.
			count += descendingVerticalLineLength
			leftoverLines := subtractOverlappingX(topLine, bottomLine)
			fmt.Println("Left with", leftoverLines, "when subtracting overlap from", topLine, bottomLine)
			horizontalLines = append(horizontalLines, leftoverLines...)
			break
		}
		horizontalLines = sortedHorizontal(horizontalLines)
	}
	return count
}

func descendingVerticalLineAtCornerLength(bottomLine Line, verticalLines []Line) int {
	for _, vertical := range verticalLines {
		if vertical.top() == bottomLine.start || vertical.top() == bottomLine.end {
			return vertical.verticalLength() - 2 // inclusive length, don't count corners
		}
	}
	return 0
}

func subtractOverlappingX(l1 Line, l2 Line) []Line {
	start, end := getOverlappingStartEndX(l1, l2)
	remaining := removeSnippet(l1, start, end)
	remaining = append(remaining, removeSnippet(l2, start, end)...)
	return remaining
}

// assume that both start and end of line are inclusive.
// what are limitations of this strategy?
func removeSnippet(l Line, xStart, xEnd int) []Line {
	var remaining []Line
	if xStart != l.minX() {
		remaining = append(remaining, Line{
			start: l.start,
			end:   Point{y: l.end.y, x: xStart - 1},
		})
	}

	if xEnd != l.maxX() {
		remaining = append(remaining, Line{
			start: Point{y: l.start.y, x: xEnd + 1},
			end:   l.end,
		})
	}

	return remaining
}

func getOverlappingStartEndX(l1 Line, l2 Line) (int, int) {
	left := l1
	right := l2
	if l2.minX() < l1.minX() {
		left = l2
		right = l1
	}
	start := right.minX()
	end := int(math.Min(float64(right.maxX()), float64(left.maxX())))
	return start, end
}

func linesOverlap(l1 Line, l2 Line) bool {
	return overlappingX(l1, l2) >= 0
}

func overlappingX(l1 Line, l2 Line) int {
	start, end := getOverlappingStartEndX(l1, l2)
	return end - start
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
	hexLength = strings.Replace(hexLength, "#", "", -1)
	parsed, e := strconv.ParseInt(hexLength, 16, 64)
	if e != nil {
		panic(e)
	}
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

func (l Line) top() Point {
	above := l.start
	if l.start.y > l.end.y { // smaller is higher.
		above = l.end
	}
	return above
}

func (l Line) bottom() Point {
	above := l.start
	if l.start.y < l.end.y { // bigger is lower.
		above = l.end
	}
	return above
}

func (l Line) verticalLength() int {
	return l.bottom().y - l.top().y + 1
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

func sortedHorizontal(lines []Line) []Line {
	var horizontal []Line
	for _, line := range lines {
		if !line.isVertical() {
			horizontal = append(horizontal, line)
		}
	}
	sort.SliceStable(horizontal, func(i, j int) bool {
		return horizontal[i].start.y < horizontal[j].end.y
	})
	return horizontal
}

func sortedVertical(lines []Line) []Line {
	var vertical []Line
	for _, line := range lines {
		if line.isVertical() {
			vertical = append(vertical, line)
		}
	}
	sort.SliceStable(vertical, func(i, j int) bool {
		return vertical[i].start.x < vertical[j].end.x
	})
	return vertical
}
