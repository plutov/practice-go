package anagram

import (
	"sort"
)

func FindAnagrams(dictionary []string, word string) []string {
	wordSlice := runeSlice(word)
	sort.Sort(wordSlice)
	var anagrams []string
	for _, w := range dictionary {
		if len(w) != len(word) || w == word {
			continue
		}
		s := runeSlice(w)
		sort.Sort(s)
		if !same(wordSlice, s) {
			continue
		}
		anagrams = append(anagrams, w)
	}
	return anagrams
}

func same(a, b runeSlice) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

type runeSlice []rune

func (s runeSlice) Len() int           { return len(s) }
func (s runeSlice) Less(i, j int) bool { return s[i] < s[j] }
func (s runeSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s runeSlice) String() string     { return string(s) }
