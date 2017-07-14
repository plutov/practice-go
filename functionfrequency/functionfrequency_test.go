package functionfrequency

import (
	"io/ioutil"
	"reflect"
	"testing"
)

var expected = map[string][]string{
	"gocode0.txt": []string{"fmt.Println", "Honda", "append"},
	"gocode1.txt": []string{"cmd.CommandPath", "errors.Errorf", "len"},
	"gocode2.txt": []string{"fmt.Errorf", "os.FileMode", "strings.Replace"},
	"gocode3.txt": []string{"recursiveInterpolate", "errors.Errorf", "template.Substitute"},
}

func TestFunctionFrequency(t *testing.T) {
	for fileName, functions := range expected {
		actual := FunctionFrequency(getGoCode(fileName))
		if !reflect.DeepEqual(actual, functions) {
			t.Errorf("FunctionFrequency() expected %v, got %v. File: %s", functions, actual, fileName)
		}
	}
}

func BenchmarkFunctionFrequency(b *testing.B) {
	for fileName := range expected {
	code := getGoCode(fileName)
		for i := 0; i < b.N; i++ {
			FunctionFrequency(code)
		}
	}
}

func getGoCode(fileName string) []byte {
	code, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	return code
}
