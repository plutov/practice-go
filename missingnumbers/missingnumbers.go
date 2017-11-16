package missingnumbers

import (
	"sort"
)

// Missing func
func Missing(numbers []int) []int {
	res := make([]int, 2)

	// sort in ascending order to enable checking by successor calculated from index
	sort.Ints(numbers)

	offset := 0
	// array is sorted in natural order, should be starting with '1' so the value of every
	// entry in the array should be its index incremented by 1 altered by the already identified
	// missing numbers
	for idx, val := range numbers {
		expected := idx + offset + 1
		if expected != val {
			res[offset] = expected
			offset++
		}
	}

	// edge cases: missing numbers at the end
	if res[0] == 0 && res[1] == 0 {
		res[0] = len(numbers) + 1
		res[1] = len(numbers) + 2
	}

	if res[1] == 0 {
		res[1] = len(numbers) + 2
	}

	return res
}
