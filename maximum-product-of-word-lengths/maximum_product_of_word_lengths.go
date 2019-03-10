func maxProduct(words []string) int {
	i, j, result := 0, 0, 0
	for i = 0; i < len(words); i++ {
		chars := []rune(words[i])
		for j = i + 1; j < len(words); j++ {
			if isContain(chars, words[j]) {
				continue
			}
			result = max(result, len(words[i])*len(words[j]))
		}
	}
	return result
}

func isContain(chars []rune, word string) bool {
	for _, char := range chars {
		for _, c := range word {
			if char == c {
				return true
			}
		}
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
