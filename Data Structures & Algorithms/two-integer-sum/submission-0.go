func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		if indexOfComplement, found := seen[complement]; found {
			return []int{indexOfComplement, i}
		}

		seen[num] = i
	}

	return nil
}