package nodedegree

import (
	"fmt"
	"testing"
)

var smallGraph = [][]int{
	[]int{1, 2},
	[]int{1, 3},
}

var bigGraph = [][]int{
	[]int{1, 2},
	[]int{1, 3},
	[]int{2, 3},
	[]int{1, 4},
	[]int{3, 4},
	[]int{1, 5},
	[]int{2, 5},
	[]int{1, 6},
	[]int{2, 6},
	[]int{3, 6},
	[]int{3, 7},
	[]int{5, 7},
	[]int{6, 7},
	[]int{3, 8},
	[]int{4, 8},
	[]int{6, 8},
	[]int{7, 8},
	[]int{2, 9},
	[]int{5, 9},
	[]int{6, 9},
	[]int{2, 10},
	[]int{9, 10},
	[]int{6, 11},
	[]int{7, 11},
	[]int{8, 11},
	[]int{9, 11},
	[]int{10, 11},
	[]int{1, 12},
	[]int{6, 12},
	[]int{7, 12},
	[]int{8, 12},
	[]int{11, 12},
	[]int{6, 13},
	[]int{7, 13},
	[]int{9, 13},
	[]int{10, 13},
	[]int{11, 13},
	[]int{5, 14},
	[]int{8, 14},
	[]int{12, 14},
	[]int{13, 14},
	[]int{1, 15},
	[]int{2, 15},
	[]int{5, 15},
	[]int{9, 15},
	[]int{10, 15},
	[]int{11, 15},
	[]int{12, 15},
	[]int{13, 15},
	[]int{1, 16},
	[]int{2, 16},
	[]int{5, 16},
	[]int{6, 16},
	[]int{11, 16},
	[]int{12, 16},
	[]int{13, 16},
	[]int{14, 16},
	[]int{15, 16},
}

func TestDegree(t *testing.T) {
	var tests = []struct {
		nodes      int
		graph      [][]int
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
