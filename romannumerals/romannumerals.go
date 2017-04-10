package romannumerals

func div(num int, denom int) (int, int) {
	return num / denom, num % denom
}

type symbolValue struct {
	Symbol string
	Value  int
}

var romanNumerals = []symbolValue{
	symbolValue{"M", 1000},
	symbolValue{"D", 500},
	symbolValue{"C", 100},
	symbolValue{"L", 50},
	symbolValue{"X", 10},
	symbolValue{"V", 5},
	symbolValue{"I", 1},
}

// Encode takes an integer as its parameter and
// returns a string containing the Roman numeral representation of that
// integer.
func Encode(number int) (string, bool) {

	if number <= 0 {
		return "", false
	}

	stringRep := ""

	for idx, symVal := range romanNumerals {
		q, r := div(number, symVal.Value)

		if symVal.Value != 1000 && q == 4 || q == 9 {
			stringRep = stringRep + symVal.Symbol + romanNumerals[idx-1].Symbol
		} else {
			for q > 0 {
				stringRep = stringRep + symVal.Symbol
				q--
			}
		}

		number = r

		if number == 0 {
			break
		}
	}

	return stringRep, true
}

// Decode takes a Roman numeral as its argument and
//returns its value as a numeric decimal integer.
func Decode(stringRep string) (int, bool) {

	if stringRep == "" {
		return 0, false
	}

	value := 0
	last := 1000

	for _, numeral := range stringRep {
		var curr int

		switch numeral {
		case 'M':
			curr = 1000
		case 'D':
			curr = 500
		case 'C':
			curr = 100
		case 'L':
			curr = 50
		case 'X':
			curr = 10
		case 'V':
			curr = 5
		case 'I':
			curr = 1
		default:
			return 0, false
		}

		if curr > last {
			value -= last
			value += (curr - last)
		} else {
			value += curr
		}

		last = curr
	}

	return value, true
}
