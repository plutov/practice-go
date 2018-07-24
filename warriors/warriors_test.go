package warriors

import "testing"

type testCase struct {
	input  string
	output int
}

var tests = []testCase{
	testCase{"100100\n001010\n000000\n110000\n111000\n010100", 3},
	testCase{"111\n111\n111", 1},
	testCase{"101\n101\n101", 2},
}

func TestRun(t *testing.T) {
	for _, test := range tests {
		actual := Count(test.input)
		if actual != test.output {
			t.Fatalf("Count(%s) expected %d, got %d", test.input, test.output, actual)
		}
	}
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(tests[0].input)
	}
}
