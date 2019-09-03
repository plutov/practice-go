package nodedegree

import (
	"fmt"
	"math"
)

// Degree func
func Degree(nodes int, graph [][2]int, node int) (int, error) {
	return DegreeLinearReverse(nodes, graph, node)
}

// DegreeLinear search for occurences of node
// by interating the complete graph.
// Access nodes within the for loop by index.
func DegreeLinear(nodes int, graph [][2]int, node int) (int, error) {

	if node > nodes {
		return 0, fmt.Errorf("node %d not found in the graph", node)
	}

	degree := 0
	for i := range graph {
		if graph[i][0] == node || graph[i][1] == node {
			degree++
		}
	}

	return degree, nil
}

// DegreeLinearCopy search for occurences of node
// by interating the complete graph.
// Access nodes within the for loop by value.
func DegreeLinearCopy(nodes int, graph [][2]int, node int) (int, error) {

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

// DegreeLinearReverse search for occurences of node
// by interating the sorted graph in reverse order.
// Stop after the second end of a connection gets too small.
func DegreeLinearReverse(nodes int, graph [][2]int, node int) (int, error) {

	if node > nodes {
		return 0, fmt.Errorf("node %d not found in the graph", node)
	}

	degree := 0
	var i int
	for i = len(graph) - 1; i >= 0 && graph[i][1] > node; i-- {
		if graph[i][0] == node {
			degree++
		}
	}

	for j := i; j >= 0 && graph[j][1] == node; j-- {
		degree++
	}

	return degree, nil
}

// DegreeStepReverse search for occurences of node
// by skipping nodes in between. Step size is adjusted accordingly
// whether the last step was too short or too long.
func DegreeStepReverse(nodes int, graph [][2]int, node int) (int, error) {

	if node > nodes {
		return 0, fmt.Errorf("node %d not found in the graph", node)
	}

	lenGraph := len(graph)
	last := graph[lenGraph-1][1]

	step := lenGraph / last
	if step == 0 {
		step = 1
	}

	degree := 0

	next := last
	i := lenGraph - 1
	for ; next > node; i -= step {

		i, step = adjust(i, step, graph, next)
		if graph[i][1] != next {
			next--
			continue
		}

		i = find(i, graph, node)
		if graph[i][0] == node {
			degree++
		}

		next--
	}

	if i < 0 {
		return degree, nil
	}

	i, _ = adjust(i, 0, graph, node)
	for j := i + 1; j < len(graph) && graph[j][1] == node; j++ {
		degree++
	}
	for j := i; j >= 0 && graph[j][1] == node; j-- {
		degree++
	}
	return degree, nil
}

func adjust(i, step int, graph [][2]int, node int) (int, int) {
	if graph[i][1] == node {
		return i, step
	}
	if graph[i][1] > node {
		for i > 0 && graph[i][1] > node {
			i--
		}
		return i, step + 1
	}
	for i < len(graph)-1 && graph[i][1] < node {
		i++
	}
	return i, step - 1
}

func find(i int, graph [][2]int, node int) int {
	if graph[i][0] == node {
		return i
	}
	if graph[i][0] > node {
		for i > 0 && graph[i][0] > node {
			i--
		}
		return i
	}
	for i < len(graph)-1 && graph[i][0] < node {
		i++
	}
	return i
}

// DegreeInterpol interpolation search for occurences of node in graph.
func DegreeInterpol(nodes int, graph [][2]int, node int) (int, error) {

	if node > nodes {
		return 0, fmt.Errorf("node %d not found in the graph", node)
	}

	// shift factor needed to convert nodes to integers
	log2 := uint(math.Log2(float64(nodes))) + 1

	// search on the first end of connections
	end := len(graph)
	degree := 0
	for n := graph[len(graph)-1][1]; n > node; n-- {
		i := interpolSearch(graph[:end], [2]int{node, n}, log2)
		if graph[i][0] == node {
			degree++
		}
		end = i
	}

	// search on the second end of connections
	i := interpolSearch(graph, [2]int{node - 1, node}, log2)
	for j := i + 1; j < len(graph) && graph[j][1] == node; j++ {
		degree++
	}
	for j := i; j >= 0 && graph[j][1] == node; j-- {
		degree++
	}

	return degree, nil
}

// num convert nodes to integers.
func num(n [2]int, log2 uint) int64 {
	return int64(n[1]<<log2) + int64(n[0])
}

// interpolSearch search for index of node in graph by linear interpolation.
// Time complexity O(log log n).
func interpolSearch(graph [][2]int, node [2]int, log2 uint) int {

	l := 0
	r := len(graph) - 1
	n := num(node, log2)

	x := r
	for r != l {
		gl := num(graph[l], log2)
		gr := num(graph[r], log2)
		if n < gl || n > gr {
			return x
		}
		xf := float64(n-gl) / float64(gr-gl) * float64(r-l)
		x = l + int(xf)
		if graph[x] == node {
			break
		}
		if num(graph[x], log2) < n {
			x++
			l = x
			continue
		}
		x--
		r = x
	}

	return x
}
