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
			curr = (curr - diff) % 100
			if curr < 0 {
				curr += 100
			}
		} else {
			curr = (curr + diff) % 100
		}

		if curr == 0 {
			res++
		}
	}

	fmt.Println(res)
}
