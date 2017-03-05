package anagram

import "strings"

func FindAnagrams(dictionary []string, word string) (result []string) {
	word = normalize(word)
	if len(word) == 0 {
		return nil
	}

	charDir := parseCharDic(word)
	for _, value := range dictionary {
		ww := normalize(value)
		// ignore exact match or empty word(not anagram)
		if ww == word || len(ww) == 0 {
			continue
		}

		if compareDics(charDir, parseCharDic(ww)) {
			result = append(result, value)
		}
	}
	return result
}

func normalize(s string) string {
	return strings.Replace(strings.ToLower(s), " ", "", -1)
}

func parseCharDic(word string) (result map[rune]int) {
	result = make(map[rune]int)
	for _, char := range word {
		result[char] = result[char] + 1
	}

	return result
}

func compareDics(dic1, dic2 map[rune]int) bool {
	if len(dic1) != len(dic2) {
		return false
	}

	for key, value := range dic1 {
		if dic2[key] != value {
			return false
		}
	}

	return true
}
