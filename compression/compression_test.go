package compression

import (
	"fmt"
	"os"
	"testing"
)

var (
	originalContent string
	originalLength  int64
)

func init() {
	code, _ := os.ReadFile("dataset.txt")
	originalContent = string(code)

	originalLength = int64(len(originalContent))
}

func TestEncodeDecode(t *testing.T) {
	encoded := Encode(originalContent)
	decoded := Decode(encoded)

	if originalContent != decoded {
		t.Errorf("Decode(Encode(originalContent)) = %s, want %s", decoded, originalContent)
	}

	// print compression ratio
	encodedLength := int64(len(encoded))
	fmt.Printf("\nOriginal length: %d, Encoded length: %d, Compression ratio: %.2f\n\n", originalLength, encodedLength, float64(encodedLength)/float64(originalLength))
}

func BenchmarkEncodeDecode(b *testing.B) {
	b.SetBytes(originalLength)
	for i := 0; i < b.N; i++ {
		encoded := Encode(originalContent)
		Decode(encoded)
	}
}
