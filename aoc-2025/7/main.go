package main

import (
	"fmt"
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
	for i, line := range lines {
		input[i] = make([]string, len(line))
		for j := 0; j < len(line); j++ {
			if line[j] == 'S' {
				input[i][j] = "|"
			} else {
				input[i][j] = string(line[j])
			}
		}
	}

	allInputs := allPossibleInputs(input)
	fmt.Println("Total inputs:", len(allInputs))

	uniqueTimelines := make(map[string]bool)
	for _, input := range allInputs {
		for i, line := range input {
			for j, char := range line {
				if i == 0 {
					continue
				}

				if char == "<" && j > 0 && input[i-1][j] == "|" {
					input[i][j-1] = "|"
					input[i][j] = "."
				} else if char == ">" && j < len(line)-1 && input[i-1][j] == "|" {
					input[i][j+1] = "|"
					input[i][j] = "."
					j++
				} else if char == "." && input[i-1][j] == "|" {
					input[i][j] = "|"
				} else if char != "|" {
					input[i][j] = "."
				}
			}
		}
		timelineStr := ""
		for _, line := range input {
			timelineStr += strings.Join(line, "")
		}
		uniqueTimelines[timelineStr] = true
	}

	fmt.Println(len(uniqueTimelines))
}

func allPossibleInputs(input [][]string) [][][]string {
	for i, line := range input {
		for j, char := range line {
			if char == "^" {
				leftInput := copyInput(input)
				leftInput[i][j] = "<"
				left := allPossibleInputs(leftInput)

				rightInput := copyInput(input)
				rightInput[i][j] = ">"
				right := allPossibleInputs(rightInput)

				return append(left, right...)
			}
		}
	}

	return [][][]string{input}
}

func copyInput(input [][]string) [][]string {
	newInput := make([][]string, len(input))
	for i := range input {
		newInput[i] = make([]string, len(input[i]))
		copy(newInput[i], input[i])
	}
	return newInput
}
