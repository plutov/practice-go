package secretmessage

import (
	"sort"
	"strings"
)

// pair tracks frequency of a character
type pair struct {
	c string
	n int
}

// Decode sorts encoded by descending rune frequency, discards all but one of each rune, then returns everything before "_"
func Decode(encoded string) string {
	var (
		chars = strings.Split(encoded, "")
		m     = make(map[string]int)
		s     = make([]pair, 1)
		out   = ""
		i     = -1
	)

	// count frequencies:
	for _, c := range chars {
		m[c]++
	}

	// Order the runes by descending frequency:
	for k, v := range m {
		s = append(s, pair{c: k, n: v})
	}
	sort.Slice(s, func(i, j int) bool {
		return s[i].n > s[j].n
	})

	// Append each character to `out` until we reach _:
	for {
		i++
		if s[i].c == "_" {
			break
		}
		out += s[i].c
	}

	return out
}
