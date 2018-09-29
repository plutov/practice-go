package brokennode

func FindBrokenNodes(brokenNodes int, reports []bool) string {

	result := make([]byte, len(reports))

	var conf uint64 = 1<<uint64(brokenNodes) - 1
	var n uint64 = 1 << uint64(len(reports))

out:
	for {
		for !check(conf, reports) {
			if conf = nextPerm(conf); conf >= n {
				break out
			}
		}

		merge(result, conf)
		if conf = nextPerm(conf); conf >= n {
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

// merge a node configuration into results.
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

// check consistence of a node configuration with reports.
func check(conf uint64, reports []bool) bool {

	// copy first node behind last
	c := conf | (conf&1)<<uint64(len(reports))

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

// nextPerm computes a new node configuration.
// https://graphics.stanford.edu/~seander/bithacks.html#NextBitPermutation
func nextPerm(v uint64) uint64 {
	t := v | (v - 1) + 1
	return t | ((t&-t)/(v&-v)>>1 - 1)
}
