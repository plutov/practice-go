package collage

// Variations call f with all variations of k out of n nummers (no repetions).
func Variations(n, k int, f func([]int)) {
	Combinations(n, k, func(tuple []int) {
		Permutations(tuple, f)
	})
}

// Combinations call f with all combinations of k out of n nummers (no repetions).
func Combinations(n, k int, f func([]int)) {

	tuple := make([]int, k)

	var comb func(int, int)
	comb = func(i, next int) {
		for j := next; j < n; j++ {
			tuple[i] = j
			if i == k-1 {
				f(tuple)
				continue
			}
			comb(i+1, j+1)
		}
	}

	comb(0, 0)
}

// Permutations call f with all permutations of numbers in a tuple.
func Permutations(tuple []int, f func([]int)) {

	tupleCopy := make([]int, len(tuple))
	copy(tupleCopy, tuple)

	heapsAlgorithm(tupleCopy, len(tuple), f)
}

func heapsAlgorithm(tuple []int, n int, f func([]int)) {

	if n == 1 {
		f(tuple)
		return
	}

	for i := 0; i < (n - 1); i++ {
		heapsAlgorithm(tuple, n-1, f)
		if n&1 == 0 {
			tuple[n-1], tuple[i] = tuple[i], tuple[n-1]
		} else {
			tuple[n-1], tuple[0] = tuple[0], tuple[n-1]
		}
	}
	heapsAlgorithm(tuple, n-1, f)
}

// NumVariations calc n!/(n-k)!
func NumVariations(n, k int) int {
	num := n
	for i := n - 1; i >= n-k+1; i-- {
		num *= i
	}
	return num
}
