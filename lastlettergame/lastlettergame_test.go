package lastlettergame

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

var (
	// allwords will be filled in init() func
	allwords []string
	// dic will be inited and filled in init() func
	dic      map[string]bool
	expected = []string{"machamp", "petilil", "landorus", "scrafty", "yamask", "kricketune", "emboar", "registeel", "loudred", "darmanitan", "nosepass", "simisear", "relicanth", "heatmor", "rufflet", "trapinch", "haxorus", "seaking", "girafarig", "gabite", "exeggcute", "emolga", "audino"}
)

func init() {
	content, err := ioutil.ReadFile("pokemons.txt")
	if err != nil {
		panic(err)
	}

	allwords = strings.Split(string(content), "\n")
	if len(allwords) > 0 && allwords[len(allwords)-1] == "" {
		allwords = allwords[0 : len(allwords)-1]
	}

	dic = make(map[string]bool, len(allwords))
	for _, word := range allwords {
		dic[word] = true
	}
	fmt.Printf("%d words in dictionary\n", len(allwords))
}

func TestSequence(t *testing.T) {
	actual := Sequence(allwords)
	if !validateSequenceRes(actual) {
		t.Errorf("Sequence(%v) expected %v, got %v", allwords, expected, actual)
	}
}

func BenchmarkSequence(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sequence(allwords)
	}
}

func validateSequenceRes(result []string) bool {
	if len(result) != len(expected) {
		return false
	} else if len(result) < 1 {
		return true
	} else if _, ok := dic[result[0]]; !ok {
		return false
	}

	var unique map[string]bool = make(map[string]bool, len(expected))
	unique[result[0]] = true
	for i := 1; i < len(result); i++ {
		if _, ok := dic[result[i]]; !ok {
			return false
		} else if _, ok := unique[result[i]]; ok {
			return false
		}

		unique[result[i]] = true
		if !canBeNextWord(result[i-1], result[i]) {
			return false
		}
	}
	return true
}

func canBeNextWord(prev, next string) bool {
	if len(prev) < 2 || len(next) < 2 {
		return false
	}
	var rprev []rune = []rune(prev)
	var rnext []rune = []rune(next)
	return rprev[len(rprev)-1] == rnext[0]
}
