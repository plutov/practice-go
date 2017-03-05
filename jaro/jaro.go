package jaro

import (
	"strings"
)

func Distance(word1 string, word2 string) float64 {
	// ignore case
	word1, word2 = strings.ToLower(word1), strings.ToLower(word2)
	l1, l2 := len(word1), len(word2)

	// special cases
	switch {
	case l1 == 0 && l2 == 0: // both words are empty
		// return 0 ?
		return 1
	case l1 == 0 || l2 == 0: // either word is empty
		return 0
	case word1 == word2: // same string
		return 1
	}

	// check range
	w1, w2 := []rune(word1), []rune(word2)
	l1, l2 = len(w1), len(w2)
	r := l1
	if l2 > l1 {
		r = l2
	}
	if r /= 2; r > 0 {
		r -= 1
	}

	// count matching characters
	m := 0
	for i, c := range w1 {
		s, e := i-r, i+r+1
		if s >= l2 {
			break
		}
		if s < 0 {
			s = 0
		}
		if e > l2 {
			e = l2
		}
		for j := s; j < e; j++ {
			if w2[j] == c {
				m, w1[i], w2[j] = m+1, 0, 0
				break
			}
		}
	}

	// no matches
	if m == 0 {
		return 0
	}

	// transpositions
	t, rw1, rw2 := 0, []rune(word1), []rune(word2)
	for i, j, k, c := 0, 0, 0, 0; c < m; i++ {
		if w1[i] != 0 {
			continue
		}
		for j = k; w2[j] != 0; j++ {
		}
		if rw1[i] != rw2[j] {
			t++
		}
		k, c = j+1, c+1
	}

	// Jaro distance
	return float64((l1+l2)*m*m+l1*l2*(m-t/2)) / float64(3*l1*l2*m)
}
