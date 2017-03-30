package shorthash

func GenerateShortHashes(dictionary string, len int) []string {
	result := []string{}

	if len == 0 {
		return result
	}

	for _, c := range dictionary {
		result = append(result, string(c))

		for _, t := range GenerateShortHashes(dictionary, len-1) {
			result = append(result, string(c)+t)
		}
	}

	return result
}
