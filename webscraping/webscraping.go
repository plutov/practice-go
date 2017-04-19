package webscraping

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

const timeServiceURL = "http://tycho.usno.navy.mil/cgi-bin/timer.pl"

var parser = regexp.MustCompile(`<BR>(.*? ([A-Z]+))\t`)

func GetTime(timezone string) string {
	page, err := fetchPage()
	if err != nil {
		return ""
	}

	times := parsePage(page)

	if time, ok := times[timezone]; ok {
		return time
	}

	return ""
}

func fetchPage() (string, error) {
	client := http.Client{
		Timeout: time.Duration(10 * time.Second),
	}
	resp, err := client.Get(timeServiceURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	page, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(page), nil
}

func parsePage(page string) map[string]string {
	result := map[string]string{}

	matches := parser.FindAllStringSubmatch(page, -1)
	for _, m := range matches {
		result[m[2]] = m[1]
	}

	return result
}
