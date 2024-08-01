package url_shortener

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestShortenAndExpand(t *testing.T) {
	storage := &MemoryStorage{}
	shortener := NewURLShortener(storage)

	testURL := "https://example.com"
	short, err := shortener.Shorten(testURL)
	if err != nil {
		t.Fatalf("Failed to shorten URL: %v", err)
	}

	expanded, err := shortener.Expand(short)
	if err != nil {
		t.Fatalf("Failed to expand URL: %v", err)
	}

	if expanded != testURL {
		t.Errorf("Expected %s, got %s", testURL, expanded)
	}
}

func TestHandleShorten(t *testing.T) {
	storage := &MemoryStorage{}
	shortener := NewURLShortener(storage)

	form := url.Values{}
	form.Add("url", "https://example.com")
	req, err := http.NewRequest("POST", "/shorten", strings.NewReader(form.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(shortener.HandleShorten)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if !strings.Contains(rr.Body.String(), "Shortened URL:") {
		t.Errorf("Handler returned unexpected body: got %v", rr.Body.String())
	}
}

func TestHandleExpand(t *testing.T) {
	storage := &MemoryStorage{}
	shortener := NewURLShortener(storage)

	testURL := "https://example.com"
	short, _ := shortener.Shorten(testURL)

	req, err := http.NewRequest("GET", "/expand/"+short, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(shortener.HandleExpand)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusFound {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusFound)
	}

	location := rr.Header().Get("Location")
	if location != testURL {
		t.Errorf("Handler returned unexpected location: got %v want %v", location, testURL)
	}
}

func BenchmarkShorten(b *testing.B) {
	storage := &MemoryStorage{}
	shortener := NewURLShortener(storage)

	for i := 0; i < b.N; i++ {
		shortener.Shorten("http://example.com")
	}
}

func BenchmarkExpand(b *testing.B) {
	storage := &MemoryStorage{}
	shortener := NewURLShortener(storage)
	short, _ := shortener.Shorten("http://example.com")

	for i := 0; i < b.N; i++ {
		shortener.Expand(short)
	}
}
