package lastlettergame

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
	"testing"
)

var (
	// dic will be filled in init() func
	dic      []string
	expected = []string{"machamp", "petilil", "landorus", "scrafty", "yamask", "kricketune", "emboar", "registeel", "loudred", "darmanitan", "nosepass", "simisear", "relicanth", "heatmor", "rufflet", "trapinch", "haxorus", "seaking", "girafarig", "gabite", "exeggcute", "emolga", "audino"}
)

func init() {
	content, err := ioutil.ReadFile("pokemons.txt")
	if err != nil {
		panic(err)
	}

	dic = strings.Split(string(content), "\n")
	if len(dic) > 0 && dic[len(dic)-1] == "" {
		dic = dic[0 : len(dic)-1]
	}
	fmt.Printf("%d words in dictionary\n", len(dic))
}

func TestSequence(t *testing.T) {
	actual := Sequence(dic)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Sequence(%v) expected %v, got %v", dic, expected, actual)
	}
}

func BenchmarkSequence(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sequence(dic)
	}
}
