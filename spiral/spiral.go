package main

import (
	"fmt"
	"strconv"
)

/*
returns the element at position x,y in the spiral of size n
y: line number between 0 and n-1
x: column number between 0 and n-1
*/
func element(n, x, y int) int {
	/*
	y: number of line [0 .. n-1]
	   The recursive structure of an even spiral:

	   +-----------------------------+
	   |          top line           |
	   +-------------------------+---+
	   |                         | r |
	   |                         | i |
	   |   odd subspiral         | g |
	   |     * size n-1          | h |
	   |     * line/col numbers  | t |
	   |       * y_sub = y - 1   |   |
	   |       * x_sub = x       | c |
	   |                         | o |
	   |                         | l |
	   +-------------------------+---+

	   The recursive structure of an odd spiral:

	   +---+-------------------------+
	   |   |                         |
	   | l |                         |
	   | e |   even subspiral        |
	   | f |     * size n-1          |
	   | t |     * line/col numbers  |
	   |   |       * y_sub = y       |
	   | c |       * x_sub = x - 1   |
	   | o |                         |
	   | l |                         |
	   +---+-------------------------+
	   |        bottom line          |
	   +-----------------------------+
	*/
	//
	// Due to recursion we only have to bother with numbers along the edges.
	//
	sqr := n * n
	// even
	if n%2 == 0 {
		if y == 0 { // top line: counted down with x from sqr - 1
			return sqr - 1 - x 
		}
		if x == n-1 { // right column: counted down with y from sqr - n
			return sqr - n - y
		}
		return element(n-1, x, y-1) // recursion to odd subspiral
	}
	// odd
	if y == n-1 { // bottom line: counted up with x from sqr - n
		return sqr - n + x
	}
	if x == 0 { // left column: counted up with y from sqr - n - (n - 1)
		return sqr - n - (n - 1) + y
	}
	return element(n-1, x-1, y) // recursion to even subspiral
}


func Spiral(n int) {
	str := strconv.Itoa(n * n - 1)
	length := len(str)
	for y := 0; y < n; y++ {
		for x := 0; x < n; x++ {
			fmt.Printf("%[2]*[1]v ", element(n, x, y), length)
		}
		fmt.Println()
	}
}

func main() {
	Spiral(10)
}
