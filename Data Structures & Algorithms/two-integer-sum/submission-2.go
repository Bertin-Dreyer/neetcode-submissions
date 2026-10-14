func twoSum(nums []int, target int) []int {
    twoSumMap := make(map[int]int)

    for i, num := range nums {
        
        complement := target - num
        
        if index, exists := twoSumMap[complement]; exists {
            return []int{index, i}
    }
        
        twoSumMap[num] = i
        
    }
    
    return nil
}