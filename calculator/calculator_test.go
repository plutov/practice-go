package calculator

import (
	"fmt"
	"testing"
)

var tests = []struct {
	expr     string
	result   float64
	isErrNil bool
}{
	{"1 + 2", 3.0, true},
	{"1 + 10 * 10 + 1", 102.0, true},
	{"1 / 2 / 8", 0.0625, true},
	{"1 / ( 2 / 8 )", 4.0, true},
	{"1 - 2 - 3", -4.0, true},
	{"1 - ( 2 - 3 )", 2.0, true},
	{"(2 + 1) - (-3) * 4", 15.0, true},
	{"(", 0.0, false},
	{")", 0.0, false},
	{"(2 + 1) - (-3) * 4 +", 0.0, false},
	{"", 0.0, false},
	{"1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9 + 10", 55.0, true},
	{"() + ()", 0.0, false},
	{"(1 + 2) + (3 + 4)", 10.0, true},
	{"((5 - 2) * 3 + 1) / ((3 + 1) - 2)", 5.0, true},
}

func TestEval(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.expr, func(t *testing.T) {
			result, err := Eval(tt.expr)
			if result != tt.result {
				t.Errorf("Eval(%s) got %f, want %f", tt.expr, result, tt.result)
			}
			if (err == nil) != tt.isErrNil {
				t.Errorf("Eval(%s) got error %v, want error %v", tt.expr, err, tt.isErrNil)
			}
		})
	}
}

func BenchmarkEval(b *testing.B) {
	for b.Loop() {
		_, err := Eval("(2 + (1)) - (-3) * 4")
		if err != nil {
			fmt.Printf("Eval failed: %v", err)
		}
	}
}
