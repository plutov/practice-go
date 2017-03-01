package anagram

func FindAnagrams(dictionary []string, word string) (result []string) {
	cachedParsing := parseCharDic(word)
	for _, value := range dictionary {
		if isAnagram(word, value, cachedParsing) {
			result = append(result, value)
		}
	}
	return result
}

func isAnagram(word1, word2 string, parsedWord1 map[rune]int) bool {
	if word1 == word2 {
		return false
	}
	if len(word1) != len(word2) {
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
