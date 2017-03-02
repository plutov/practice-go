package anagram

import (
	"fmt"
	"io/ioutil"
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
	{"Eleven plus two", []string{"Twelve plus one"}},
	{"Clint Eastwood", []string{"Old West Action"}},
	{"Protectionism", []string{"Cite no imports", "Nice to imports"}},
	{"Protection", []string{}},
	{"Funeral", []string{"Real fun"}},
	{"The Doors", []string{}},
	{"funeral", []string{"Real fun"}},
}

func init() {
	content, err := ioutil.ReadFile("dictionary.txt")
	if err != nil {
		panic(err)
	}

	dictionary = strings.Split(string(content), "\n")
	fmt.Printf("%d words in dictionary\n", len(dictionary))
}

func TestFindAnagrams(t *testing.T) {
	for _, test := range tests {
		actual := FindAnagrams(dictionary, test.word)
		if len(actual) != len(test.expected) {
			t.Fatalf("FindAnagrams(%s) expected length %d, got %d: %v", test.word, len(test.expected), len(actual), actual)
		}

		for k, v := range test.expected {
			if actual[k] != v {
				t.Fatalf("FindAnagrams(%s) expected %v, got %v", test.word, test.expected, actual)
			}
		}
	}
}

func BenchmarkTestAnagrams(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			FindAnagrams(dictionary, test.word)
		}
	}
}
