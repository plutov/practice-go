package wordladder

import "testing"

var tests = []struct {
	from     string
	to       string
	dic      []string
	expected int
}{
	{"from", "to", []string{}, 0},
	{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}, 5},
	{"hot", "dog", []string{"hot", "dog", "cog", "pot", "dot"}, 3},
	{"a", "b", []string{"c", "b"}, 2},
	{"lost", "cost", []string{"most", "fist", "lost", "cost", "fish"}, 2},
	{"talk", "tail", []string{"talk", "tons", "fall", "tail", "gale", "hall", "negs"}, 0},
	{"hot", "dog", []string{"hot", "dog"}, 0},
}

func TestWordLadder(t *testing.T) {
	for _, test := range tests {
		actual := WordLadder(test.from, test.to, test.dic)
		if actual != test.expected {
			t.Errorf("WordLadder(%s, %s, %v) expected %d, got %d", test.from, test.to, test.dic, test.expected, actual)
		}
	}
}

func BenchmarkWordLadder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			WordLadder(test.from, test.to, test.dic)
		}
	}
}
