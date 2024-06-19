package main

import "fmt"

func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMax(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = maxValue(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		dp[i] = maxValue(nums[i]+dp[i-2], dp[i-1])
	}

	return dp[len(nums)-1]
}

func main() {
	arr := []int{5, 1, 1, 5}
	fmt.Println("Max Sum:", findMax(arr))
}
