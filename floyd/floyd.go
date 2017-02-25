package floyd

// Triangle makes a Floyd's triangle matrix with rows count.
func Triangle(rows int) [][]int {
	triangle := make([][]int, rows)
	var i int
	for r := 0; r < rows; r++ {
		triangle[r] = make([]int, r+1)
		for c := 0; c < r+1; c++ {
			i++
			triangle[r][c] = i
		}
	}
	return triangle
}
