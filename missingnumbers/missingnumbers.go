package missingnumbers

// Missing func
func Missing(numbers []int) []int {
	numbers = append(numbers, numbers[0], numbers[1])
	for _, val := range numbers {
		if val < 0 {
			val = -val
		}
		if numbers[val-1] > 0 {
			numbers[val-1] = -numbers[val-1]
		}
	}

	n1 := 0
	for i, val := range numbers {
		if val > 0 {
			if n1 == 0 {
				n1 = i + 1
			} else {
				return []int{n1, i + 1}
			}
		}
	}
	return []int{}
}
