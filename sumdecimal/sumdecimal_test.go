package sumdecimal

import "testing"

var tests = []struct {
	n        int
	expected int
}{
	{2, 4482},
	{-1, 0},
	{1, 0},
	{100, 0},
	{0, 0},
	{4627, 4426},
	{1234, 4477},
}

func TestSumDecimal(t *testing.T) {
	for _, test := range tests {
		actual := SumDecimal(test.n)
		if actual != test.expected {
			t.Errorf("SumDecimal(%d) expected %d, got %d", test.n, test.expected, actual)
		}
	}
}

func BenchmarkSumDecimal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			SumDecimal(test.n)
		}
	}
}
