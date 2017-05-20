package functionfrequency

import (
	"io/ioutil"
	"reflect"
	"testing"
)

var expected = []string{"fmt.Println", "Honda", "append"}

func TestFunctionFrequency(t *testing.T) {
	code := getGoCode()
	actual := FunctionFrequency(code)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("FunctionFrequency() expected %v, got %v", expected, actual)
	}
}

func BenchmarkFunctionFrequency(b *testing.B) {
	code := getGoCode()
	for i := 0; i < b.N; i++ {
		FunctionFrequency(code)
	}
}

func getGoCode() []byte {
	code, err := ioutil.ReadFile("gocode.txt")
	if err != nil {
		panic(err)
	}

	return code
}
