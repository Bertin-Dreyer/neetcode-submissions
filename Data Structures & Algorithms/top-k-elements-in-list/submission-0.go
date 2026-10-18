import "slices"
func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	topSlice := []int{}

	for _, num := range nums {
		freq[num]++
	}

	for  num := range freq{
		topSlice = append(topSlice, num)
	}

	slices.SortFunc(topSlice, func(a, b int) int {
    	return freq[b] - freq[a] 
	})

	return topSlice[:k]
}
