package anagram

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

var dictionary []string

var tests = []struct {
	word     string
	expected []string
}{
	{"", []string{}},
	{" ", []string{}},
	{"angel", []string{"angle", "galen", "glean", "lange"}},
	{"evil", []string{"levi", "live", "veil", "vile"}},
	{"le batt", []string{}},
}

func init() {
	r, err := http.Get("http://www.puzzlers.org/pub/wordlists/unixdict.txt")
	if err != nil {
		panic(err)
		return
	}
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
		return
	}

	dictionary = strings.Split(string(content), "\n")
	fmt.Printf("%d words in dictionary\n", len(dictionary))
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
