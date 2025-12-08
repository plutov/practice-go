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

	sum := 0
	banks := strings.Split(string(file), "\n")
	for _, bank := range banks {
		if bank == "" {
			continue
		}
		bateries := parseBank(bank)
		rounds := 12
		nextIndex := 0
		joltages := []int{}

		for round := 12; round > 0; round-- {
			to := len(bateries) - round + 1

			max := 0
			maxIndex := 0
			for i, v := range bateries[nextIndex:to] {
				if v > max {
					max = v
					maxIndex = i
				}
			}
			nextIndex += maxIndex + 1

			joltages = append(joltages, max)
		}

		for i := 0; i < len(joltages); i++ {
			sum += joltages[i] * int(math.Pow10(rounds-i-1))
		}
	}

	fmt.Println(sum)
}

func parseBank(bank string) []int {
	res := make([]int, len(bank))
	for i, c := range bank {
		res[i], _ = strconv.Atoi(string(c))
	}
	return res
}
