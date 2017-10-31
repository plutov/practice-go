package coins

// Piles func
func Piles(n int) int {
	return piles(n, n)
}

type cache map[int]map[int]int

// m is our memoizing cache
var m = make(cache)

// piles implements the recursive partition function
// that memoizes values in a pkg local cache
//
// see: https://en.wikipedia.org/wiki/Partition_(number_theory)
// see: https://projecteuler.net/problem=78
func piles(n, k int) int {
	// base cases
	if n == 0 {
		return 1
	}
	if k == 0 || n < 0 {
		return 0
	}
	// ensure cache map
	if _, nok := m[n]; !nok {
		m[n] = make(map[int]int)
	}
	// check in cache
	if v := m[n][k]; v != 0 {
		return v
	}
	// compute recursively & return
	m[n][k] = piles(n-k, k) + piles(n, k-1)
	return m[n][k]
}
