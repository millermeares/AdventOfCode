package seventeen

import (
	"math"
)

// technically this is SPFA, but same idea as dijkstra's - just less complicated.
func dijkstra(input []string, minStraight int, maxStraight int) int {
	maze := initializeMaze(input, maxStraight)
	mazeWeights := toIntArray(input)
	start := maze[Point{x: 0, y: 0}]
	// the paths to points are not clear.
	startPathToPoint := start.getPathToPoint(Point{x: 0, y: 0}, 0)
	startPathToPoint.costFromStart = 0 // start here, it is free.
	queue := []*PathToPoint{startPathToPoint}

	for len(queue) > 0 {
		first := queue[0]
		queue = queue[1:] // remove first
		neighbors := first.getNeighbors(maze, mazeWeights, minStraight, maxStraight)
		for _, neighbor := range neighbors {
			if neighbor.costFromStart <= first.costFromStart+neighbor.myPointWeight(maze) {
				continue // won't update minimum value for this route, so end.
			}
			neighbor.costFromStart = first.costFromStart + neighbor.myPointWeight(maze)
			queue = append(queue, neighbor)
		}
	}

	yEnd := len(input) - 1
	end := Point{y: yEnd, x: len(input[yEnd]) - 1}
	return maze[end].minCostFromStart(minStraight)
}

func initializeMaze(input []string, maxStraight int) map[Point]AggregatePoint {
	weights := toIntArray(input)
	points := map[Point]AggregatePoint{}
	for y := 0; y < len(weights); y++ {
		for x := 0; x < len(weights[y]); x++ {
			p := Point{x: x, y: y}
			points[p] = makeAggregatePoint(p, weights, maxStraight)
		}
	}
	return points
}

type AggregatePoint struct {
	p          Point
	pointPaths []*PathToPoint
	weight     int
}

func makeAggregatePoint(p Point, weights [][]int, maxStraightLength int) AggregatePoint {
	var p2ps []*PathToPoint
	for i := 1; i <= maxStraightLength; i++ {
		p2ps = append(p2ps, &PathToPoint{previousPoint: Point{x: p.x - 1, y: p.y}, straightLineLength: i, costFromStart: math.MaxInt, curPoint: p})
		p2ps = append(p2ps, &PathToPoint{previousPoint: Point{x: p.x + 1, y: p.y}, straightLineLength: i, costFromStart: math.MaxInt, curPoint: p})
		p2ps = append(p2ps, &PathToPoint{previousPoint: Point{x: p.x, y: p.y - 1}, straightLineLength: i, costFromStart: math.MaxInt, curPoint: p})
		p2ps = append(p2ps, &PathToPoint{previousPoint: Point{x: p.x, y: p.y + 1}, straightLineLength: i, costFromStart: math.MaxInt, curPoint: p})
	}
	// this should only really be used at the start node when no paths have happened yet?
	p2ps = append(p2ps, &PathToPoint{previousPoint: Point{x: p.x, y: p.y}, straightLineLength: 0, costFromStart: math.MaxInt})
	return AggregatePoint{p: p, pointPaths: p2ps, weight: weights[p.y][p.x]}
}

func (ap AggregatePoint) minCostFromStart(minStraight int) int {
	min := math.MaxInt
	for _, p2p := range ap.pointPaths {
		if p2p.straightLineLength < minStraight {
			continue // must travel at least this far.
		}
		if p2p.costFromStart < min {
			min = p2p.costFromStart
		}
	}
	return min
}

func (ap AggregatePoint) getPathToPoint(prev Point, straightLineLength int) *PathToPoint {
	for _, p2p := range ap.pointPaths {
		if p2p.previousPoint == prev && straightLineLength == p2p.straightLineLength {
			return p2p
		}
	}
	panic("Could not find path to point")
}

func (p2p *PathToPoint) getNeighbors(maze map[Point]AggregatePoint, mazeWeights [][]int, minStraight int, maxStraight int) []*PathToPoint {
	adjacent := p2p.curPoint.adjacentPoints(mazeWeights)
	adjacent = removeEquivalentPoint(adjacent, p2p.previousPoint)

	var neighbors []*PathToPoint
	for _, pot := range adjacent {
		if pointsAreInStraightLine([]Point{p2p.previousPoint, p2p.curPoint, pot}) {
			if p2p.straightLineLength == maxStraight { // not valid
				continue
			}
			neighbors = append(neighbors, maze[pot].getPathToPoint(p2p.curPoint, p2p.straightLineLength+1))
			continue
		}
		// we can only turn if we have traveled >= minStraight.
		if p2p.straightLineLength >= minStraight {
			neighbors = append(neighbors, maze[pot].getPathToPoint(p2p.curPoint, 1))
		}
	}
	return neighbors
}

func (p2p PathToPoint) myPointWeight(maze map[Point]AggregatePoint) int {
	return maze[p2p.curPoint].weight
}
