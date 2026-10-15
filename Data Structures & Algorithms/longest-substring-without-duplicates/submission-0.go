func lengthOfLongestSubstring(s string) int {
    runes := []rune(s)
    seen := make(map[rune]bool)
    left := 0
    best := 0

    for right := 0; right < len(runes); right++ {
        for seen[runes[right]] {
            delete(seen, runes[left])
            left++
        }
        seen[runes[right]] = true

        size := right - left + 1
        if size > best {
            best = size
        }
    }
    return best
}