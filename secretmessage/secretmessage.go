package secretmessage

import "sort"

// Decode func
func Decode(encoded string) string {
	// count: size 27 for main case of alpha+_
	counts := make(map[rune]int, 27)
	for _, r := range encoded {
		counts[r]++
	}
	// only get runes with count >= counts[_]
	// cap 26 for base case of alpha chars
	runes := make([]rune, 0, 26)
	for k, v := range counts {
		if k != '_' && v >= counts['_'] {
			runes = append(runes, k)
		}
	}
	// sort runes by count desc
	sort.Slice(runes, func(i, j int) bool {
		return counts[runes[i]] > counts[runes[j]]
	})
	// return string
	return string(runes)
}
