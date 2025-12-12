package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type (
	State   uint64
	Machine struct {
		state  State
		pushes []State
		ampers []int
	}
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\n")
	machines := parseLines(lines)

	count1 := puzzle1(machines)
	fmt.Println("Puzzle I. Count [494]: ", count1)
}

func puzzle1(machines []Machine) int {
	count := 0
	for _, m := range machines {
		combs := bfsXOR(m.state, m.pushes)
		count += len(combs)
	}
	return count
}

type Sample struct {
	state State
	comb  []State
}

func onPush(state State, push State) State {
	return state ^ push
}

func bfsXOR(target State, numbers []State) []State {
	queue := []Sample{{state: 0, comb: []State{}}}

	best := make(map[State]int)
	best[0] = 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.state == target {
			return current.comb
		}

		for _, push := range numbers {
			newVal := onPush(current.state, push)
			newLen := len(current.comb) + 1

			// если ещё не было или есть короче путь
			if prevLen, ok := best[newVal]; !ok || newLen < prevLen {
				best[newVal] = newLen
				newComb := append([]State{}, current.comb...)
				newComb = append(newComb, push)
				queue = append(queue, Sample{state: newVal, comb: newComb})
			}
		}
	}
	return nil
}

func parseLines(lines []string) []Machine {
	machines := make([]Machine, len(lines))
	for j, line := range lines {
		if line == "" {
			continue
		}
		machines[j] = parseLine(line)
	}
	return machines
}

func parseLine(line string) Machine {
	parts := strings.Split(line, " ")
	machine := Machine{}
	count := 0
	machine.state, count = parseState(parts[0])
	machine.ampers = parseAmpers(parts[len(parts)-1])
	machine.pushes = parsePushes(parts[1:len(parts)-1], count)
	return machine
}

func parsePushes(str []string, length int) []State {
	var states []State
	for _, sv := range str {
		si := strings.TrimSpace(sv)
		si = strings.TrimPrefix(si, "(")
		si = strings.TrimSuffix(si, ")")
		ns := strings.Split(si, ",")
		bools := make([]bool, length)
		for _, n := range ns {
			parsed, _ := strconv.Atoi(n)
			bools[parsed] = true
		}
		states = append(states, BoolsToBitmask(bools))
	}
	return states
}

func parseAmpers(str string) []int {
	s := strings.TrimSpace(str)
	s = strings.TrimPrefix(s, "{")
	s = strings.TrimSuffix(s, "}")
	ns := strings.Split(s, ",")
	nums := make([]int, len(ns))
	for j, n := range ns {
		nums[j], _ = strconv.Atoi(n)
	}
	return nums
}

func parseState(str string) (State, int) {
	s := strings.TrimSpace(str)
	s = strings.TrimPrefix(s, "[")
	s = strings.TrimSuffix(s, "]")
	var bools []bool
	for _, b := range s {
		if b == '#' {
			bools = append(bools, true)
		} else {
			bools = append(bools, false)
		}
	}
	return BoolsToBitmask(bools), len(s)
}

func BoolsToBitmask(bools []bool) State {
	var mask State
	n := len(bools)
	for j, b := range bools {
		if j >= 64 {
			break
		}
		if b {
			mask |= 1 << uint(n-1-j)
		}
	}
	return State(mask)
}
