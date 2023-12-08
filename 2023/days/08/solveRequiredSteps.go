package eight

import (
	"days"
)

func GetDay() days.Day {
	return days.MakeDay(SolveRequiredSteps, Part2, "08")
}

func SolveRequiredSteps(input []string) int {
	steps := input[0]
	if input[1] != "" {
		panic("misunderstood input" + input[1])
	}
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
	return -1
}
