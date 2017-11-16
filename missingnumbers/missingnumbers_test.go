package missingnumbers

import (
	"math/rand"
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

func TestMissingFuzzy(t *testing.T) {

	expected := make([]int, 2)

	for i := 0; i < 10; i++ {

		n := 10000
		numbers := rand.Perm(n)
		for i := range numbers {
			numbers[i]++
		}

		x1 := rand.Intn(n)
		x2 := rand.Intn(n - 1)

		expected[0] = numbers[x1]
		numbers = append(numbers[:x1], numbers[x1+1:]...)
		expected[1] = numbers[x2]
		numbers = append(numbers[:x2], numbers[x2+1:]...)
		if expected[0] > expected[1] {
			expected[0], expected[1] = expected[1], expected[0]
		}

		actual := Missing(numbers)

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("test %d: expected %v was %v", i, expected, actual)
		}
	}
}

func BenchmarkMissingFuzzy_1e2(b *testing.B) {
	benchmarkMissingFuzzy(b, 1e2)
}

func BenchmarkMissingFuzzy_1e3(b *testing.B) {
	benchmarkMissingFuzzy(b, 1e3)
}

func BenchmarkMissingFuzzy_1e4(b *testing.B) {
	benchmarkMissingFuzzy(b, 1e4)
}

func BenchmarkMissingFuzzy_1e5(b *testing.B) {
	benchmarkMissingFuzzy(b, 1e5)
}

func BenchmarkMissingFuzzy_1e6(b *testing.B) {
	benchmarkMissingFuzzy(b, 1e6)
}

func benchmarkMissingFuzzy(b *testing.B, n int) {

	for i := 0; i < b.N; i++ {

		b.StopTimer()
		numbers := rand.Perm(n)
		for i := range numbers {
			numbers[i]++
		}
		x1 := rand.Intn(n)
		x2 := rand.Intn(n - 1)
		numbers = append(numbers[:x1], numbers[x1+1:]...)
		numbers = append(numbers[:x2], numbers[x2+1:]...)
		b.StartTimer()

		_ = Missing(numbers)
	}
}
