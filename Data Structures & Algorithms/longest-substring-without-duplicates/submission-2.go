func lengthOfLongestSubstring(s string) int {
	// Sliding Window
	// The checker
	seen := make(map[rune]int)
	start := 0
	best := 0


	for end, char := range s {

		// if we have seen the character, we move the start of the window to 1 past the end.
		if i, exists := seen[char] ; exists && i >= start {
			start = i + 1
		}
		// add character as we go
		seen[char] = end

		// if the current lenght is bigger than the prev recorded best, we do this:
		if end-start + 1 > best {
			best = end - start + 1
		}
	}

	return best
}
