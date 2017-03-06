package mergesort

import (
	"reflect"
	"testing"
)

var tests = []struct {
	input    []int
	expected []int
}{
	{[]int{}, []int{}},
	{[]int{0}, []int{0}},
	{[]int{-100}, []int{-100}},
	{[]int{-100, -200}, []int{-200, -100}},
	{[]int{1, 0, -1}, []int{-1, 0, 1}},
	{[]int{12, 11, 13, 5, 6, 7}, []int{5, 6, 7, 11, 12, 13}},
	{[]int{3, 4, 2, 1, 7, 5, 8, 9, 0, 6}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
	{[]int{100, 200, 100, 200, 100, 200, 100, 200}, []int{100, 100, 100, 100, 200, 200, 200, 200}},
	{[]int{2, 3, 14, 10, 8, 1, 12, 9}, []int{1, 2, 3, 8, 9, 10, 12, 14}},
}

// Add input with 10000 elements
func init() {
	input := make([]int, 10000)
	expected := []int{}
	for i := 9999; i >= 0; i-- {
		if i%3 == 0 {
			expected = append(expected, -i)
		}
	}
	for i := 0; i < 10000; i++ {
		if i%3 == 0 {
			input[i] = -i
		} else {
			input[i] = i
			expected = append(expected, i)
		}
	}

	tests = append(tests, struct {
		input    []int
		expected []int
	}{input, expected})
}

func TestMergeSort(t *testing.T) {
	for _, test := range tests {
		actual := MergeSort(test.input)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("MergeSort(%v) expected %v, got %v", test.input, test.expected, actual)
		}
	}
}

func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			MergeSort(test.input)
		}
	}
}
