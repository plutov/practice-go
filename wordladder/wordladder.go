package wordladder

import "math"

// WordLadder is used to count the length of the shortest transformation sequence from first word to second word
func WordLadder(from string, to string, dic []string) int {
	if len(from) != len(to) {
		return 0
	}
	u := indexInSlice(from, dic)
	v := indexInSlice(to, dic)
	if v == -1 {
		return 0
	}
	if u == -1 {
		u = len(dic)
		dic = append(dic, from)
	}
	graph := generateGraph(dic)
	return dijkstra(graph, u, v)
}

func indexInSlice(s string, list []string) int {
	for i := 0; i < len(list); i++ {
		if s == list[i] {
			return i
		}
	}
	return -1
}

func generateGraph(dic []string) [][]int {
	g := [][]int{}
	for i := 0; i < len(dic); i++ {
		r := []int{}
		for j := 0; j < len(dic); j++ {
			t := distance(dic[i], dic[j])
			r = append(r, t)
		}
		g = append(g, r)
	}
	return g
}

func distance(s string, t string) int {
	if len(s) != len(t) {
		return 0
	}
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] != t[i] {
			count++
		}
	}
	if count == 1 {
		return 1
	}
	return 0
}

func dijkstra(graph [][]int, u int, v int) int {
	n := len(graph[0])
	visited := make([]bool, n)
	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = math.MaxInt64
	}
	dist[u] = 0
	for i := 0; i < n; i++ {
		t := minDistance(dist, visited)
		if t == -1 {
			return 0
		}
		if t == v {
			return dist[t] + 1
		}
		visited[t] = true
		for j := 0; j < n; j++ {
			if !visited[j] && graph[t][j] == 1 && dist[t] != math.MaxInt64 && dist[t]+graph[t][j] < dist[j] {
				dist[j] = dist[t] + graph[t][j]
			}
		}
	}
	return 0
}

func minDistance(dist []int, visited []bool) int {
	min := math.MaxInt64
	index := -1
	for i := 0; i < len(dist); i++ {
		if !visited[i] && dist[i] < min {
			min = dist[i]
			index = i
		}
	}
	return index
}
