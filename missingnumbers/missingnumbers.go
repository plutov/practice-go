package missingnumbers

// Missing func
func Missing(numbers []int) []int {

	len := len(numbers)
	var miss, sum, max, outside int
	for i := range numbers {
		for j := numbers[i] - 1; j != i; j = numbers[i] - 1 {

			if numbers[i] == 0 {
				break
			}

			if j >= len {
				numbers[i] = 0
				outside++
				miss = i + 1
				sum += j + 1
				if j+1 > max {
					max = j + 1
				}
				break
			}

			if j+1 == miss {
				miss = i + 1
			}

			if j < i {
				sum += j + 1
			}

			numbers[i], numbers[j] = numbers[j], numbers[i]
		}

		if numbers[i] != 0 {
			sum += i + 1
		}
	}

	switch outside {
	case 0:
		return []int{len + 1, len + 2}
	case 1:
		if max == len+1 {
			return []int{miss, len + 2}
		}
		return []int{miss, len + 1}
	default:
		gauss := (max*max + max) / 2
		miss2 := gauss - sum - miss
		swapAsc(&miss, &miss2)
		return []int{miss, miss2}
	}
}

func swapAsc(a, b *int) {
	if *b < *a {
		*a, *b = *b, *a
	}
}
