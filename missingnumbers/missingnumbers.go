package missingnumbers

// Missing takes slice of 1..n numbers exept two and finds missing ones.
func Missing(numbers []int) []int {
	out := []int{-1, -1}

	i := 0
	tempI := 0
	tempV := 0
	for {
		if i == len(numbers) {
			break
		}

		switch {
		case i+1 == numbers[i]:
			i += 1
			continue
		case numbers[i] > len(numbers):
			out[numbers[i]%(len(numbers)+1)] = i
			i += 1
			continue
		}

		tempI = numbers[i] - 1
		tempV = numbers[tempI]

		numbers[tempI] = numbers[i]
		numbers[i] = tempV
	}

	for oi := range out {
		if out[oi] == -1 {
			out[oi] = len(numbers) + oi + 1
			continue
		}
		out[oi] += 1
	}

	return out
}
