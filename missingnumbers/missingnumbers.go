package missingnumbers

import (
	"math"
)

// Missing func
func Missing(numbers []int) []int {

	// (could be parallelized)
	sum, sumSq := 0, 0
	for _, n := range numbers {
		sum += n
		sumSq += n * n
	}

	N := len(numbers) + 2
	gauss := N * (N + 1) / 2
	gaussSq := N * (N + 1) * (2*N + 1) / 6
	diff := gauss - sum
	diffSq := gaussSq - sumSq

	// missing numbers x, y:
	// x + y = diff
	// x^2 + y^2 = diffSq
	// x^2 + (diff - x)(diff - x) = diffSq
	// x^2 + (x^2 - 2*diff*x + diff^2) = diffSq
	// 2*x^2 - 2*diff*x + diff^2 - diffSq = 0
	a, b, c := 2, -2*diff, diff*diff-diffSq

	// x1, x2 = (-b +- sqrt(b^2 - 4ac)) / 2a
	sqrt := int(math.Sqrt(float64(b*b - 4*a*c)))
	x1 := (-b - sqrt) / (2 * a)
	x2 := (-b + sqrt) / (2 * a)

	return []int{x1, x2}
}
