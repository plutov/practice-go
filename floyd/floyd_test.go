package floyd

import (
	"testing"
)

var tests = []struct {
	rowsCount int
	expected  [][]int
}{
	{0, [][]int{}},
	{1, [][]int{[]int{1}}},
	{2, [][]int{[]int{1}, []int{2, 3}}},
	{3, [][]int{[]int{1}, []int{2, 3}, []int{4, 5, 6}}},
	{4, [][]int{[]int{1}, []int{2, 3}, []int{4, 5, 6}, []int{7, 8, 9, 10}}},
	{5, [][]int{[]int{1}, []int{2, 3}, []int{4, 5, 6}, []int{7, 8, 9, 10}, []int{11, 12, 13, 14, 15}}},
	{6, [][]int{[]int{1}, []int{2, 3}, []int{4, 5, 6}, []int{7, 8, 9, 10}, []int{11, 12, 13, 14, 15}, []int{16, 17, 18, 19, 20, 21}}},
}

func TestFloydTriangle(t *testing.T) {
	for _, test := range tests {
		actual := Triangle(test.rowsCount)
		if len(actual) != len(test.expected) {
			t.Fatalf("FloydTriangle(%d) expected length %d, got %d", test.rowsCount, len(test.expected), len(actual))
		}
		for k, v := range test.expected {
			if len(actual[k]) != len(v) {
				t.Fatalf("FloydTriangle(%d) expected length %d for row %d, got %d", test.rowsCount, len(v), k, len(actual[k]))
			}
			for k2, v2 := range v {
				if actual[k][k2] != v2 {
					t.Fatalf("FloydTriangle(%d) expected %d for row %d and column %d, got %d", test.rowsCount, v2, k, k2, actual[k][k2])
				}
			}
		}
	}
}

func BenchmarkTestFloydTriangle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			Triangle(test.rowsCount)
		}
	}
}
