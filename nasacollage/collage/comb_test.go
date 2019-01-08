package collage_test

import (
	"fmt"
	"testing"

	"github.com/shogg/practice-go/nasacollage/collage"
)

func ExampleCombinations() {

	collage.Combinations(5, 3, func(tuple []int) {
		fmt.Println(tuple)
	})

	// Output:
	// [0 1 2]
	// [0 1 3]
	// [0 1 4]
	// [0 2 3]
	// [0 2 4]
	// [0 3 4]
	// [1 2 3]
	// [1 2 4]
	// [1 3 4]
	// [2 3 4]
}
func ExampleCombinations3Choose2() {

	collage.Combinations(3, 2, func(tuple []int) {
		fmt.Println(tuple)
	})

	// Output:
	// [0 1]
	// [0 2]
	// [1 2]
}

func ExamplePermutations() {

	collage.Permutations([]int{0, 2, 4}, func(tuple []int) {
		fmt.Println(tuple)
	})

	// Output:
	// [0 2 4]
	// [2 0 4]
	// [4 0 2]
	// [0 4 2]
	// [2 4 0]
	// [4 2 0]
}

func ExampleVariations() {

	collage.Variations(4, 3, func(tuple []int) {
		fmt.Println(tuple)
	})

	// Output:
	// [0 1 2]
	// [1 0 2]
	// [2 0 1]
	// [0 2 1]
	// [1 2 0]
	// [2 1 0]
	// [0 1 3]
	// [1 0 3]
	// [3 0 1]
	// [0 3 1]
	// [1 3 0]
	// [3 1 0]
	// [0 2 3]
	// [2 0 3]
	// [3 0 2]
	// [0 3 2]
	// [2 3 0]
	// [3 2 0]
	// [1 2 3]
	// [2 1 3]
	// [3 1 2]
	// [1 3 2]
	// [2 3 1]
	// [3 2 1]
}

func ExampleVariations3Choose2() {

	collage.Variations(3, 2, func(tuple []int) {
		fmt.Println(tuple)
	})

	// Output:
	// [0 1]
	// [1 0]
	// [0 2]
	// [2 0]
	// [1 2]
	// [2 1]
}

func BenchmarkVariations(b *testing.B) {

	calls := 0
	for i := 0; i < b.N; i++ {
		collage.Variations(1000, 3, func(tuple []int) {
			calls++
		})
		fmt.Println(calls)
	}
}
