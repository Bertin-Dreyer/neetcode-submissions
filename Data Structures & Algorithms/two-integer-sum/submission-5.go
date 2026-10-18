func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)
    for i, num := range nums {
		check := target - num
		// Check if we have seen the target yet.
		if j, exists := seen[check] ; exists {
			return []int{j,i}
		}

		seen[num] = i
	}
	return []int{}
}
