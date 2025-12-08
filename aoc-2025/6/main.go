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
	lines := strings.Split(string(file), "\n")
	lines = lines[:len(lines)-1]
	numbers := []string{}
	operation := ""
	for col := len(lines[0]) - 1; col >= 0; col-- {
		allEmpty := true
		number := ""
		for row := 0; row < len(lines); row++ {
			val := " "
			if col < len(lines[row]) {
				val = string(lines[row][col])
			}

			if row == len(lines)-1 {
				if val != " " {
					operation = val
				}
			} else {
				if val != " " {
					allEmpty = false
					number += val
				}
			}
		}
		if number != "" {
			numbers = append(numbers, number)
		}

		if allEmpty || col == 0 {
			numbersInt := make([]int, len(numbers))
			for i, num := range numbers {
				if num != "" {
					numbersInt[i], _ = strconv.Atoi(num)
				} else {
					numbersInt[i] = 0
				}
			}
			res += doMath(numbersInt, operation)
			fmt.Println("col:", col, "numbers:", numbers, "operation:", operation, "result:", res)

			numbers = []string{}
			operation = ""
		}
	}

	fmt.Println(res)
}

func doMath(numbers []int, operation string) int {
	switch operation {
	case "+":
		sum := 0
		for _, num := range numbers {
			sum += num
		}
		return sum
	case "*":
		product := 1
		for _, num := range numbers {
			product *= num
		}
		return product
	}

	return 0
}
