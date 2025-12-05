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

	res := 0
	input := strings.Split(string(file), "\n")
	state := make([][]rune, len(input))
	for i, line := range input {
		state[i] = []rune(line)
	}

	for {
		removed := 0

		rows := make([][]rune, len(state))
		for i, row := range state {
			rows[i] = []rune(row)
		}

		for x, row := range rows {
			for y, v := range row {
				if v != '@' {
					continue
				}

				tmp := 0
				dx := []int{-1, 0, 1, -1, 1, -1, 0, 1}
				dy := []int{-1, -1, -1, 0, 0, 1, 1, 1}

				for i := range len(dx) {
					nx := x + dx[i]
					ny := y + dy[i]

					if nx >= 0 && nx < len(rows) && ny >= 0 && ny < len(row) {
						if len(rows[nx]) > 0 && rows[nx][ny] == '@' {
							tmp++
						}
					}
				}

				if tmp < 4 {
					res++
					removed++
					state[x][y] = '.'
				}
			}
		}

		if removed == 0 {
			break
		}
	}

	fmt.Println(res)
}
