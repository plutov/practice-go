package romannumerals

import "testing"

var encodeTests = []struct {
	roman   string
	integer int
	ok      bool
}{
	{"", -1, false},
	{"", 0, false},
	{"I", 1, true},
	{"XVIII", 18, true},
	{"XXXVIII", 38, true},
	{"CX", 110, true},
	{"CXXXVIII", 138, true},
	{"CCCLXVI", 366, true},
	{"DCXXXIV", 634, true},
	{"MMDCCXIV", 2714, true},
	{"MMMCD", 3400, true},
}

var decodeTests = []struct {
	roman   string
	integer int
	ok      bool
}{
	{"I", 1, true},
	{"", 0, false},
	{"1", 0, false},
	{"T", 0, false},
	{"X X X", 0, false},
	{"X0", 0, false},
	{"XVIII", 18, true},
	{"XXXVIII", 38, true},
	{"CX", 110, true},
	{"CXXXVIII", 138, true},
	{"CCCLXVI", 366, true},
	{"DCXXXIV", 634, true},
	{"MMDCCXIV", 2714, true},
	{"MMMCD", 3400, true},
}

func TestEncode(t *testing.T) {
	for _, test := range encodeTests {
		roman, ok := Encode(test.integer)

		if roman != test.roman || ok != test.ok {
			t.Errorf("Encode(%d) expected (\"%s\", %t), got (\"%s\", %t)", test.integer, test.roman, test.ok, roman, ok)
		}
	}
}

func TestDecode(t *testing.T) {
	for _, test := range decodeTests {
		integer, ok := Decode(test.roman)

		if integer != test.integer || ok != test.ok {
			t.Errorf("Decode(\"%s\") expected (%d, %t), got (%d, %t)", test.roman, test.integer, test.ok, integer, ok)
		}
	}
}

func BenchmarkEncode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range encodeTests {
			Encode(test.integer)
		}
	}
}

func BenchmarkDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range decodeTests {
			Decode(test.roman)
		}
	}
}
