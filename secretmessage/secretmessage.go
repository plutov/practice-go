package secretmessage

import "sort"

// Decode func
func Decode(encoded string) string {
	// count: size 27 for main case of alpha+_
	counts := make([]int, 27)
	for _, r := range encoded {
		if r == '_' {
			counts[26]++
		} else {
			counts[r-'a']++
		}
	}
	// only get runes with count >= counts[_]
	// cap 26 for base case of alpha chars
	runes := make([]rune, 0, 26)
	for k, v := range counts {
		if k != 26 && v >= counts[26] {
			runes = append(runes, rune('a'+k))
		}
	}
	// sort runes by count desc
	sort.Slice(runes, func(i, j int) bool {
		return counts[runes[i]-'a'] > counts[runes[j]-'a']
	})
	// return string
	return string(runes)
}
