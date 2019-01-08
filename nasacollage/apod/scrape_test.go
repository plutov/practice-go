package apod_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/shogg/practice-go/nasacollage/apod"
)

func TestScrapeImageURLs(t *testing.T) {

	// skip if travis
	if _, ok := os.LookupEnv("TRAVIS"); ok {
		t.Skip("travis build")
	}

	var links []string
	err := apod.ScrapeImageURLs(
		"https://apod.nasa.gov/apod/archivepix.html",
		func(link string) {
			links = append(links, link)
			fmt.Println(link)
		})
	if err != nil {
		t.Fatal(err)
	}

	if len(links) == 0 {
		t.Error("no result")
	}
}
