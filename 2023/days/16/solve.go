package sixteen

import (
	"days"
)

func GetDay() days.Day {
	return days.MakeDay(Part1, Part2, "16")
}

type Point struct {
	x int
	y int
}

func Part1(input []string) int {
	leftOfStart := Point{x: -1, y: 0} // establish direction
	start := Point{x: 0, y: 0}
	return countEnergized(leftOfStart, start, input)
}

func countEnergized(beforeStart Point, start Point, input []string) int {
	energized := boolArrayOfSize(input)
	visited := map[Point][]Point{}
	traverse(beforeStart, start, input, energized, visited)
	return countTrue(energized)
}

// need to record visited and not visit if i have already.
func traverse(prev Point, cur Point, maze []string, energized [][]bool, visited map[Point][]Point) {
	// if "next" is not on the grid, stop.
	if !cur.isOnGrid(maze) {
		return // light reflected out of maze.
	}
	// populate if not there yet.
	_, prevBefore := visited[prev]
	if !prevBefore {
		visited[prev] = make([]Point, 0)
	}
	froms := visited[prev]
	seenBefore := contains(froms, cur)
	if seenBefore {
		// already traversed in this direction, we can stop here.
		return
	} else {
		froms = append(froms, cur)
		visited[prev] = froms
	}

	energized[cur.y][cur.x] = true
	xDiff, yDiff := prev.getDiff(cur)
	currentChar := maze[cur.y][cur.x]
	if currentChar == '.' || (currentChar == '|' && xDiff == 0) || (currentChar == '-' && yDiff == 0) {
		// continue in same direction.
		next := Point{
			x: cur.x + xDiff,
			y: cur.y + yDiff,
		}
		traverse(cur, next, maze, energized, visited)
		return
	}
	// splitter case can cause up and down or left and right.
	if currentChar == '|' && xDiff != 0 {
		upPoint := Point{x: cur.x, y: cur.y - 1}
		downPoint := Point{x: cur.x, y: cur.y + 1}
		traverse(cur, upPoint, maze, energized, visited)
		traverse(cur, downPoint, maze, energized, visited)
		return
	}
	if currentChar == '-' && yDiff != 0 {
		traverse(cur, Point{y: cur.y, x: cur.x - 1}, maze, energized, visited)
		traverse(cur, Point{y: cur.y, x: cur.x + 1}, maze, energized, visited)
		return
	}

	nextPoint := getNextPoint(prev, cur, rune(currentChar))
	traverse(cur, nextPoint, maze, energized, visited)
}

func getNextPoint(prev Point, cur Point, currentChar rune) Point {
	xDiff, yDiff := prev.getDiff(cur)
	if currentChar == '/' {
		if xDiff > 0 {
			return Point{x: cur.x, y: cur.y - 1}
		}
		if xDiff < 0 {
			return Point{x: cur.x, y: cur.y + 1}
		}
		if yDiff > 0 {
			return Point{y: cur.y, x: cur.x - 1}
		}
		if yDiff < 0 {
			return Point{y: cur.y, x: cur.x + 1}
		}
	}
	if currentChar == '\\' {
		if xDiff > 0 {
			return Point{x: cur.x, y: cur.y + 1}
		}
		if xDiff < 0 {
			return Point{x: cur.x, y: cur.y - 1}
		}
		if yDiff > 0 {
			return Point{y: cur.y, x: cur.x + 1}
		}
		if yDiff < 0 {
			return Point{y: cur.y, x: cur.x - 1}
		}
	}

	panic("Evaluating unexpected character" + string(currentChar))
}

func (p *Point) isOnGrid(maze []string) bool {
	return !(p.x < 0 || p.y < 0 || p.y >= len(maze) || p.x >= len(maze[p.y]))
}

// return xDiff, yDiff
func (prev *Point) getDiff(cur Point) (int, int) {
	return (cur.x - prev.x), (cur.y - prev.y)
}

func countTrue(energized [][]bool) int {
	sum := 0
	for _, line := range energized {
		for _, v := range line {
			if v {
				sum++
			}
		}
	}
	return sum
}

func boolArrayOfSize(input []string) [][]bool {
	traversal := make([][]bool, len(input))
	for i := 0; i < len(input); i++ {
		traversal[i] = make([]bool, len(input[i]))
	}
	return traversal
}

func contains(points []Point, finding Point) bool {
	for _, p := range points {
		if p == finding {
			return true
		}
	}
	return false
}

func Part2(input []string) int {
	max := 0
	// check left going right and right going left.
	for y, line := range input {
		start := Point{x: 0, y: y}
		leftOfStart := Point{x: -1, y: y}
		iter := countEnergized(leftOfStart, start, input)
		if iter > max {
			max = iter
		}

		start = Point{x: len(line) - 1, y: y}
		rightOfStart := Point{x: len(line), y: y}
		iter = countEnergized(rightOfStart, start, input)
		if iter > max {
			max = iter
		}
	}

	for x := range input[0] {
		start := Point{x: x, y: 0}
		aboveStart := Point{x: x, y: -1}
		iter := countEnergized(aboveStart, start, input)
		if iter > max {
			max = iter
		}
		start = Point{x: x, y: len(input) - 1}
		belowStart := Point{x: x, y: len(input)}
		iter = countEnergized(belowStart, start, input)
		if iter > max {
			max = iter
		}
	}
	return max
}
