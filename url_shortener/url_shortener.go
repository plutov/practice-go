package url_shortener

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"sync"
)

type Storage interface {
	Save(key, value string) error
	Load(key string) (string, error)
}

type MemoryStorage struct {
	data sync.Map
}

func (m *MemoryStorage) Save(key, value string) error {
	m.data.Store(key, value)
	return nil
}

func (m *MemoryStorage) Load(key string) (string, error) {
	if value, ok := m.data.Load(key); ok {
		return value.(string), nil
	}
	return "", errors.New("key not found")
}

type URLShortener struct {
	storage Storage
}

func NewURLShortener(storage Storage) *URLShortener {
	return &URLShortener{storage: storage}
}

func (us *URLShortener) Shorten(url string) (string, error) {
	hash := sha256.Sum256([]byte(url))
	encoded := base64.URLEncoding.EncodeToString(hash[:])
	short := encoded[:8]

	err := us.storage.Save(short, url)
	if err != nil {
		return "", fmt.Errorf("failed to save the URL: %v", err)
	}

	return short, nil
}

func (us *URLShortener) Expand(short string) (string, error) {
	url, err := us.storage.Load(short)
	if err != nil {
		return "", fmt.Errorf("failed to load the URL: %v", err)
	}
	return url, nil
}

func (us *URLShortener) HandleShorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	url := r.FormValue("url")
	if url == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	short, err := us.Shorten(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Shortened URL: %s\n", short)
}

func (us *URLShortener) HandleExpand(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	short := r.URL.Path[len("/expand/"):]
	if short == "" {
		http.Error(w, "Short code is required", http.StatusBadRequest)
		return
	}

	url, err := us.Expand(short)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	http.Redirect(w, r, url, http.StatusFound)
}
