func hasDuplicate(nums []int) bool {
    freqCheck := make(map[int]int)

	for _, num := range nums {
		freqCheck[num]++
	}

	for _, value := range freqCheck {
		if value > 1 {
			return true
		}
	}

	return false
}
