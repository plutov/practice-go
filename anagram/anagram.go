package anagram

import "strings"

func FindAnagrams(dictionary []string, word string) (result []string) {
	loweredWord := strings.TrimSpace(strings.ToLower(word))
	cachedParsing := parseCharDic(loweredWord)
	for _, value := range dictionary {
		if isAnagram(loweredWord, strings.ToLower(value), cachedParsing) {
			result = append(result, value)
		}
	}
	return result
}

func isAnagram(word1, word2 string, parsedWord1 map[rune]int) bool {
	if word1 == word2 {
		return false
	}

	dic2 := parseCharDic(word2)
	return compareDics(parsedWord1, dic2)
}

func parseCharDic(word string) (result map[rune]int) {
	result = make(map[rune]int)
	for _, char := range word {
		result[char] = result[char] + 1
	}
	delete(result, ' ')

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
