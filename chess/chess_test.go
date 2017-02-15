package chess

import "testing"

var tests = []struct {
	white  string
	black  string
	attack bool
	error  bool
}{
	{"f1", "g2", false, false},
	{"d4", "f4", false, false},
	{"a8", "b6", true, false},
	{"b7", "d8", true, false},
	{"c6", "b8", true, false},
	{"d1", "f2", true, false},
	{"e3", "g2", true, false},
	{"f8", "h7", true, false},
	{"g1", "h3", true, false},
	{"h4", "g2", true, false},
	{"b4", "b4", false, true},
	{"a8", "b9", false, true},
	{"a0", "b1", false, true},
	{"g3", "i5", false, true},
	{"not", "valid", true, true},
	{"", "", false, true},
}

func TestCanKnightAttack(t *testing.T) {
	for _, test := range tests {
		switch attack, err := CanKnightAttack(test.white, test.black); {
		case err != nil:
			if !test.error {
				t.Fatalf("CanKnightAttack(%s, %s) returned error %q. Error not expected.",
					test.white, test.black, err)
			}
		case test.error:
			t.Fatalf("CanKnightAttack(%s, %s) = %t, %v. Expected error.",
				test.white, test.black, attack, err)
		case attack != test.attack:
			t.Fatalf("CanKnightAttack(%s, %s) = %t, want %t.",
				test.white, test.black, attack, test.attack)
		}
	}
}

func BenchmarkCanKnightAttack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			CanKnightAttack(test.white, test.black)
		}
	}
}
