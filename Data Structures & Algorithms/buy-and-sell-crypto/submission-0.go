func maxProfit(prices []int) int {

	//price[i]'s value is the price on the i'th day
	// price[7] = 10 is the price on the 7th day that was/is 10
	// sell and buy is unique days
	left := 0
	maxProfit := 0

	for right := 1; right < len(prices); right++ {
		if prices[right] > prices[left] {
			profit := prices[right] - prices[left]
			if profit > maxProfit {
				maxProfit = profit
			}
		} else {
			left = right
		}
	}
	return maxProfit
}