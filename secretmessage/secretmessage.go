package secretmessage

import (
	"sort"
	"strings"
)

func Encode(e string) string {
	m := make(map[rune]int)
	for _, r := range e {
		count := m[r]
		count++
		m[r] = count
	}

	type RuneCount struct {
		R rune
		C int
	}

	var totals []RuneCount

	for k, v := range m {
		totals = append(totals, RuneCount{R: k, C: v})
	}

	sort.Slice(totals, func(i, j int) bool { return totals[i].C > totals[j].C })

	or := make([]rune, len(totals), len(totals))
	for i := range totals {
		or[i] = totals[i].R
	}

	s := string(or)
	if idx := strings.Index(s, "_"); idx != -1 {
		s = s[:idx]
	}

	return s
}
