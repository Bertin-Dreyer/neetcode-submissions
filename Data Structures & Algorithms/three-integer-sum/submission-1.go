import "slices"
func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	output := [][]int{}

	for i := range nums {
		left := i + 1
		right := len(nums) - 1
		// sum := nums[i] + nums[left] + nums[right]
		target := -nums[i]
		
		if i > 0 && nums[i] == nums[i-1] {
        	continue 
    		}

		for left < right {

			if target > nums[left] + nums[right] {
						left++
						continue
					}

					if target < nums[left] + nums[right] {
						right--
						continue
					}else{
						output = append(output, []int{nums[i], nums[left], nums[right]})
						left++
    					right--
					}
					for left < right && nums[left] == nums[left-1] {
						left++
					}
		}
		// output = append(output, nums[i], nums[left], nums[right])
		// if left + right = -nums[i] then we found the 3 ?
	}
	return output
}