func isAnagram(s string, t string) bool {
	if len(s) != len(t){
		return false
	}
	anagramMap := make(map[rune]int)

	sRunes := []rune(s)
	tRunes := []rune(t)
	
	for _, char := range sRunes {
		anagramMap[char]++
	}

	for _, char := range tRunes {
		anagramMap[char]--

		if anagramMap[char]< 0 {
			return false
		}
	}

	for _, value := range anagramMap {
		if value != 0 {
			return false
		}

	}
	return true
}
