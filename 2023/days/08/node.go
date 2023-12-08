package eight

import "strings"

type Node struct {
	id          string
	leftNodeId  string
	rightNodeId string
}

func (n *Node) getStepNodeId(step byte) string {
	if step != 'R' && step != 'L' {
		panic("Unexpected input: " + string(step))
	}
	if step == 'R' {
		return n.rightNodeId
	}
	return n.leftNodeId
}

func parseInput(input []string) map[string]Node {
	graph := map[string]Node{}
	for _, line := range input {
		node := parseNode(line)
		graph[node.id] = node
	}
	return graph
}

func parseNode(input string) Node {
	split := strings.Split(input, " = ")
	id := split[0]
	rl := strings.Replace(split[1], "(", "", -1)
	rl = strings.Replace(rl, ")", "", -1)
	rl = strings.Replace(rl, " ", "", -1)
	both := strings.Split(rl, ",")
	return Node{
		id:          id,
		leftNodeId:  both[0],
		rightNodeId: both[1],
	}
}
