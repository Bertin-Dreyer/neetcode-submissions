func hasDuplicate(nums []int) bool {
    dupeMap := make(map[int]int,len(nums))

	for _, num := range nums {
		dupeMap[num]++

		if dupeMap[num] > 1 {
			return true
		}

	}
	return false
}
