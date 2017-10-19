package secretmessage

import (
    "fmt"
)

// Encode func
func Encode(encoded string) string {
        var keys []string
	var counter map[string]int
        var currentKey string

	counter = make(map[string]int)

	for _, runeVal := range encoded {
            currentKey = string(runeVal)
            _, inMap := counter[currentKey]
	    if !inMap {
                counter[currentKey] = 1
		keys = append(keys, currentKey)
	    } else {
                counter[currentKey] = counter[currentKey] + 1
	    }
	}

	fmt.Println("%v", counter)
	fmt.Println("%v", keys)

	// @TODO do insertion sort on counter
	return ""
}
