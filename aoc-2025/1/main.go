package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	curr := 50
	res := 0
	lines := strings.Split(string(file), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}

		diff, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		if strings.HasPrefix(line, "L") {
			new := curr - diff
			if new <= 0 {
				res += int(math.Abs(float64(new / 100)))
				if curr != 0 {
					res++
				}
			}

			curr = new % 100
			if curr < 0 {
				curr += 100
			}
		} else {
			new := curr + diff
			if new >= 100 {
				res += (new / 100)
			}
			curr = new % 100
		}

		fmt.Printf("Current: %d, Line: %s, Diff: %d, Res: %d\n", curr, line, diff, res)
	}

	fmt.Println(res)
}
