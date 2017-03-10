package countsort

func CountSort(input []int, max int) []int {
	dic := make([]int, max+1)
	res := make([]int, len(input))
	for _, v := range input {
		dic[v]++
	}

	j := 0
	for k, c := range dic {
		for i := 0; i < c; i++ {
			res[j] = k
			j++
		}
	}

	return res
}
