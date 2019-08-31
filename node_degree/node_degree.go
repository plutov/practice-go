package nodedegree

import (
	"fmt"
)

// Degree func
func Degree(nodes int, graph [][2]int, node int) (int, error) {
	if node > nodes {
		return 0, fmt.Errorf("node %d not found in the graph", node)
	}

	degree := 0

	for _, n := range graph {
		if n[0] == node || n[1] == node {
			degree++
		}
	}

	return degree, nil
}
