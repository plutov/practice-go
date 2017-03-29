package buildword

import (
	// "fmt"
	"math"
	"strconv"
)

// BuildWord("buildword", []string{"buil", "dwor", "bu", "uild", "wo", "rd"}) = 4

type Fragment struct {
	frag   string
	index  int
	length int
}

func BuildWord(word string, fragments []string) int {
	var candidates []Fragment
	lw := len(word)
	contents := make(map[string]bool)
	for _, frag := range fragments {
		lf := len(frag)
		for i := 0; i+lf <= lw; i++ {
			f := word[i : i+lf]
			if f == frag && !contents[frag+strconv.Itoa(i)] {
				var n Fragment
				n.frag = frag
				n.index = i
				n.length = lf
				candidates = append(candidates, n)
				contents[frag+strconv.Itoa(i)] = true
			}
		}
	}

/*
	for _, cand := range candidates {
		fmt.Printf("%q at %d, length %d\n", cand.frag, cand.index, cand.length)
	}
*/

	var min int = math.MaxInt32
	if len(candidates) > 0 {
		min = ChooseFragment("", 0, len(word), candidates, 0)
	}

	if min == math.MaxInt32 { return 0 }
	return min
}

func ChooseFragment(sofar string, position int, length int, candidates []Fragment, fragCount int) int {
	// fmt.Printf("Enter ChooseFragment(%q, %d, %d, %d)\n", sofar, position, length, fragCount)
	// from candidates[] find all entries that start at position, and
	// don't end up longer than length. If there's any length left over,
	// call recursively.
	var curr []Fragment

	for _, frag := range candidates {
		if frag.index == position && position+frag.length <= length {
			if position+frag.length == length {
				fragCount++
				// fmt.Printf("-> Returning %d on fragment %q, %q\n", fragCount, frag, sofar)
				return fragCount
			}
			curr = append(curr, frag)
		}
	}
	// fmt.Printf("\tChooseFragment(%q, %d, %d) from among %q\n", sofar, position, length, curr)

	var min int = math.MaxInt32

	if len(curr) > 0 {
		for _, frag := range curr {
			if position+frag.length <= length {
				m := ChooseFragment(sofar+frag.frag, position+frag.length, length, candidates, fragCount+1)

				if m < min {
					min = m
				}
			}
		}
	}

	return min
}
