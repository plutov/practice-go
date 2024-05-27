package collage_test

import (
	"fmt"
	"testing"

	"github.com/plutov/practice-go/nasacollage/collage"
)

func BenchmarkVariations(b *testing.B) {
	calls := 0
	for i := 0; i < b.N; i++ {
		collage.Variations(1000, 3, func(tuple []int) {
			calls++
		})
		fmt.Println(calls)
	}
}
