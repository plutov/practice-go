package apod

import (
	"bufio"
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

var regexpPageURL = regexp.MustCompile(`(?i)(ap\d{6}.html)`)
var regexpImageURL = regexp.MustCompile(`(?i)src="(image/[^"]*)"`)

func ScrapeImageURLs(mainURL string, callback func(string)) error {

	lastSlash := strings.LastIndex(mainURL, "/")
	baseURL := mainURL[:lastSlash]

	pages, err := scrape(mainURL, regexpPageURL)
	if err != nil {
		return err
	}

	for _, page := range pages {

		pageURL := fmt.Sprintf("%s/%s", baseURL, page)
		images, err := scrape(pageURL, regexpImageURL)
		if err != nil {
			return fmt.Errorf("%s: %s", page, err)
		}

		for _, img := range images {
			imgURL := fmt.Sprintf("%s/%s", baseURL, img)
			callback(imgURL)
		}

	}

	return nil
}

func scrape(url string, re *regexp.Regexp) ([]string, error) {

	var links []string

	mainpage, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if mainpage.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s", mainpage.Status)
	}

	defer mainpage.Body.Close()

	scanner := bufio.NewScanner(mainpage.Body)
	for scanner.Scan() {
		line := scanner.Bytes()
		match := re.FindSubmatch(line)
		if match == nil {
			continue
		}

		links = append(links, string(match[1]))
	}

	return links, nil
}
