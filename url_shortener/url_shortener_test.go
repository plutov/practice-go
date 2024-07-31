package url_shortener

import (
	"testing"
)

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
