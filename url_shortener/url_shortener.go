package url_shortener

import (
	"bufio"
	"fmt"
	"hash/fnv"
	"os"
	"strings"
	"sync"
)

var (
	urlMap = sync.Map{}
)

func Shorten(url string) string {
	hasher := fnv.New32a()
	hasher.Write([]byte(url))
	short := fmt.Sprintf("%x", hasher.Sum32())
	urlMap.Store(short, url)
	return short
}

func Expand(short string) string {
	if url, ok := urlMap.Load(short); ok {
		return url.(string)
	}
	return ""
}

func Run() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a URL to shorten or a short code to expand:")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if strings.HasPrefix(input, "http://") || strings.HasPrefix(input, "https://") {
		short := Shorten(input)
		fmt.Printf("Alright, here's your shortened URL: %s\n", short)
	} else {
		original := Expand(input)
		if original != "" {
			fmt.Printf("Here's the original URL: %s\n", original)
		} else {
			fmt.Println("Short code not found")
		}
	}
}
