package snowflakes

// OverlaidTriangles retunrs number of triangles which are M levels deep
func OverlaidTriangles(n, m int) int {

	tri, t := make([]int, n+1), make([]int, n+1)
	tri[0] = 1

	for k := 0; k < n-1; k++ {

		for j := 0; j <= k; j++ {
			off := j & 0x7ffffffffffffffe
			j1 := j & 1
			for i := 0; i < 3; i++ {
				t[off+i] += tri[j] * triGen[j1][i]
			}
		}

		tri, t = t, tri
		for i := range t {
			t[i] = 0
		}
	}

	if m&1 == 0 {
		return 0
	}
	return tri[m-1]
}

var triGen = [2][4]int{
	{6, 6, 0},
	{-1, 1, 1},
}
