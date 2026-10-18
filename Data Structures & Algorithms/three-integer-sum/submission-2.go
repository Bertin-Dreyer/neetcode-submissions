import "slices"
func threeSum(nums []int) [][]int {
	output := [][]int{}
	// Sort the slice 
	slices.Sort(nums)

	for i := range nums {
		// Turn into 2sum
		target := -(nums[i])
		right := len(nums) - 1
		left := i + 1

		if i > 0 && nums[i] == nums[i - 1] {
				continue
			}

		for left < right  {
			check := nums[left] + nums[right]

			if target > check {
				left++
				continue
			}

			if target < check {
				right--
				continue
			}

			if target == check {
				output = append(output, []int{nums[i], nums[left], nums[right]})
				right--
				left++

				for left < right && nums[left] == nums[left - 1] {
				left++
				}
			}
		}
	}
	return output
}
