package main

import "fmt"

type point struct{ x, y int }

func Spiral(n int) {

	p := topLeft(n)
	minX := p.x
	for row := 0; row < n; row++ {

		for col := 0; col < n; col++ {
			val := spiralValue(p)
			fmt.Printf("%4d", val)
			p.x++
		}

		fmt.Println()
		p.x = minX
		p.y++
	}
}

// Value at given point in spiral.
// Origin {0, 0} is the spiral center with value 0.
func spiralValue(p point) int {

	// radius/winding number
	r := radius(p)
	if r == 0 {
		return 0
	}

	// nxn winding dimension
	n := 2*r + 1

	// value at top-right corner of current winding
	val := 4 * (r - 1) * r
	val += n - 1

	// right column
	// exclude last winding position at {r, r}
	if p != (point{r, r}) && le(p, point{r, -r}) {
		return val - r - p.y
	}

	// top row
	val += n - 1 // value at next corner ccw
	if le(p, point{-r, -r}) {
		return val - r - p.x
	}

	// left column
	val += n - 1
	if le(p, point{-r, r}) {
		return val - r + p.y
	}

	// bottom row
	val += n - 1
	return val - r + p.x
}

// lesser or equal winding position
func le(a, b point) bool {

	if a.x == b.x {
		if a.x > 0 {
			return a.y >= b.y
		}
		return a.y <= b.y
	}

	if a.y == b.y {
		if a.y > 0 {
			return a.x <= b.x
		}
		return a.x >= b.x
	}

	return false
}

// radius/winding number
func radius(p point) int {
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}
	return max(abs(p.x), abs(p.y))
}

// top-left coordinates of nxn spiral
func topLeft(n int) point {
	if n%2 == 1 {
		return point{-(n - 1) / 2, -(n - 1) / 2}
	}

	return point{-(n - 2) / 2, -n / 2}
}

func main() {
	Spiral(10)
}
