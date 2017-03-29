package buildword

import "testing"

var tests = []struct {
	word      string
	fragments []string
	expected  int
}{
	{"buildword", []string{"buil", "dwor", "bu", "ild", "wo", "rd"}, 4},
	{"answer", []string{"wer", "ans"}, 2},
	{"aaaaaa", []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa"}, 1},
	{"veeerrryy ttrrricckkkyy tteessssttt", []string{"tt", "t", "ee", "e", "v", "rr", "rrr", "kkk", "cc", "ssss", "y", "yy", " ", "ii", "i"}, 18},
	{"golang", []string{"g", "lang", "golan"}, 2},
	{"golang", []string{"g", "lang", "gola"}, 0},
}

func TestBuildWord(t *testing.T) {
	for _, test := range tests {
		actual := BuildWord(test.word, test.fragments)
		if actual != test.expected {
			t.Errorf("BuildWord(%s, %v) expected %d, got %d", test.word, test.fragments, test.expected, actual)
		}
	}
}

func BenchmarkBuildWord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			BuildWord(test.word, test.fragments)
		}
	}
}
