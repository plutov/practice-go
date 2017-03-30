package shorthash

import (
	"reflect"
	"sort"
	"testing"
)

var tests = []struct {
	dictionary string
	maxLen     int
	expected   []string
}{
	{"ab", 1, []string{"a", "b"}},
	{"ab", 2, []string{"a", "b", "aa", "bb", "ab", "ba"}},
	{"123", 3, []string{"1", "2", "3", "11", "12", "13", "22", "23", "33", "111", "112", "113", "122", "123", "133", "222", "223", "233", "333"}},
	{"ab", 3, []string{"a", "b", "aa", "bb", "ab", "ba", "aaa", "baa", "aba", "aab", "bbb", "abb", "bab", "bba"}},
	{"a", 4, []string{"a", "aa", "aaa", "aaaa"}},
	{"磨宿", 1, []string{"磨", "宿"}},
	{"", 1, []string{}},
	{"a", 0, []string{}},
}

func TestGenerateShortHashes(t *testing.T) {
	for _, test := range tests {
		actual := GenerateShortHashes(test.dictionary, test.maxLen)
		expected := test.expected

		sort.Strings(expected)
		sort.Strings(actual)

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("GenerateShortHashes(%s, %d) expected %v, got %v", test.dictionary, test.maxLen, expected, actual)
		}
	}
}

func BenchmarkGenerateShortHashes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			GenerateShortHashes(test.dictionary, test.maxLen)
		}
	}
}
