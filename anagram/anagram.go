package anagram

import (
	"sort"
	"strings"
	"unicode"
)

type MyRune []rune

func (r MyRune) Len() int           { return len(r) }
func (r MyRune) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r MyRune) Less(i, j int) bool { return r[i] < r[j] }

func Normalize(s string) string {
	var r MyRune
	for _, b := range s {
		b = unicode.ToLower(b)
		if b >= 'a' && b <= 'z' {
			r = append(r, b)
		}
	}
	sort.Sort(r)
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
