package secretmessage

import "sort"

// Decode func
func Decode(encoded string) string {
	// count
	counts := make(map[rune]int)
	for _, r := range encoded {
		counts[r]++
	}
	// only get runes >= counts[_]
	var runes []rune
	for k, v := range counts {
		if k != '_' && v >= counts['_'] {
			runes = append(runes, k)
		}
	}
	// sort runes by count
	sort.Slice(runes, func(i, j int) bool {
		return counts[runes[i]] > counts[runes[j]]
	})
	// return string
	return string(runes)
}
