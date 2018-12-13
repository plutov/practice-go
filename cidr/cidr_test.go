package cidr

import (
	"fmt"
	"testing"
)

var tests = []struct {
	ips  []string
	want int
}{
	{[]string{"0000:0000:0000:0000:0000:0000:0000:0000"}, 128},
	{[]string{"0000:0000:0000:0000:0000:0000:0000:0000", "0000:0000:0000:0000:0000:0000:0000:0001"}, 127},
	{[]string{"0123:4567:89ab:cdef:fedc:ba98:7654:3210", "0246:8ace:1357:9bdf:0246:8ace:1357:9bdf"}, 248},
	{[]string{"0000:0000:0000:0000:0000:0000:0000:ffff", "0000:0000:0000:0000:0000:0000:0000:0000"}, 142},
	{[]string{"ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff", "0000:0000:0000:0000:0000:0000:0000:0000"}, 254},
}

func TestNumberOfMasks(t *testing.T) {
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test%d", i), func(t *testing.T) {
			got := NumberOfMasks(tt.ips)
			if got != tt.want {
				t.Errorf("NumberOfMasks(%v) got %d, want %d", tt.ips, got, tt.want)
			}
		})
	}
}

func BenchmarkNumberOfMasks(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NumberOfMasks(tests[4].ips)
	}
}
