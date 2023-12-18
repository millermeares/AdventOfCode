package seventeen

import (
	"days"
	"math"
)

func GetDay() days.Day {
	return days.MakeDay(Part1, Part2, "17")
}

type Point struct {
	x int
	y int
}

type PathToPoint struct {
	straightLineLength int
	previousPoint      Point
}

func Part1(input []string) int {
	visited := boolArrayOfSize(input)
	maze := toIntArray(input)
	start := Point{x: 0, y: 0}
	path := []Point{start}
	visited[0][0] = true
	return cheapestPathDFS(path, maze, visited, map[Point]map[PathToPoint]int{})
}

// answer 882 and 800 are too high.
func cheapestPathDFS(path []Point, maze [][]int, visited [][]bool, memo map[Point]map[PathToPoint]int) int {
	last := path[len(path)-1]
	if last.isEnd(maze) {
		return 0
	}
	pathToPoint := getPathToPoint(path)
	lc, e := memo[last]
	if e {
		cost, lenExists := lc[pathToPoint]
		if lenExists {
			return cost
		}
	}

	children := possibleNextSteps(path, maze)
	min := math.MaxInt
	for _, child := range children {
		if visited[child.y][child.x] {
			continue
		}
		path = append(path, child)
		visited[child.y][child.x] = true
		amount := cheapestPathDFS(path, maze, visited, memo)
		visited[child.y][child.x] = false
		path = path[:len(path)-1]
		if amount == math.MaxInt {
			continue
		}
		amount += maze[child.y][child.x]
		if min > amount {
			min = amount
		}
	}
	lenCost, e := memo[last]
	if !e {
		lenCost = map[PathToPoint]int{}
	}
	lenCost[pathToPoint] = min
	memo[last] = lenCost
	return min
}

func (p Point) isEnd(maze [][]int) bool {
	return p.y == len(maze)-1 && p.x == len(maze[p.y])-1
}

func possibleNextSteps(path []Point, maze [][]int) []Point {
	currentPoint := path[len(path)-1]
	if len(path) == 1 {
		// we are at the beginning. the two legal options are down and right.
		return []Point{
			{x: currentPoint.x + 1, y: currentPoint.y},
			{x: currentPoint.x, y: currentPoint.y + 1},
		}
	}

	adjacentPoints := currentPoint.adjacentPoints(maze)
	// remove previous point as an option because we can't reverse.
	adjacentPoints = removeEquivalentPoint(adjacentPoints, path[len(path)-2])
	if len(path) >= 4 {
		if pointsAreInStraightLine(path[len(path)-4:]) {
			// if last 3 are straight line, don't allow "straight".
			adjacentPoints = removeStraightOption(path, adjacentPoints)
		}
	}

	return adjacentPoints
}

func removeStraightOption(path []Point, options []Point) []Point {
	last := path[len(path)-1]
	shortestPath := path[len(path)-2:]
	sX := sameX(shortestPath)
	sY := sameY(shortestPath)
	for _, option := range options {
		if sX && last.x == option.x {
			return removeEquivalentPoint(options, option)
		}
		if sY && last.y == option.y {
			return removeEquivalentPoint(options, option)
		}
	}
	return options
}

func removeEquivalentPoint(points []Point, toRemove Point) []Point {
	idx := -1
	for i, p := range points {
		if p == toRemove {
			idx = i
		}
	}
	return append(points[:idx], points[idx+1:]...)
}

func pointsAreInStraightLine(points []Point) bool {
	// fmt.Println("Evaluating if points are in straight line", points)
	return sameY(points) || sameX(points) // this works since there is no reversing.
}

func sameY(points []Point) bool {
	for i := 0; i < len(points)-1; i++ {
		if points[i].y != points[i+1].y {
			return false
		}
	}
	return true
}

func sameX(points []Point) bool {
	for i := 0; i < len(points)-1; i++ {
		if points[i].x != points[i+1].x {
			return false
		}
	}
	return true
}

func getPathToPoint(points []Point) PathToPoint {
	straightLineLength := straightLineLength(points)
	if len(points) <= 1 {
		return PathToPoint{
			straightLineLength: straightLineLength,
		}
	}
	return PathToPoint{
		straightLineLength: straightLineLength,
		previousPoint:      points[len(points)-2],
	}
}

func straightLineLength(points []Point) int {
	if len(points) == 1 {
		return 1
	}
	lastTwo := points[len(points)-2:]
	sX := sameX(lastTwo)
	sY := sameY(lastTwo)

	count := 1
	for i := len(points) - 1; i >= 1; i-- {
		if sX && points[i].x != points[i-1].x {
			break
		}
		if sY && points[i].y != points[i-1].y {
			break
		}
		count++
	}
	return count
}

func (p Point) adjacentPoints(maze [][]int) []Point {
	var points []Point
	if p.x != 0 {
		points = append(points, Point{x: p.x - 1, y: p.y})
	}
	if p.x != len(maze[p.y])-1 {
		points = append(points, Point{x: p.x + 1, y: p.y})
	}
	if p.y != 0 {
		points = append(points, Point{x: p.x, y: p.y - 1})
	}
	if p.y != len(maze)-1 {
		points = append(points, Point{x: p.x, y: p.y + 1})
	}
	return points
}

func Part2(input []string) int {
	return -1
}

func boolArrayOfSize(input []string) [][]bool {
	traversal := make([][]bool, len(input))
	for i := 0; i < len(input); i++ {
		traversal[i] = make([]bool, len(input[i]))
	}
	return traversal
}
func toIntArray(input []string) [][]int {
	rows := make([][]int, len(input))
	for i, line := range input {
		var nums []int
		for _, c := range line {
			n := int(c - '0')
			nums = append(nums, n)
		}
		rows[i] = nums
	}
	return rows
}
