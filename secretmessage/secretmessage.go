package secretmessage

import "sort"

//RuneCounter counts the frequency of runes using a lookup table
type RuneCounter struct {
	lookuptable map[rune]int
}

//NewRuneCounter create new rune counter
func NewRuneCounter() RuneCounter {
	return RuneCounter{
		make(map[rune]int),
	}
}

//Count counts the frequecy of a rune in a stirng
func (counter *RuneCounter) Count(str string) {
	for _, r := range str {
		counter.lookuptable[r] = counter.lookuptable[r] + 1
	}
}

//CountedRune represnts the rune counts
type CountedRune struct {
	Rune  rune
	Count int
}

//Runes gets the CountedRunes
func (counter RuneCounter) Runes() []CountedRune {
	runes := make([]CountedRune, len(counter.lookuptable))
	i := 0
	for r, c := range counter.lookuptable {
		runes[i] = CountedRune{r, c}
		i++
	}

	return runes
}

//CountedRunes implements sort.Interface to sort counted runes by count
type CountedRunes []CountedRune

func (a CountedRunes) Len() int           { return len(a) }
func (a CountedRunes) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a CountedRunes) Less(i, j int) bool { return a[i].Count < a[j].Count }

// Decode func
func Decode(encoded string) string {
	counter := NewRuneCounter()
	counter.Count(encoded)
	runes := counter.Runes()

	sort.Sort(sort.Reverse(CountedRunes(runes)))

	var decoded []rune
	for _, r := range runes {
		if r.Rune == '_' {
			break
		}
		decoded = append(decoded, r.Rune)
	}
	return string(decoded)
}
