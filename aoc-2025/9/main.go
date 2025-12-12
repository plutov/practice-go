package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

var xMap = make(map[int]int)
var yMap = make(map[int]int)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(file), "\n")

	points := make([]Point, len(lines))
	allX := make([]int, 0)
	allY := make([]int, 0)

	for i, line := range lines {
		if line == "" {
			continue
		}

		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		points[i] = Point{x, y}
		allX = append(allX, x)
		allY = append(allY, y)
	}

	sort.Ints(allX)
	sort.Ints(allY)
	uniqX := unique(allX)
	uniqY := unique(allY)

	for i, x := range uniqX {
		xMap[x] = i
	}
	for i, y := range uniqY {
		yMap[y] = i
	}

	grid := make([][]rune, len(uniqY))
	for i := range grid {
		grid[i] = make([]rune, len(uniqX))
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}

	zPoints := make([]Point, len(points))
	for i, p := range points {
		grid[yMap[p.y]][xMap[p.x]] = '#'
		zPoints[i] = Point{x: xMap[p.x], y: yMap[p.y]}
	}

	for i, a := range zPoints {
		b := zPoints[(i+1)%len(zPoints)]

		if a.x == b.x {
			y0, y1 := min(a.y, b.y), max(a.y, b.y)
			for y := y0; y <= y1; y++ {
				grid[y][a.x] = '#'
			}
		} else if a.y == b.y {
			x0, x1 := min(a.x, b.x), max(a.x, b.x)
			for x := x0; x <= x1; x++ {
				grid[a.y][x] = '#'
			}
		}
	}

	floodFill(grid, getInsidePoint(grid))

	maxArea := 0
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			a, b := points[i], points[j]
			if isEnclosed(a, b, grid) {
				area := (abs(a.x-b.x) + 1) * (abs(a.y-b.y) + 1)
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}

	fmt.Println(maxArea)
}

func unique(s []int) []int {
	if len(s) == 0 {
		return s
	}
	result := []int{s[0]}
	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1] {
			result = append(result, s[i])
		}
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func isEnclosed(a, b Point, grid [][]rune) bool {
	x1, x2 := min(xMap[a.x], xMap[b.x]), max(xMap[a.x], xMap[b.x])
	y1, y2 := min(yMap[a.y], yMap[b.y]), max(yMap[a.y], yMap[b.y])

	for x := x1; x <= x2; x++ {
		if grid[y1][x] == '.' || grid[y2][x] == '.' {
			return false
		}
	}

	for y := y1; y <= y2; y++ {
		if grid[y][x1] == '.' || grid[y][x2] == '.' {
			return false
		}
	}
	return true
}

func floodFill(grid [][]rune, start Point) {
	stack := []Point{start}
	dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for len(stack) > 0 {
		p := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if grid[p.y][p.x] != '.' {
			continue
		}
		grid[p.y][p.x] = 'X'

		for _, d := range dirs {
			nx, ny := p.x+d[0], p.y+d[1]
			if ny >= 0 && ny < len(grid) && nx >= 0 && nx < len(grid[0]) {
				if grid[ny][nx] == '.' {
					stack = append(stack, Point{x: nx, y: ny})
				}
			}
		}
	}
}

func getInsidePoint(grid [][]rune) Point {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] != '.' {
				continue
			}

			hitsLeft := 0
			prev := '.'

			for i := x; i >= 0; i-- {
				cur := grid[y][i]
				if cur != prev {
					hitsLeft++
				}
				prev = cur
			}

			if hitsLeft%2 == 1 {
				return Point{x: x, y: y}
			}
		}
	}
	panic("no inside point found")
}
