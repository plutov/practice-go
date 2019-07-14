package anagram

import "strings"

func isSubset(s, S string) bool {
	used := make([]bool, len(S))
OuterLoop:
	for _, r := range s {
		if r == ' ' {
			continue
		}
		for i, R := range S {
			if !used[i] && r == R {
				used[i] = true
				continue OuterLoop
			}
		}
		return false
	}
	return true
}

func nonEmpty(s string) bool {
	for _, r := range s {
		if r != ' ' {
			return true
		}
	}
	return false
}

func isAnagram(s1, s2 string) bool {
	s1 = strings.ToLower(s1)
	s2 = strings.ToLower(s2)
	return s1 != s2 && nonEmpty(s1) && nonEmpty(s2) && isSubset(s1, s2) && isSubset(s2, s1)
}

// FindAnagrams returns all the anagrams of `word` in dictionary.
// Faster than a hashtable-based method for short words.
func FindAnagrams2(dictionary []string, word string) []string {
	// assumes that words are short
	result := make([]string, 0)
	for _, w := range dictionary {
		if isAnagram(word, w) {
			result = append(result, w)
		}
	}
	return result
}
