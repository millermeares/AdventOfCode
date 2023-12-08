package eight

import (
	"days"
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
	stepCount := 0
	for ; !allNodesEndInZ(nodes); stepCount++ {
		stepToTake := steps[stepCount%len(steps)] // stepCount mod len(steps)
		for i := 0; i < len(nodes); i++ {
			nextNodeId := nodes[i].getStepNodeId(stepToTake)
			nodes[i] = graph[nextNodeId]
		}
	}
	// for each A node to Z node, can i answer the question: "how many steps until it gets to a Z
	return stepCount
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
	for _, node := range nodes {
		if !node.endsWithZ() {
			return false
		}
	}
	return true
}
