package mergesort

func MergeSort(input []int) []int {
	if len(input) <= 1 {
		return input
	}

	return merge(MergeSort(input[:len(input)/2]), MergeSort(input[len(input)/2:]))
}

func merge(first []int, second []int) []int {
	a := make([]int, len(first)+len(second))
	i, j, k := 0, 0, 0
	for i < len(first) && j < len(second) {
		if first[i] < second[j] {
			a[k] = first[i]
			i++
		} else {
			a[k] = second[j]
			j++
		}
		k++
	}

	for i < len(first) {
		a[k] = first[i]
		i++
		k++
	}

	for j < len(second) {
		a[k] = second[j]
		j++
		k++
	}

	return a
}
