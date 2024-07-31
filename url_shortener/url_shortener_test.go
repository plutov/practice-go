package url_shortener

import (
	"testing"
)

func TestShortenAndExpand(t *testing.T) {
	url := "https://example.com"
	short := Shorten(url)
	expanded := Expand(short)

	if expanded != url {
		t.Errorf("Expected %s, got %s", url, expanded)
	}
}

func BenchmarkShorten(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Shorten("http://example.com")
	}
}

func BenchmarkExpand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Expand("shortURL")
	}
}
