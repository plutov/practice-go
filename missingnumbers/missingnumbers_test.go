package missingnumbers

import (
	"reflect"
	"testing"
)

type testCase struct {
	input    []int
	expected []int
}

var tests = []testCase{
	testCase{[]int{4, 2, 3}, []int{1, 5}},
	testCase{[]int{1, 2, 3, 4}, []int{5, 6}},
}

func init() {
	var bigSlice []int
	for i := 1; i <= 1000; i++ {
		if i != 100 && i != 900 {
			bigSlice = append(bigSlice, i)
		}
	}

	tests = append(tests, testCase{bigSlice, []int{100, 900}})
}

func TestMissing(t *testing.T) {
	for _, test := range tests {
		actual := Missing(test.input)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Fatalf("Missing(%v) expected %v, got %v", test.input, test.expected, actual)
		}
	}
}

func BenchmarkMissing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			Missing(test.input)
		}
	}
}
