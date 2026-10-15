func isPalindrome(s string) bool {
	cleaned := []rune{}
	for _, char := range s {
		if unicode.IsDigit(char) || unicode.IsLetter(char) {
			cleaned = append(cleaned, unicode.ToLower(char))
		}
	}

	top := len(cleaned) - 1
	start := 0

	for start < top {
		if cleaned[start] == cleaned[top] {
			start++
			top--
			continue
		}else{
			return false
		}
	} 
	return true
}
