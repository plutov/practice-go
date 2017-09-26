package coins

import "testing"

var tests = []struct {
	coins    int
	expected int
}{
	{1, 1},
	{2, 2},
	{3, 3},
	{4, 5},
	{5, 7},
	{6, 11},
	{7, 15},
	{10, 42},
	{15, 176},
	{22, 1002},
	{30, 5604},
	{100, 190569292},
}

func TestPiles(t *testing.T) {
	for _, test := range tests {
		actual := Piles(test.coins)

		if actual != test.expected {
			t.Errorf("Piles(%d) expected %d, got %d", test.coins, test.expected, actual)
		}
	}
}

func BenchmarkPiles(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Piles(100)
	}
}
