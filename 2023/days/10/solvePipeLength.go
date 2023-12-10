package ten

import (
	"days"
	"fmt"
)

func GetDay() days.Day {
	return days.MakeDay(SolvePipeLength, Part2, "10")
}

type Point struct {
	x int
	y int
}

func SolvePipeLength(input []string) int {
	// input is kinda already parsed.
	startPoint := getStart(input)
	length := getRestOfPipeLength(input, startPoint, startPoint, startPoint)
	return length / 2
}

func Part2(input []string) int {
	start := getStart(input)
	path := []Point{}
	trappedPoints := 0
	path = getPointsInPipe(input, start, start, path)
	fmt.Println(path)
	pathMap := getPathMap(input, path)
	printNestedBool(pathMap)
	for y, line := range input {
		for x := range line {
			if pathMap[y][x] {
				continue
			}
			p := Point{x: x, y: y}
			visited := boolArrayOfSize(input)
			if isPointTrappedByLoop(input, p, pathMap, visited) {
				trappedPoints++
			}
		}
	}
	return trappedPoints
}

func isPointTrappedByLoop(maze []string, p Point, loop [][]bool, visited [][]bool) bool {
	// fmt.Println("Evaluating", p)
	if p.isOnEdge(maze) {
		return false
	}
	isTrapped := true
	visited[p.y][p.x] = true
	neighbors := getNeighbors(maze, p)
	for _, neighbor := range neighbors {
		if visited[neighbor.y][neighbor.x] {
			fmt.Println(neighbor, "already visited")
			continue // already evaluated.
		}
		if loop[neighbor.y][neighbor.x] {
			fmt.Println(neighbor, "on loop")
			continue // on loop.
		}
		neighborTrapped := isPointTrappedByLoop(maze, neighbor, loop, visited)
		if !neighborTrapped {
			isTrapped = neighborTrapped
			break
		}
	}
	visited[p.y][p.x] = false
	return isTrapped
}

func getNeighbors(maze []string, p Point) []Point {
	var neighbors []Point
	abovePoint := Point{x: p.x, y: p.y - 1}
	belowPoint := Point{x: p.x, y: p.y + 1}
	rightPoint := Point{x: p.x + 1, y: p.y}
	leftPoint := Point{x: p.x - 1, y: p.y}
	if p.x != 0 {
		neighbors = append(neighbors, leftPoint)
	}
	if p.y != 0 {
		neighbors = append(neighbors, abovePoint)
	}
	if p.y != len(maze)-1 {
		neighbors = append(neighbors, belowPoint)
	}
	if p.x != len(maze[p.y])-1 {
		neighbors = append(neighbors, rightPoint)
	}
	return neighbors
}

func boolArrayOfSize(maze []string) [][]bool {
	traversal := make([][]bool, len(maze))
	for i := 0; i < len(maze); i++ {
		traversal[i] = make([]bool, len(maze[i]))
	}
	return traversal
}

func (p Point) isOnEdge(maze []string) bool {
	return p.y == len(maze)-1 || p.y == 0 || p.x == 0 || p.x == len(maze[p.y])-1
}

func getRestOfPipeLength(maze []string, prev Point, current Point, start Point) int {
	if prev != start && current == start {
		return 0 // we are at the start and have traversed through the pipe.
	}
	nextPoint := getNextStep(maze, prev, current)
	return 1 + getRestOfPipeLength(maze, current, nextPoint, start)
}

func getPointsInPipe(maze []string, current Point, start Point, traversed []Point) []Point {
	fmt.Println("Traversed:", traversed, "Current", current)
	if len(traversed) != 0 && current == start {
		return traversed
	}
	lastPoint := current
	if len(traversed) > 0 {
		lastPoint = traversed[len(traversed)-1]
	}
	nextPoint := getNextStep(maze, lastPoint, current)
	fmt.Println("Next point", nextPoint)
	return getPointsInPipe(maze, nextPoint, start, append(traversed, current))
}

func getPathMap(maze []string, path []Point) [][]bool {
	// should be a copy of maze but with false booleans, rather than
	traversal := boolArrayOfSize(maze)
	for _, point := range path {
		traversal[point.y][point.x] = true
	}
	return traversal
}

func getNextStep(maze []string, prev Point, current Point) Point {
	connector1, connector2 := getConnectingPoints(maze, current)
	if connector1 == prev {
		return connector2
	}
	return connector1
}

func getConnectingPoints(maze []string, current Point) (Point, Point) {
	currentChar := maze[current.y][current.x]
	abovePoint := Point{x: current.x, y: current.y - 1}
	belowPoint := Point{x: current.x, y: current.y + 1}
	rightPoint := Point{x: current.x + 1, y: current.y}
	leftPoint := Point{x: current.x - 1, y: current.y}
	if currentChar == 'S' {
		connectingPoint := firstConnectingPoint(maze, current)
		return connectingPoint, connectingPoint
	}
	if currentChar == '|' {
		return belowPoint, abovePoint
	}
	if currentChar == '-' {
		return leftPoint, rightPoint
	}
	if currentChar == 'L' {
		return abovePoint, rightPoint
	}
	if currentChar == 'J' {
		return abovePoint, leftPoint
	}
	if currentChar == '7' {
		return belowPoint, leftPoint
	}
	if currentChar == 'F' {
		return belowPoint, rightPoint
	}
	panic("Unrecognized character: " + string(currentChar))
}

func firstConnectingPoint(maze []string, start Point) Point {
	// infer the connection based on what is near.
	if start.x != 0 {
		leftChar := maze[start.y][start.x-1]
		if leftChar == 'F' || leftChar == 'L' || leftChar == '-' {
			return Point{y: start.y, x: start.x - 1}
		}
	}
	if start.x != len(maze[start.y])-1 {
		rightChar := maze[start.y][start.x+1]
		if rightChar == '-' || rightChar == 'J' || rightChar == '7' {
			return Point{y: start.y, x: start.x + 1}
		}
	}
	if start.y != 0 {
		aboveChar := maze[start.y-1][start.x]
		if aboveChar == '|' || aboveChar == '7' || aboveChar == 'F' {
			return Point{y: start.y - 1, x: start.x}
		}
	}
	if start.y != len(maze)-1 {
		belowChar := maze[start.y+1][start.x]
		if belowChar == '|' || belowChar == 'J' || belowChar == 'L' {
			return Point{y: start.y + 1, x: start.x}
		}
	}
	panic("No connectors detected")
}

func getStart(input []string) Point {
	for y, line := range input {
		for x, char := range line {
			if char == 'S' {
				return Point{x: x, y: y}
			}
		}
	}
	panic("Did not find start")
}

func printNestedBool(visited [][]bool) {
	for _, line := range visited {
		fmt.Println(line)
	}
}
