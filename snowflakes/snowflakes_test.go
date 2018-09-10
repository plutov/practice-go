package snowflakes

import "testing"

var tests = []struct {
	name string
	n    int
	m    int
	want int
}{
	{"1_1", 1, 1, 1},
	{"3_1", 3, 1, 30},
	{"3_3", 3, 3, 6},
	{"11_1", 11, 1, 3027630},
	{"11_3", 11, 3, 19862070},
}

func TestOverlaidTriangles(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := OverlaidTriangles(tt.n, tt.m)
			if got != tt.want {
				t.Errorf("OverlaidTriangles(%d, %d) got %d, want %d", tt.n, tt.m, got, tt.want)
			}
		})
	}
}

func BenchmarkOverlaidTriangles(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OverlaidTriangles(11, 3)
	}
}
