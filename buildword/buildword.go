package buildword

import (
	"math"
	"strconv"
)

// Possible fragments of desired word,
// in usable form.
type Fragment struct {
	index  int // Where fragment fits in desired word
	length int // string length of fragment
}

func BuildWord(word string, fragments []string) int {

	var candidates []Fragment
	lw := len(word)
	contents := make(map[string]bool)

	for _, frag := range fragments {
		lf := len(frag)
		// Find all indexes into argument word where
		// this fragment could possibly fit. A given
		// fragment might match more than one index
		// into word.
		for i := 0; i+lf <= lw; i++ {
			matchingWordPart := word[i : i+lf]
			if matchingWordPart == frag {
				indexSuffix := strconv.Itoa(i)
				if !contents[frag+indexSuffix] {
					var n Fragment
					n.index = i
					n.length = lf
					candidates = append(candidates, n)
					contents[frag+indexSuffix] = true
				}
			}
		}
	}

	var min int = math.MaxInt32
	if len(candidates) > 0 {
		min = ChooseFragment(0, len(word), candidates, 0)
	}

	if min == math.MaxInt32 {
		return 0
	}
	return min
}

// from candidates[] find all entries that start at position, and
// don't end up longer than length. If there's any length left over,
// call recursively.
func ChooseFragment(position int, length int, candidates []Fragment, fragCount int) int {

	// Fragments that match substring of desired word at index position.
	var curr []Fragment

	for _, frag := range candidates {
		if frag.index == position {
			if position+frag.length == length {
				// This fragment matches substring of desired word,
				// and it completes the desired word.
				return fragCount + 1
			}
			curr = append(curr, frag)
		}
	}

	var min int = math.MaxInt32

	for _, frag := range curr {
		if position+frag.length <= length {
			m := ChooseFragment(position+frag.length, length, candidates, fragCount+1)
			if m < min {
				min = m
			}
		}
	}

	return min
}
