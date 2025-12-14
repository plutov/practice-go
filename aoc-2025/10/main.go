package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type ButtonCombination struct {
	Joltages []int
	Presses  int
}

type Machine struct {
	Lights   []int
	Buttons  [][]int
	Patterns map[string][]ButtonCombination
	Joltages []int
}

func parseInput(input string) []Machine {
	var machines []Machine

	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		lightSpec := parts[0]
		joltageSpec := parts[len(parts)-1]
		buttonSpecs := parts[1 : len(parts)-1]

		// Parse lights
		lights := make([]int, 0)
		for _, c := range lightSpec[1 : len(lightSpec)-1] {
			if c == '#' {
				lights = append(lights, 1)
			} else {
				lights = append(lights, 0)
			}
		}

		// Parse buttons
		buttons := make([][]int, 0)
		for _, buttonSpec := range buttonSpecs {
			buttonStr := buttonSpec[1 : len(buttonSpec)-1]
			buttonParts := strings.Split(buttonStr, ",")
			button := make([]int, 0)
			for _, part := range buttonParts {
				n, _ := strconv.Atoi(part)
				button = append(button, n)
			}
			buttons = append(buttons, button)
		}

		// Build patterns
		patterns := make(map[string][]ButtonCombination)
		for n := 0; n < (1 << len(buttons)); n++ {
			lightResult := make([]int, len(lights))
			joltageMultiplier := make([]int, len(lights))
			presses := 0

			for buttonIndex := 0; buttonIndex < len(buttons); buttonIndex++ {
				if (n & (1 << buttonIndex)) != 0 {
					btn := buttons[buttonIndex]
					for _, light := range btn {
						lightResult[light] ^= 1
						joltageMultiplier[light]++
					}
					presses++
				}
			}

			key := tupleToString(lightResult)
			patterns[key] = append(patterns[key], ButtonCombination{
				Joltages: joltageMultiplier,
				Presses:  presses,
			})
		}

		// Parse joltages
		joltageStr := joltageSpec[1 : len(joltageSpec)-1]
		joltageParts := strings.Split(joltageStr, ",")
		joltages := make([]int, 0)
		for _, part := range joltageParts {
			n, _ := strconv.Atoi(part)
			joltages = append(joltages, n)
		}

		machines = append(machines, Machine{
			Lights:   lights,
			Buttons:  buttons,
			Patterns: patterns,
			Joltages: joltages,
		})
	}

	return machines
}

func tupleToString(arr []int) string {
	parts := make([]string, len(arr))
	for i, v := range arr {
		parts[i] = strconv.Itoa(v)
	}
	return strings.Join(parts, ",")
}

func part1(machines []Machine) int {
	total := 0
	for _, machine := range machines {
		minPresses := int(^uint(0) >> 1) // Max int
		key := tupleToString(machine.Lights)
		if combinations, ok := machine.Patterns[key]; ok {
			for _, c := range combinations {
				if c.Presses < minPresses {
					minPresses = c.Presses
				}
			}
			total += minPresses
		}
	}
	return total
}

func subtractAndHalf(jolt1, jolt2 []int) []int {
	result := make([]int, len(jolt1))
	for i := range jolt1 {
		result[i] = (jolt1[i] - jolt2[i]) / 2
	}
	return result
}

func isValidJoltage(joltages []int) bool {
	for _, j := range joltages {
		if j < 0 {
			return false
		}
	}
	return true
}

func solve(joltages []int, patterns map[string][]ButtonCombination, target []int) *int {
	// Check if we've reached the target
	if slicesEqual(joltages, target) {
		zero := 0
		return &zero
	}

	// Pull out the least significant bit from the joltages
	lsb := make([]int, len(joltages))
	for i, j := range joltages {
		lsb[i] = j & 1
	}

	lsbKey := tupleToString(lsb)
	if combinations, ok := patterns[lsbKey]; ok {
		var presses []int
		for _, c := range combinations {
			// Remove joltages produced by this button
			newJoltage := subtractAndHalf(joltages, c.Joltages)
			if isValidJoltage(newJoltage) {
				if rest := solve(newJoltage, patterns, target); rest != nil {
					presses = append(presses, c.Presses+2*(*rest))
				}
			}
		}
		if len(presses) > 0 {
			minVal := presses[0]
			for _, p := range presses {
				if p < minVal {
					minVal = p
				}
			}
			return &minVal
		}
	}

	return nil
}

func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func part2(machines []Machine) int {
	total := 0
	for _, m := range machines {
		target := make([]int, len(m.Joltages))
		if solution := solve(m.Joltages, m.Patterns, target); solution != nil {
			total += *solution
		}
	}
	return total
}

func main() {
	data, _ := os.ReadFile("input.txt")
	input := string(data)
	machines := parseInput(input)

	fmt.Printf("Part 1: %d\n", part1(machines))

	start := time.Now()
	result := part2(machines)
	elapsed := time.Since(start).Seconds()
	fmt.Printf("Part 2: %d (%.2fs)\n", result, elapsed)
}
