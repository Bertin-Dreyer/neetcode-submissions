func lengthOfLongestSubstring(s string) int {
	// Sliding Window
	seen := make(map[rune]int)
	start := 0
	best := 0

	for end , char := range s {

		if i, exist := seen[char] ; exist && i >= start {
			start = i + 1
		} 
		
		seen[char] = end

		if end-start + 1 > best {
			best = end - start + 1
		}
	}
	return best
}
