package twentyone

import (
	"days"
)

func GetDay() days.Day {
	return days.MakeDay(Part1, Part2, "21")
}

func Part1(input []string) int {
	return countViableSpots(input, 64, false)
}

func countViableSpots(maze []string, steps int, allowDimensions bool) int {
	memo := map[Point]map[int][]Point{}
	start := findStartPoint(maze)
	return len(getPossibleLandingSpots(start, maze, steps, memo, allowDimensions))
}

func findStartPoint(maze []string) Point {
	for y := 0; y < len(maze); y++ {
		for x := 0; x < len(maze[y]); x++ {
			if maze[y][x] == 'S' {
				return Point{x: x, y: y, dimension: 0}
			}
		}
	}
	panic("Couldn't find start")
}

func getPossibleLandingSpots(p Point, maze []string, remainingSteps int, memo map[Point]map[int][]Point, allowDimensions bool) []Point {
	if remainingSteps == 1 {
		return getViableNeighbors(p, maze, allowDimensions)
	}
	_, e := memo[p]
	if !e {
		memo[p] = map[int][]Point{}
	}
	stepsCalculated := memo[p]
	stepCalculations, thisStepCalculated := stepsCalculated[remainingSteps]
	if thisStepCalculated {
		return stepCalculations
	}

	neighbors := getViableNeighbors(p, maze, allowDimensions)
	distinct := map[Point]bool{}
	for _, neighbor := range neighbors {
		neighborSpots := getPossibleLandingSpots(neighbor, maze, remainingSteps-1, memo, allowDimensions)
		for _, spot := range neighborSpots {
			distinct[spot] = true
		}
	}

	var distinctSpots []Point
	for k := range distinct {
		distinctSpots = append(distinctSpots, k)
	}
	memo[p][remainingSteps] = distinctSpots
	return distinctSpots
}

func getViableNeighbors(p Point, maze []string, allowDimensions bool) []Point {
	var points []Point
	if p.x != 0 {
		if maze[p.y][p.x-1] != '#' {
			points = append(points, Point{x: p.x - 1, y: p.y, dimension: p.dimension})
		}
	} else if allowDimensions {
		if maze[p.y][len(maze[p.y])-1] != '#' {
			points = append(points, Point{
				dimension: p.dimension - 1,
				x:         len(maze[p.y]) - 1,
				y:         p.y,
			})
		}
	}
	if p.y != 0 {
		if maze[p.y-1][p.x] != '#' {
			points = append(points, Point{x: p.x, y: p.y - 1, dimension: p.dimension})
		}
	} else if allowDimensions {
		if maze[len(maze)-1][p.x] != '#' {
			points = append(points, Point{
				dimension: p.dimension - 1,
				x:         p.x,
				y:         len(maze) - 1,
			})
		}
	}
	if p.y != len(maze)-1 {
		if maze[p.y+1][p.x] != '#' {
			points = append(points, Point{x: p.x, y: p.y + 1, dimension: p.dimension})
		}
	} else if allowDimensions {
		if maze[0][p.x] != '#' {
			points = append(points, Point{
				dimension: p.dimension + 1,
				x:         p.x,
				y:         0,
			})
		}
	}
	if p.x != len(maze[p.y])-1 {
		if maze[p.y][p.x+1] != '#' {
			points = append(points, Point{x: p.x + 1, y: p.y, dimension: p.dimension})
		}
	} else if allowDimensions {
		// wrap around to next dimension.
		if maze[p.y][0] != '#' {
			points = append(points, Point{
				dimension: p.dimension + 1,
				x:         0,
				y:         p.y,
			})
		}
	}

	return points
}

func Part2(input []string) int {
	return countViableSpots(input, 26501365, true)
}

// in theory, i think i could memoize without the dimension (maybe). calculation should be the same across dimensions from the same square
type Point struct {
	x         int
	y         int
	dimension int
}
