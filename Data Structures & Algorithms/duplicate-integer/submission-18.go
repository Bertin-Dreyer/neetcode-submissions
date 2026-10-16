func hasDuplicate(nums []int) bool {
	// seen := make(map[int]bool)
	// Memory trick since empty struct hold 0 bytes.
	seen := make(map[int]struct{})

	for _, num := range nums {
		
		if _, seen := seen[num] ; seen {
			return true
		}
		seen[num] = struct{}{}
	}
	return false
}
