package main

import (
	"os"
	"strings"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(file), "\n")
	lines = lines[:len(lines)-1]

	input := make([][]string, len(lines))
	indexOfRay := -1
	for i, line := range lines {
		input[i] = make([]string, len(line))
		for j := 0; j < len(line); j++ {
			if line[j] == 'S' {
				input[i][j] = "|"
				indexOfRay = j
			} else {
				input[i][j] = string(line[j])
			}
		}
	}

	rayCounts := make(map[int]int)
	rayCounts[indexOfRay] = 1

	for i, line := range input {
		if i == 0 {
			continue
		}

		for j, char := range line {
			if char == "^" && rayCounts[j] > 0 {
				if j > 0 {
					rayCounts[j-1] += rayCounts[j]
				}
				if j < len(line)-1 {
					rayCounts[j+1] += rayCounts[j]
				}
				rayCounts[j] = 0
			} else {
				continue
			}
		}

	}

	res := 0
	for _, count := range rayCounts {
		res += count
	}
}
