package countsort

import (
	"reflect"
	"testing"
)

var test = struct {
	input    []int
	expected []int
	max      int
}{}

// Add input with 10000 elements
func init() {
	test.max = 10000

	for i := 10000; i > 0; i-- {
		test.input = append(test.input, i)
		test.expected = append(test.expected, 10001-i)
	}
}

func TestCountSort(t *testing.T) {
	actual := CountSort(test.input, test.max)
	if !reflect.DeepEqual(actual, test.expected) {
		t.Errorf("CountSort(%v, %d) expected %v, got %v", test.input, test.max, test.expected, actual)
	}
}

func BenchmarkCountSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountSort(test.input, test.max)
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bubble(test.input)
	}
}

func bubble(input []int) {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(input)-1; i++ {
			if input[i+1] < input[i] {
				swap(input, i, i+1)
				swapped = true
			}
		}
	}
}

func swap(input []int, i, j int) {
	tmp := input[j]
	input[j] = input[i]
	input[i] = tmp
}
