package reverseparentheses

import "testing"

var tests = []struct {
	origin   string
	expected string
}{
	{"(bar)", "rab"},
	{"foo(bar)baz", "foorabbaz"},
	{"foo(bar(baz))blim", "foobazrabblim"},
	{"(abc)d(efg)", "cbadgfe"},
	{"foobarbaz", "foobarbaz"},
	{"((bar))", "bar"},
	{"g(o)(((la)))(ng)", "goalgn"},
	{"foo()bar", "foobar"},
}

func TestReverse(t *testing.T) {
	for _, test := range tests {
		actual := Reverse(test.origin)

		if actual != test.expected {
			t.Errorf("Reverse(%s) expected \"%s\", got \"%s\"", test.origin, test.expected, actual)
		}
	}
}
