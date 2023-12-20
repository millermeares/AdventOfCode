package seventeen

import (
	"days"
)

func GetDay() days.Day {
	return days.MakeDay(Part1, Part2, "17")
}

func Part1(input []string) int {
	return dijkstra(input)
}

type Point struct {
	x int
	y int
}

type PathToPoint struct {
	straightLineLength int
	curPoint           Point
	previousPoint      Point
	costFromStart      int
}

func removeEquivalentPoint(points []Point, toRemove Point) []Point {
	idx := -1
	for i, p := range points {
		if p == toRemove {
			idx = i
		}
	}
	if idx == -1 {
		return points
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

func (p Point) adjacentPoints(maze [][]int) []Point {
	var points []Point
	if p.y != 0 {
		points = append(points, Point{x: p.x, y: p.y - 1})
	}
	if p.y != len(maze)-1 {
		points = append(points, Point{x: p.x, y: p.y + 1})
	}
	if p.x != 0 {
		points = append(points, Point{x: p.x - 1, y: p.y})
	}
	if p.x != len(maze[p.y])-1 {
		points = append(points, Point{x: p.x + 1, y: p.y})
	}
	return points
}

func Part2(input []string) int {
	return -1
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
