package url_shortener

import (
	"testing"
)

func TestRun(t *testing.T) {
	url := "https://example.com"
	short := Shorten(url)
	expanded := Expand(short)

	if expanded != url {
		t.Errorf("Expected %s, got %s", url, expanded)
	}
}

func BenchmarkRun(b *testing.B) {
	url := "https://example.com"
	for i := 0; i < b.N; i++ {
		short := Shorten(url)
		Expand(short)
	}
}
