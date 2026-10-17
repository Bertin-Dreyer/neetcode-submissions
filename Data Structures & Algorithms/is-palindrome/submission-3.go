func isPalindrome(s string) bool {
 
 var cleaned []rune

 for _, char := range s {
	if unicode.IsDigit(char) || unicode.IsLetter(char) {
		cleaned = append(cleaned, unicode.ToLower(char))
	}
 }

 right := len(cleaned) - 1

 for left := 0 ; left < right ; left++ {
	if cleaned[left] != cleaned[right]{
		return false
	}
	right--
 }
 return true
}
