package template

import "testing"

func TestRun(t *testing.T) {}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Run()
	}
}
