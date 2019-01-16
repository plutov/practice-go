package collage

import "math"

func Disjoint(a, b []int) bool {

	for i := range a {
		for j := range b {
			if a[i] == b[j] {
				return false
			}
		}
	}

	return true
}

func Min(n int, f func(int) int) int {

	min, mini := math.MaxInt16, math.MaxInt16
	for i := 0; i < n; i++ {
		m := f(i)
		if m < min {
			min = m
			mini = i
		}
	}

	return mini
}

func Max(n int, f func(int) int) int {

	max, maxi := 0, 0
	for i := 0; i < n; i++ {
		m := f(i)
		if m > max {
			max = m
			maxi = i
		}
	}

	return maxi
}
