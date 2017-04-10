package romannumerals

import "strings"

type romanArabic struct {
	roman  string
	arabic int
}

var numbers = []romanArabic{
	{"M", 1000},
	{"CM", 900},
	{"D", 500},
	{"CD", 400},
	{"C", 100},
	{"XC", 90},
	{"L", 50},
	{"XL", 40},
	{"X", 10},
	{"IX", 9},
	{"V", 5},
	{"IV", 4},
	{"I", 1},
}

func Encode(n int) (string, bool) {
	result := ""

	for _, ra := range numbers {
		for n >= ra.arabic {
			n -= ra.arabic
			result += ra.roman
		}
	}

	return result, len(result) > 0
}

func Decode(s string) (int, bool) {
	result := 0

	for _, ra := range numbers {
		for strings.HasPrefix(s, ra.roman) {
			s = s[len(ra.roman):]
			result += ra.arabic
		}
	}

	if len(s) != 0 {
		return 0, false
	}

	return result, result > 0
}
