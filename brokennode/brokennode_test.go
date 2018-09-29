package brokennode

import "testing"

var tests = []struct {
	brokenNodes int
	reports     []bool
	want        string
}{
	{1, []bool{true, false, false}, "WWB"},
	{2, []bool{true, false, false}, "B??"},
	{3, []bool{true, false, false}, "BBB"},
	{1, []bool{false, true}, "WB"},
	{1, []bool{true, true, true, false}, "BWWW"},
	{2, []bool{false, true, true, true}, "WBBW"},
	{3, []bool{false, true, false, false}, "?B??"},
	{1, []bool{true, true, false, true, true}, "WWWBW"},
	{2, []bool{true, false, true, true, false}, "B??WW"},
	{3, []bool{false, false, true, false, true, false}, "??????"},
	{4, []bool{true, false, true, true, false, false, true, false}, "????WB??"},
	{3, []bool{true, true, true, true, true, true, true, true, true, false, false, true, true, true, false, false, true, true, true, true, true, false, true, true, true, true}, "WWWWWWWWWWBWWWWBWWWWWWBWWW"},
}

func TestFindBrokenNodes(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			got := FindBrokenNodes(tt.brokenNodes, tt.reports)
			if got != tt.want {
				t.Errorf("FindBrokenNodes(%d, %v) got %s, want %s", tt.brokenNodes, tt.reports, got, tt.want)
			}
		})
	}
}

func BenchmarkFindBrokenNodes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindBrokenNodes(tests[0].brokenNodes, tests[0].reports)
	}
}
