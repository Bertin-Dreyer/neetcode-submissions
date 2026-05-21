func hasDuplicate(nums []int) bool {
    dupeCheck := make(map[int]bool)

    for _, value := range nums{
        if dupeCheck[value] == true{
            return true
        }
        dupeCheck[value] = true
    }

    return false
}
