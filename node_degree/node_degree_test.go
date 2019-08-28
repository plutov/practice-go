package nodedegree

import (
	"fmt"
	"testing"
)

var smallGraph = [][2]int{
	{1, 2},
	{1, 3},
}

var bigGraph = [][2]int{
	{1, 2},
	{1, 3},
	{2, 3},
	{1, 4},
	{3, 4},
	{1, 5},
	{2, 5},
	{1, 6},
	{2, 6},
	{3, 6},
	{3, 7},
	{5, 7},
	{6, 7},
	{3, 8},
	{4, 8},
	{6, 8},
	{7, 8},
	{2, 9},
	{5, 9},
	{6, 9},
	{2, 10},
	{9, 10},
	{6, 11},
	{7, 11},
	{8, 11},
	{9, 11},
	{10, 11},
	{1, 12},
	{6, 12},
	{7, 12},
	{8, 12},
	{11, 12},
	{6, 13},
	{7, 13},
	{9, 13},
	{10, 13},
	{11, 13},
	{5, 14},
	{8, 14},
	{12, 14},
	{13, 14},
	{1, 15},
	{2, 15},
	{5, 15},
	{9, 15},
	{10, 15},
	{11, 15},
	{12, 15},
	{13, 15},
	{1, 16},
	{2, 16},
	{5, 16},
	{6, 16},
	{11, 16},
	{12, 16},
	{13, 16},
	{14, 16},
	{15, 16},
}

func TestDegree(t *testing.T) {
	var tests = []struct {
		nodes      int
		graph      [][2]int
		node       int
		wantDegree int
		wantErrMsg string
	}{
		{3, smallGraph, 1, 2, ""},
		{3, smallGraph, 2, 1, ""},
		{3, smallGraph, 3, 1, ""},
		{3, smallGraph, 4, 0, "node 4 not found in the graph"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.nodes), func(t *testing.T) {
			got, err := Degree(tt.nodes, tt.graph, tt.node)
			if got != tt.wantDegree {
				t.Fatalf("got %d, expecting %d", got, tt.wantDegree)
			}
			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}
			if errMsg != tt.wantErrMsg {
				t.Fatalf("got '%s' error message, expecting '%s' error message", errMsg, tt.wantErrMsg)
			}
		})
	}
}

func BenchmarkDegree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Degree(16, bigGraph, 6)
	}
}
