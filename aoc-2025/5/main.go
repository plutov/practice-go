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

	res := 0
	rangeLines := strings.Split(string(file), "\n")
	ranges := [][]int{}

	for _, line := range rangeLines {
		if line == "" {
			break
		}
		parts := strings.Split(line, "-")
		valMin, _ := strconv.Atoi(parts[0])
		valMax, _ := strconv.Atoi(parts[1])
		ranges = append(ranges, []int{valMin, valMax})
	}

	for {
		merged := false
		indexToRemove := -1

		for i, r1 := range ranges {
			for j, r2 := range ranges {
				if i == j {
					continue
				}

				if r2[0] >= r1[0] && r2[0] <= r1[1] {
					merged = true
					ranges[i][1] = max(r1[1], r2[1])
					indexToRemove = j
					break
				} else if r2[1] >= r1[0] && r2[1] <= r1[1] {
					merged = true
					ranges[i][0] = min(r1[0], r2[0])
					indexToRemove = j
					break
				}
			}

			if merged {
				break
			}
		}

		if !merged {
			break
		}

		if indexToRemove != -1 {
			ranges = append(ranges[:indexToRemove], ranges[indexToRemove+1:]...)
		}
	}

	for _, r := range ranges {
		res += r[1] - r[0] + 1
	}

	fmt.Println(res)
}
