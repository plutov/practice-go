package anagram

import (
	"slices"
	"strings"
	"unicode"
)

func Normalize(s string) string {
	var r []rune
	for _, b := range s {
		b = unicode.ToLower(b)
		if b >= 'a' && b <= 'z' {
			r = append(r, b)
		}
	}
	slices.Sort(r)
	return string(r)
}

func FindAnagrams(dictionary []string, word string) (result []string) {
	n := Normalize(word)
	if len(n) == 0 {
		return []string{}
	}
	var res []string
	for _, s := range dictionary {
		n2 := Normalize(s)
		if n == n2 && !strings.EqualFold(s, word) {
			res = append(res, s)
		}
	}
	return res
}
