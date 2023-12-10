package eight

import (
	"days"
	"fmt"
)

func GetDay() days.Day {
	return days.MakeDay(SolveRequiredSteps, Part2, "08")
}

func SolveRequiredSteps(input []string) int {
	steps := input[0]
	graph := parseInput(input[2:])
	currentNode := graph["AAA"]
	stepCount := 0
	for ; currentNode.id != "ZZZ"; stepCount++ {
		stepToTake := steps[stepCount%len(steps)] // stepCount mod len(steps)
		nextNodeId := currentNode.getStepNodeId(stepToTake)
		currentNode = graph[nextNodeId]
	}
	return stepCount
}

func Part2(input []string) int {
	// start at every path that ends with "A"?
	steps := input[0]
	graph := parseInput(input[2:])

	nodes := getNodesEndingWithA(graph)
	var cycles []Cycle
	for _, node := range nodes {
		c := getCycle(node, graph, steps)
		fmt.Println(c)
		fmt.Println(c.zInCycle(graph, steps))
		cycles = append(cycles, c)
	}

	return getMinOverlappingZ(cycles, graph, steps)
}

func getMinOverlappingZ(cycles []Cycle, graph map[string]Node, steps string) int {
	var rates []int
	maxOffset := 0
	for _, c := range cycles {
		zIndex := c.zInCycle(graph, steps)[0] // hard-coded to only have 1 Z per cycle. I know this is true from my input but is hacky.
		zOffset := zIndex + c.stepsBeforeCycleEntry
		if zOffset > maxOffset {
			maxOffset = zOffset
		}
		rates = append(rates, c.length)
	}
	fmt.Println("Calculating LCM of rates", rates)
	fmt.Println("Adding max offset", maxOffset)
	return LCM(rates[0], rates[1], rates[2:]) + maxOffset
}

func getNodesEndingWithA(graph map[string]Node) []Node {
	var nodes []Node
	for _, v := range graph {
		if v.endsWithA() {
			nodes = append(nodes, v)
		}
	}
	return nodes
}

// copy-pasted from https://go.dev/play/p/SmzvkDjYlb and amended
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers []int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i], integers[i+1:])
	}

	return result
}
