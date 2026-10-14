func isPalindrome(s string) bool {
	start := 0
	s = strings.ToLower(s)
	
	var cleaned []rune
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
		cleaned = append(cleaned, unicode.ToLower(r))
	}
}

	for i := len(cleaned) -1 ; i >= 0 ; i-- {
		if cleaned[i] == cleaned[start] {
			start++
			continue
		}

		if cleaned[i] != cleaned[start] {
			return false
		}
	}
	return true
}
