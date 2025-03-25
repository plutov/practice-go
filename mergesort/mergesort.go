package mergesort

// MergeSort is used to sort an array of integer
func MergeSort(input []int) []int {
	aux := make([]int, len(input))

	var sort func(input []int)
	sort = func(s []int) {
		if len(s) <= 1 {
			return
		}
		m := len(s) / 2
		sort(s[:m])
		sort(s[m:])

		for k, l, r := 0, 0, m; k < len(s); k++ {
			if r >= len(s) || (l < m && s[l] <= s[r]) {
				aux[k] = s[l]
				l++
			} else {
				aux[k] = s[r]
				r++
			}
		}
		copy(s, aux)
	}

	tmp := append([]int{}, input...)
	sort(tmp)
	return tmp
}
