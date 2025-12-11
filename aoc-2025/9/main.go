package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(file), "\n")
	lines = lines[:len(lines)-1]

	maxArea := 0
	pairs := make([][2]int, len(lines))
	board := make([][]int, 100000)
	for i, line := range lines {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])

		pairs[i] = [2]int{x, y}
		if board[x] == nil {
			board[x] = make([]int, 100000)
		}
		board[x][y] = 1
	}

	for _, p := range pairs {
		for _, p1 := range pairs {
			if p1[0] <= p[0] && p1[1] <= p[1] {
				area := (p[0] - p1[0] + 1) * (p[1] - p1[1] + 1)
				if area > maxArea {
					maxArea = area
				}
			}

			if p1[0] >= p[0] && p1[1] <= p[1] {
				area := (p1[0] - p[0] + 1) * (p[1] - p1[1] + 1)
				if area > maxArea {
					maxArea = area
				}
			}

			if p1[0] <= p[0] && p1[1] >= p[1] {
				area := (p[0] - p1[0] + 1) * (p1[1] - p[1] + 1)
				if area > maxArea {
					maxArea = area
				}
			}

			if p1[0] >= p[0] && p1[1] >= p[1] {
				area := (p1[0] - p[0] + 1) * (p1[1] - p[1] + 1)
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}

	fmt.Println(maxArea)
}
