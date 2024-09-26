package main

import (
	"fmt"
	"sort"
)

// Ft_coin function
func Ft_coin(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}

	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i >= coin {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

// Ft_missing function
func Ft_missing(nums []int) int {
	n := len(nums)
	expectedSum := n * (n + 1) / 2
	actualSum := 0
	for _, num := range nums {
		actualSum += num
	}
	return expectedSum - actualSum
}

// Ft_non_overlap function
func Ft_non_overlap(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	count := 0
	end := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < end {
			count++
		} else {
			end = intervals[i][1]
		}
	}

	return count
}

// Ft_profit function
func Ft_profit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	minPrice := prices[0]
	maxProfit := 0

	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if price-minPrice > maxProfit {
			maxProfit = price - minPrice
		}
	}

	return maxProfit
}

// Ft_max_substring function
func Ft_max_substring(s string) int {
	charIndex := make(map[byte]int)
	maxLen, start := 0, 0

	for i := 0; i < len(s); i++ {
		if lastIndex, found := charIndex[s[i]]; found && lastIndex >= start {
			start = lastIndex + 1
		}
		charIndex[s[i]] = i
		maxLen = max(maxLen, i-start+1)
	}

	return maxLen
}

// Ft_min_window function
func Ft_min_window(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}

	targetCount := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		targetCount[t[i]]++
	}

	left, right, start, minLen, formed := 0, 0, 0, len(s)+1, 0
	windowCount := make(map[byte]int)

	for right < len(s) {
		windowCount[s[right]]++

		if windowCount[s[right]] == targetCount[s[right]] {
			formed++
		}

		for left <= right && formed == len(targetCount) {
			if right-left+1 < minLen {
				start = left
				minLen = right - left + 1
			}

			windowCount[s[left]]--
			if windowCount[s[left]] < targetCount[s[left]] {
				formed--
			}
			left++
		}
		right++
	}

	if minLen == len(s)+1 {
		return ""
	}

	return s[start : start+minLen]
}

// Utility functions
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Main function to test all
func main() {
	fmt.Println("Testing Ft_coin:")
	fmt.Println(Ft_coin([]int{1, 2, 5}, 11)) // Expected result: 3
	fmt.Println(Ft_coin([]int{2}, 3))        // Expected result: -1
	fmt.Println(Ft_coin([]int{1}, 0))        // Expected result: 0

	fmt.Println("\nTesting Ft_missing:")
	fmt.Println(Ft_missing([]int{3, 1, 2}))                   // Expected result: 0
	fmt.Println(Ft_missing([]int{0, 1}))                      // Expected result: 2
	fmt.Println(Ft_missing([]int{9, 6, 4, 2, 3, 5, 7, 0, 1})) // Expected result: 8

	fmt.Println("\nTesting Ft_non_overlap:")
	fmt.Println(Ft_non_overlap([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}})) // Expected result: 1
	fmt.Println(Ft_non_overlap([][]int{{1, 2}, {2, 3}}))                 // Expected result: 0
	fmt.Println(Ft_non_overlap([][]int{{1, 2}, {1, 2}, {1, 2}}))         // Expected result: 2

	fmt.Println("\nTesting Ft_profit:")
	fmt.Println(Ft_profit([]int{7, 1, 5, 3, 6, 4})) // Expected result: 5
	fmt.Println(Ft_profit([]int{7, 6, 4, 3, 1}))    // Expected result: 0

	fmt.Println("\nTesting Ft_max_substring:")
	fmt.Println(Ft_max_substring("abcabcbb")) // Expected result: 3
	fmt.Println(Ft_max_substring("bbbbb"))    // Expected result: 1

	fmt.Println("\nTesting Ft_min_window:")
	fmt.Println(Ft_min_window("ADOBECODEBANC", "ABC")) // Expected result: "BANC"
	fmt.Println(Ft_min_window("a", "aa"))              // Expected result: ""
}
