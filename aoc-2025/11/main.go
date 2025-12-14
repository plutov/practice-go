package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Graph map[string][]string

var (
	G     Graph
	cache = make(map[string]int64)
)

func readGraph(filename string) Graph {
	g := make(Graph)
	file, _ := os.Open(filename)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ": ")
		if len(parts) < 2 {
			continue
		}

		node := parts[0]
		neighbors := strings.Fields(parts[1])
		g[node] = neighbors
	}

	return g
}

func dfs(node, dst string) int64 {
	cacheKey := node + ":" + dst
	if val, ok := cache[cacheKey]; ok {
		return val
	}

	var result int64
	if node == dst {
		result = 1
	} else {
		for _, neighbor := range G[node] {
			result += dfs(neighbor, dst)
		}
	}

	cache[cacheKey] = result
	return result
}

func part1() int64 {
	return dfs("you", "out")
}

func part2() int64 {
	sequences := [][]string{
		{"svr", "fft", "dac", "out"},
		{"svr", "dac", "fft", "out"},
	}

	var total int64 = 0
	for _, seq := range sequences {
		product := int64(1)
		for i := 0; i < len(seq)-1; i++ {
			product *= dfs(seq[i], seq[i+1])
		}
		total += product
	}
	return total
}

func main() {
	G = readGraph("input.txt")

	result1 := part1()
	fmt.Printf("Part 1: %d\n", result1)

	result2 := part2()
	fmt.Printf("Part 2: %d\n", result2)
}
