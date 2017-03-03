package jaro

import "testing"

const Epsilon = 0.000001

var tests = []struct {
	word1    string
	word2    string
	expected float64
}{
	{"dixon", "dicksonx", 0.766667},
	{"dicksonx", "DIXON", 0.766667},
	{"Cosmo", "Kosmo", 0.866667},
	{"Cosmer Kramo", "Kosmoo Karme", 0.694444},
	{"MARTHA", "marhta", 0.944444},
	{"five", "ten", 0},
	{"jones", "johnson", 0.790476},
	{"impression", "impression", 1},
	{"jaro", "oraj", 0.5},
	{"", "", 1},
	{" ", " ", 1},
	{"c++", "++c", 0.555556},
}

func TestDistance(t *testing.T) {
	for _, test := range tests {
		actual := Distance(test.word1, test.word2)
		if abs(actual-test.expected) > Epsilon {
			t.Errorf("Distance(%s, %s) expected %f, got %f", test.word1, test.word2, test.expected, actual)
		}
	}
}

func BenchmarkDistance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			Distance(test.word1, test.word2)
		}
	}
}

func abs(f float64) float64 {
	if f < 0 {
		return -f
	}
	return f
}
