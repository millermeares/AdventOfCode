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
	fmt.Println(allNodesEndInZ(nodes))
	var cycles []*Cycle
	for _, node := range nodes {
		c := getCycle(node, graph, steps)
		cycles = append(cycles, &c)
	}
	printCycles(cycles, graph, steps)

	stepsBeforeCycleEntry := makeCyclesStartWhenLatestStarts(cycles, graph, steps)
	lcm := getLcmConsiderPossibleZs(cycles, graph, steps)
	return lcm + stepsBeforeCycleEntry
}

func getLcmConsiderPossibleZs(cycles []*Cycle, graph map[string]Node, steps string) int {
	var zs []int
	for _, c := range cycles {
		zs = append(zs, c.zInCycle(graph, steps)[0]) // hard-coded to only have 1 Z per cycle. I know this is true from my input but is hacky.
	}
	// LCM of just the Zs isn't enough.
	// that would be enough if the Z was reliably the last element of each cycle.
	return LCM(zs[0], zs[1], zs[2:])

}

func printCycles(cycles []*Cycle, graph map[string]Node, steps string) {
	for _, c := range cycles {
		fmt.Println(c)
		fmt.Println(c.zInCycle(graph, steps))
	}
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

func allNodesEndInZ(nodes []Node) bool {
	zCount := 0
	for _, node := range nodes {
		if node.endsWithZ() {
			// fmt.Println("Node at index ending in Z", idx)
			zCount++
		}
	}
	return zCount == len(nodes)
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
