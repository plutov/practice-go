package anagram

import (
	"strings"
)

func normalize(s string) string {
	return strings.Replace(strings.ToLower(s), " ", "", -1)
}

// FindAnagrams returns the anagrams of word from dictionary
func FindAnagrams(dictionary []string, word string) []string {
	var anagrams []string
	word = normalize(word)
	if len(word) == 0 {
		return nil
	}

	for _, w := range dictionary {
		ww := normalize(w)
		// ignore exact match or empty word(not anagram)
		if ww == word || len(ww) == 0 {
			continue
		}
		if len(strings.Trim(ww, word)) == 0 {
			anagrams = append(anagrams, w)
		}
	}
	return anagrams
}
