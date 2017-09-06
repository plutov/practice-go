package functionfrequency

import (
	"io/ioutil"
	"reflect"
	"testing"
)

var expected = map[string][][]string{
	"gocode0.txt": [][]string{{"fmt.Println", "Honda", "append"}},
	"gocode1.txt": [][]string{{"cmd.CommandPath", "errors.Errorf", "len"}, {"cmd.CommandPath", "len", "errors.Errorf"}},
	"gocode2.txt": [][]string{{"fmt.Errorf", "os.FileMode", "strings.Replace"}, {"fmt.Errorf", "os.FileMode", "bindataRead"}, {"fmt.Errorf", "os.FileMode", "time.Unix"}},
	"gocode3.txt": [][]string{{"recursiveInterpolate", "errors.Errorf", "template.Substitute"}, {"recursiveInterpolate", "errors.Errorf", "errors.Wrapf"}, {"recursiveInterpolate", "errors.Errorf", "interpolateSectionItem"}, {"recursiveInterpolate", "errors.Errorf", "len"}, {"recursiveInterpolate", "errors.Errorf", "make"}},
}

func TestFunctionFrequency(t *testing.T) {
	for fileName, functionsCollection := range expected {
		actual := FunctionFrequency(getGoCode(fileName))
		found := false

		for _, functions := range functionsCollection {
			if reflect.DeepEqual(actual, functions) {
				found = true
				break
			}
		}

		if !found {
			t.Errorf("FunctionFrequency() expected one of %v, got %v. File: %s", functionsCollection, actual, fileName)
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
