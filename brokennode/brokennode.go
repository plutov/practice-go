package brokennode

import (
	"math/bits"
)

func FindBrokenNodes(brokenNodes int, reports []bool) string {

	result := make([]byte, len(reports))

	// start-conf: broken nodes come first
	var conf uint64 = 1<<uint8(brokenNodes) - 1
	var max uint64 = 1 << uint8(len(reports))

out:
	for {
		for !check(conf, reports) {
			if conf = nextPerm(conf); conf >= max {
				break out
			}
		}

		merge(result, conf)
		if conf = nextPerm(conf); conf >= max {
			break
		}
	}

	return string(result)
}

const (
	// B broken
	B uint64 = 1
	// W working
	W uint64 = 0
)

// merge into results.
func merge(result []byte, conf uint64) {

	chr := func(c uint64) byte {
		if c&1 == B {
			return 'B'
		}
		return 'W'
	}

	for i := range result {
		switch {
		case result[i] == 0:
			result[i] = chr(conf)
		case result[i] != chr(conf):
			result[i] = '?'
		}
		conf >>= 1
	}
}

// check consistency with reports.
func check(conf uint64, reports []bool) bool {

	// duplicate first node behind last
	c := conf | ((conf & 1) << uint8(len(reports)))

	for i := range reports {
		if c&1 == B {
			c >>= 1
			continue
		}

		want := B
		if reports[i] {
			want = W
		}

		c >>= 1
		if c&1 != want {
			return false
		}
	}

	return true
}

// nextPerm next node configuration.
// https://graphics.stanford.edu/~seander/bithacks.html#NextBitPermutation
func nextPerm(v uint64) uint64 {
	t := v | (v - 1)
	return (t + 1) | ((^t&-^t - 1) >> (uint8(bits.TrailingZeros64(v)) + 1))
}
