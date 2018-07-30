package warriors

import (
	"strings"
)

// Count calculates warriors on the image
func Count(image string) int {
	count := 1
	parsed := parse(image)

	for j, row := range parsed {
		for i, col := range row {
			switch {
			case col > 1: //We've already seen part of this Byte Warrior
				paint(col, j, i, parsed)
			case col == 1: //This is a new Byte Warrior
				count++
				paint(count, j, i, parsed)
			}
		}
	}

	return count - 1

}

func parse(image string) [][]int {
	rows := strings.Split(image, "\n")
	parsed := make([][]int, len(rows)+1)
	for row, line := range rows {
		parsed[row] = make([]int, len(line))
		for col, r := range line {
			parsed[row][col] = int(r - '0')
		}
	}

	return parsed
}

func paint(color, pointJ, pointI int, image [][]int) {
	/*
		Need to paint unvisited pixels
		- - -
		- X X
		X X X
	*/
	for j := 0; j <= 1; j++ {
		for i := -j; i <= 1; i++ {
			toPaintJ := pointJ + j
			toPaintI := pointI + i

			inBounds := toPaintI >= 0 && toPaintJ < len(image) && toPaintI < len(image[toPaintJ])

			if inBounds && image[toPaintJ][toPaintI] == 1 {
				image[toPaintJ][toPaintI] = color
			}
		}
	}
}
