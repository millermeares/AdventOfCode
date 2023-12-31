package eight

type Cycle struct {
	originalNode          Node
	stepsBeforeCycleEntry int
	length                int
}

// intention is to return the index at which a Z is encountered, starting at c.originalNode.
func (c Cycle) zInCycle(graph map[string]Node, steps string) []int {
	var zIdxs []int
	currentNode := c.originalNode
	for i := c.stepsBeforeCycleEntry; i < c.stepsBeforeCycleEntry+c.length; i++ {
		if currentNode.endsWithZ() {
			zIdxs = append(zIdxs, i-c.stepsBeforeCycleEntry)
		}
		stepIdx := i % len(steps)
		nextNodeId := currentNode.getStepNodeId(steps[stepIdx])
		currentNode = graph[nextNodeId]
	}
	return zIdxs
}

func getCycle(startNode Node, graph map[string]Node, steps string) Cycle {
	currentNode := startNode
	visited := map[string][]Visit{} // record "steps" at which the node was visited
	stepCount := 0
	for ; true; stepCount++ {
		stepIdx := stepCount % len(steps)
		if contains(visited[currentNode.id], stepIdx) {
			break
		}
		visited[currentNode.id] = append(visited[currentNode.id], Visit{stepIndex: stepIdx, stepCount: stepCount})
		nextNodeId := currentNode.getStepNodeId(steps[stepIdx])
		currentNode = graph[nextNodeId]
	}
	stepIndex := stepCount % len(steps)
	stepsBeforeCycleEntry := visitForIndex(visited[currentNode.id], stepIndex).stepCount
	return Cycle{
		originalNode:          graph[currentNode.id],
		stepsBeforeCycleEntry: stepsBeforeCycleEntry,
		length:                stepCount - stepsBeforeCycleEntry,
	}
}

type Visit struct {
	stepIndex int
	stepCount int
}

func visitForIndex(list []Visit, num int) Visit {
	for _, v := range list {
		if v.stepIndex == num {
			return v
		}
	}
	panic("Visit should exist")
}

func contains(list []Visit, num int) bool {
	for _, v := range list {
		if v.stepIndex == num {
			return true
		}
	}
	return false
}
