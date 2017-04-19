package webscraping

import "testing"

var tests = []struct {
	timezone string
	expected string
}{
	{"", ""},
	{"DT", ""},
	{"Alaska", ""},
	{"UTC", "Apr. 19, 12:59:44 UTC"},
	{"HAST", "Apr. 19, 02:59:44 AM HAST"},
	{"Apr. 19, 12:59:44 UTC", ""},
	{"US Naval Observatory", ""},
	{"AM", ""},
}

func TestGetTime(t *testing.T) {
	for _, test := range tests {
		actual := GetTime(test.timezone)

		if test.expected != actual {
			t.Fatalf("GetTime(\"%s\") expected %s, got %s", test.timezone, test.expected, actual)
		}
	}
}

func BenchmarkGetTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			GetTime(test.timezone)
		}
	}
}
