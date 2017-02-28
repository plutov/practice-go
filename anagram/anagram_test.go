package anagram

import "testing"

var dictionary = []string{
	"about",
	"above",
	"angel",
	"angle",
	"batten",
	"battery",
	"battle",
	"galen",
	"evil",
	"lager",
	"large",
	"le batt",
}

var tests = []struct {
	word     string
	expected []string
}{
	{"", []string{}},
	{" ", []string{}},
	{"angel", []string{"angle", "galen"}},
	{"evil", []string{}},
	{"levi", []string{"evil"}},
	{"le batt", []string{}},
}

func TestFindAnagrams(t *testing.T) {
	for _, test := range tests {
		actual := FindAnagrams(dictionary, test.word)
		if len(actual) != len(test.expected) {
			t.Fatalf("FindAnagrams(%s) expected length %d, got %d", test.word, len(test.expected), len(actual))
		}

		for k, v := range test.expected {
			if actual[k] != v {
				t.Fatalf("FindAnagrams(%s) expected %v, got %v", test.word, test.expected, actual)
			}
		}
	}
}

func BenchmarkTestFloydTriangle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			FindAnagrams(dictionary, test.word)
		}
	}
}
